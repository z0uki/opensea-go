package opensea

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	req "github.com/imroc/req/v3"
)

const (
	APIKey  = "X-API-KEY"
	RPC_URL = "https://cloudflare-eth.com"
)

type Client struct {
	*req.Client
	*option

	Wallet  *Wallet
	eclient *ethclient.Client
}

func New(fnList ...OptionFn) (*Client, error) {
	var o = defaultOption
	for _, fn := range fnList {
		fn(o)
	}

	ec, err := ethclient.Dial(RPC_URL)

	privateKey, err := crypto.HexToECDSA(o.privateKey)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Client{
		Client: req.NewClient().
			SetBaseURL(o.baseURL).
			SetCommonRetryCount(3).
			SetCommonRetryFixedInterval(3*time.Second).
			SetCommonRetryCondition(func(resp *req.Response, err error) bool {
				return err != nil || resp.StatusCode == http.StatusTooManyRequests
			}).
			SetCommonHeader(APIKey, o.apiKey).
			SetCommonContentType("application/json").
			SetCommonHeader("Accept", "application/json"),

		option: o,
		Wallet: &Wallet{
			PrivateKey: privateKey,
			Address:    address,
		},
		eclient: ec,
	}, nil
}

func (c *Client) get(resource string, params map[string]string) (*req.Response, error) {
	rsp, err := c.R().SetQueryParams(params).Get(resource)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *Client) post(resource string, body interface{}) (*req.Response, error) {

	rsp, err := c.R().SetBody(body).Post(resource)

	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func ObjectParams(v interface{}) map[string]string {
	m := make(map[string]string)
	for k, v := range structs.Map(v) {
		//如果类型是切片 且长度为0 则不添加
		if _, ok := v.([]string); ok && len(v.([]string)) == 0 {
			continue
		}
		if v != "" && v != nil && v != 0 {
			m[k] = fmt.Sprintf("%v", v)
		}
	}
	return m
}

var (
	ErrNotFound        = errors.New("resource.not.found")
	ErrTooManyRequests = errors.New("too.many.requests")
	ErrInternalServer  = errors.New("internal.server.error")
)

func ParseRsp(rsp *req.Response, i interface{}) error {
	var statusCode = rsp.GetStatusCode()
	switch statusCode {
	case http.StatusOK:
		return rsp.Into(i)
	case http.StatusBadRequest, http.StatusNotAcceptable:
		return fmt.Errorf("bad.request:%s", rsp.String())
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusTooManyRequests:
		return ErrTooManyRequests
	case http.StatusInternalServerError:
		return ErrInternalServer
	default:
		return fmt.Errorf("unknown.http.status.code: %d", statusCode)
	}
}
