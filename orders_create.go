package opensea

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethersphere/bee/pkg/crypto"
	"github.com/ethersphere/bee/pkg/crypto/eip712"
	"github.com/fatih/structs"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"github.com/z0uki/opensea-go/model"
)

// OrdersCreateListings 挂单
func (c *Client) OrdersCreateListings(req *OrdersCreateListingsRequest) (*OrdersCreateListingsResponse, error) {

	// 1. 获取版税信息
	contract, err := c.Contract(&ContractRequest{AssetContractAddress: req.TokenAddress})
	if err != nil {
		return nil, err
	}
	var creatorBasisPoints = big.NewInt(int64(contract.DevSellerFeeBasisPoints))
	var creatorFees = contract.Collection.Fees.SellerFees
	//获取合约类型 erc721 or erc1155
	var contractType uint8
	switch contract.SchemaName {
	case "ERC721":
		contractType = ItemType_ERC721
	case "ERC1155":
		contractType = ItemType_ERC1155
	default:
		return nil, fmt.Errorf("contract type error")
	}

	// 2. 检查是否授权
	approved, err := c.IsApproved(common.HexToAddress(req.TokenAddress))
	if err != nil {
		return nil, err
	}

	if !approved {
		log.Println("not approved, set approval...")
		if _, err := c.SetApprovalForAll(common.HexToAddress(req.TokenAddress)); err != nil {
			log.Println("set approval failed")
			return nil, err
		}
	}

	// 3. 构建订单
	order, err := c.buildSellOrder(req, creatorBasisPoints, creatorFees, contractType)
	if err != nil {
		return nil, err
	}

	// 4. 发送订单
	rsp, err := c.postSellOrder(order)
	if err != nil {
		return nil, err
	}
	//fmt.Println(rsp)
	var response OrdersCreateListingsResponse
	if err := ParseRsp(rsp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// OrdersCreateCollectionOffer 发送集合offer
func (c *Client) OrdersCreateCollectionOffer(req *OrdersCreateCollectionOfferRequest) (*OrdersCreateCollectionOfferResponse, error) {

	// 1. 获取系列信息
	collection, err := c.Collection(&CollectionRequest{CollectionSlug: req.CollectionSlug})
	if err != nil {
		return nil, err
	}
	var creatorFees = collection.Fees.SellerFees

	// 2. 构建订单
	offer, err := c.buildCollectionOffer(req, creatorFees)
	if err != nil {
		return nil, err
	}

	// 3. 发送订单
	rsp, err := c.postCollectionOffer(req.CollectionSlug, offer)
	if err != nil {
		return nil, err
	}
	//fmt.Println(rsp)
	var response OrdersCreateCollectionOfferResponse
	if err := ParseRsp(rsp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type OrdersCreateListingsRequest struct {
	TokenAddress      string
	TokenId           *big.Int
	PriceWei          *big.Int
	ExpirationSeconds *big.Int
}

type OrdersCreateListingsResponse struct {
	Order model.Order `json:"order"`
}

type OrdersCreateCollectionOfferRequest struct {
	CollectionSlug    string
	Quantity          int
	PriceWei          *big.Int
	ExpirationSeconds *big.Int
}

type OrdersCreateCollectionOfferResponse struct {
	OrderHash string `json:"order_hash"`
	Chain     string `json:"chain"`
	Criteria  struct {
		Collection struct {
			Slug string `json:"slug"`
		}
		Contract struct {
			Address string `json:"address"`
		}
		Trait struct{}
	} `json:"criteria"`
	ProtocolData *model.Protocol `json:"protocol_data"`
}

func (c *Client) buildSellOrder(req *OrdersCreateListingsRequest, creatorBasisPoints *big.Int, creatorFees *model.Fee, contractType uint8) (*model.Protocol, error) {
	startTime := time.Now().Unix()
	endTime := startTime + req.ExpirationSeconds.Int64()
	rand.Seed(time.Now().UnixNano())
	salt := rand.Int63n(1000000000)

	sellerFee := model.ConsiderationItem{
		ItemType:             ItemType_NATIVE,
		Token:                ZeroAddress,
		IdentifierOrCriteria: "0",
		StartAmount:          strconv.FormatInt(CalcEarnings(req.PriceWei, creatorBasisPoints), 10),
		EndAmount:            strconv.FormatInt(CalcEarnings(req.PriceWei, creatorBasisPoints), 10),
		Recipient:            c.Wallet.Address.Hex(),
	}

	// opensea 修改了协议，不再需要手续费
	//platformFee := model.ConsiderationItem{
	//	ItemType:             ItemType_NATIVE,
	//	Token:                ZeroAddress,
	//	IdentifierOrCriteria: "0",
	//	StartAmount:          strconv.FormatInt(CalcOpenSeaFeeByBasePrice(req.PriceWei), 10),
	//	EndAmount:            strconv.FormatInt(CalcOpenSeaFeeByBasePrice(req.PriceWei), 10),
	//	Recipient:            OpenSeaFeeRecipient,
	//}

	considerations := []model.ConsiderationItem{sellerFee}

	for recipient, points := range *creatorFees {
		collectionSellerFees := model.ConsiderationItem{
			ItemType:             ItemType_NATIVE,
			Token:                ZeroAddress,
			IdentifierOrCriteria: "0",
			StartAmount:          strconv.FormatInt(CalcFeeByBasisPoints(req.PriceWei, points), 10),
			EndAmount:            strconv.FormatInt(CalcFeeByBasisPoints(req.PriceWei, points), 10),
			Recipient:            recipient,
		}
		considerations = append(considerations, collectionSellerFees)
	}

	parameters := model.Parameters{
		Offerer: c.Wallet.Address.Hex(),
		Offer: []model.OfferItem{
			{
				ItemType:             contractType,
				Token:                req.TokenAddress,
				IdentifierOrCriteria: req.TokenId.String(),
				StartAmount:          "1",
				EndAmount:            "1",
			},
		},
		Consideration:                   considerations,
		StartTime:                       strconv.FormatInt(startTime, 10),
		EndTime:                         strconv.FormatInt(endTime, 10),
		OrderType:                       0, // PARTIAL_RESTRICTED
		Zone:                            Zone,
		ZoneHash:                        ZoneHash,
		Salt:                            strconv.FormatInt(salt, 10),
		ConduitKey:                      ConduitKey,
		TotalOriginalConsiderationItems: len(considerations),
		Counter:                         "0",
	}

	signature, err := c.signParameters(parameters)
	if err != nil {
		return nil, err
	}

	order := model.Protocol{
		Parameters: parameters,
		Signature:  signature,
	}

	return &order, nil
}

func (c *Client) buildCollectionOffer(req *OrdersCreateCollectionOfferRequest, creatorFees *model.Fee) (*model.Protocol, error) {

	// getTokenConsideration
	body := map[string]interface{}{
		"offerer":  c.Wallet.Address,
		"quantity": req.Quantity,
		"criteria": map[string]interface{}{
			"collection": map[string]interface{}{
				"slug": req.CollectionSlug,
			},
		},
	}

	rsp, err := c.post("/v2/offers/build", body)
	if err != nil {
		return nil, err
	}

	//fmt.Println(rsp)

	tokenConsideration := model.ConsiderationItem{
		ItemType:             uint8(gjson.Get(rsp.String(), "partialParameters.consideration.0.itemType").Uint()),
		Token:                gjson.Get(rsp.String(), "partialParameters.consideration.0.token").String(),
		IdentifierOrCriteria: gjson.Get(rsp.String(), "partialParameters.consideration.0.identifierOrCriteria").String(),
		StartAmount:          gjson.Get(rsp.String(), "partialParameters.consideration.0.startAmount").String(),
		EndAmount:            gjson.Get(rsp.String(), "partialParameters.consideration.0.endAmount").String(),
		Recipient:            gjson.Get(rsp.String(), "partialParameters.consideration.0.recipient").String(),
	}

	platformFee := model.ConsiderationItem{
		ItemType:             ItemType_ERC20,
		Token:                WethAddress,
		IdentifierOrCriteria: "0",
		StartAmount:          strconv.FormatInt(CalcOpenSeaFeeByBasePrice(req.PriceWei), 10),
		EndAmount:            strconv.FormatInt(CalcOpenSeaFeeByBasePrice(req.PriceWei), 10),
		Recipient:            OpenSeaFeeRecipient,
	}

	considerations := []model.ConsiderationItem{tokenConsideration, platformFee}

	for recipient, points := range *creatorFees {
		collectionSellerFees := model.ConsiderationItem{
			ItemType:             ItemType_ERC20,
			Token:                WethAddress,
			IdentifierOrCriteria: "0",
			StartAmount:          strconv.FormatInt(CalcFeeByBasisPoints(req.PriceWei, points), 10),
			EndAmount:            strconv.FormatInt(CalcFeeByBasisPoints(req.PriceWei, points), 10),
			Recipient:            recipient,
		}
		considerations = append(considerations, collectionSellerFees)
	}

	startTime := time.Now().Unix()
	endTime := startTime + req.ExpirationSeconds.Int64()
	rand.Seed(time.Now().UnixNano())
	salt := rand.Int63n(1000000000)

	parameters := model.Parameters{
		Offerer: c.Wallet.Address.Hex(),
		Offer: []model.OfferItem{
			{
				ItemType:             ItemType_ERC20,
				Token:                WethAddress,
				IdentifierOrCriteria: "0",
				StartAmount:          req.PriceWei.String(),
				EndAmount:            req.PriceWei.String(),
			},
		},
		Consideration:                   considerations,
		StartTime:                       strconv.FormatInt(startTime, 10),
		EndTime:                         strconv.FormatInt(endTime, 10),
		OrderType:                       2,
		Zone:                            Zone,
		ZoneHash:                        ZoneHash,
		Salt:                            strconv.FormatInt(salt, 10),
		ConduitKey:                      ConduitKey,
		TotalOriginalConsiderationItems: len(considerations),
		Counter:                         "1",
	}

	signature, err := c.signParameters(parameters)
	if err != nil {
		return nil, err
	}

	offer := model.Protocol{
		Parameters: parameters,
		Signature:  signature,
	}

	return &offer, nil
}

func (c *Client) signParameters(offer model.Parameters) (string, error) {
	signer := crypto.NewDefaultSigner(c.Wallet.PrivateKey)
	messageToSign := eip712.TypedData{
		Domain: eip712.TypedDataDomain{
			Name:              "Seaport",
			Version:           "1.4",
			ChainId:           math.NewHexOrDecimal256(1),
			VerifyingContract: "0x00000000000001ad428e4906ae43d8f9852d0dd6",
		},
		Types:       TYPES,
		PrimaryType: "OrderComponents",
		Message:     formatParameters(structs.Map(&offer)),
	}

	signature, err := signer.SignTypedData(&messageToSign)
	if err != nil {
		return "", err
	}
	return "0x" + hex.EncodeToString(signature), nil
}

func (c *Client) postSellOrder(order *model.Protocol) (*req.Response, error) {
	return c.post("/v2/orders/ethereum/seaport/listings", order)
}

func (c *Client) postCollectionOffer(slug string, offer *model.Protocol) (*req.Response, error) {
	body := map[string]interface{}{
		"criteria": map[string]interface{}{
			"collection": map[string]interface{}{
				"slug": slug,
			},
		},
		"protocol_data": offer,
	}

	return c.post("/v2/offers", body)
}

func formatParameters(v map[string]interface{}) map[string]interface{} {
	for key, val := range v {
		if reflect.TypeOf(val).Kind() == reflect.Map { // nested map
			val = formatParameters(val.(map[string]interface{}))
		}
		if reflect.TypeOf(val).Kind() == reflect.Slice {
			var newSlice []interface{}
			for _, v := range val.([]interface{}) {
				newSlice = append(newSlice, formatParameters(v.(map[string]interface{})))
			}
			val = newSlice
		}

		if reflect.TypeOf(val).Kind() == reflect.Int {
			val = strconv.Itoa(val.(int))
		}
		if reflect.TypeOf(val).Kind() == reflect.Int64 {
			val = strconv.FormatInt(val.(int64), 10)
		}
		if reflect.TypeOf(val).Kind() == reflect.Uint8 {
			val = strconv.Itoa(int(val.(uint8)))
		}
		if key == "TotalOriginalConsiderationItems" {
			delete(v, key)
		} else {
			delete(v, key)
			v[firstLower(key)] = val
		}
	}
	return v
}

func firstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
