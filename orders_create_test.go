package opensea

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var client *Client

func init() {
	godotenv.Load()
	pkey := os.Getenv("PRIVATE_KEY")
	apiKey := os.Getenv("API_KEY")
	client, _ = New(WithAPIKey(apiKey), WithWallet(pkey))
}

func TestOrdersCreateCollectionOffer(t *testing.T) {
	req := OrdersCreateCollectionOfferRequest{
		CollectionSlug:    "bokinftofficial",
		Quantity:          1,
		PriceWei:          big.NewInt(10040000000000000),
		ExpirationSeconds: big.NewInt(901),
	}

	if rsp, err := client.OrdersCreateCollectionOffer(&req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rsp.ProtocolData)
	}
}

func TestOrdersCreateListings(t *testing.T) {
	req := OrdersCreateListingsRequest{
		TokenAddress:      "0x2bdbe17a6f148e686581bb3fb78186aadad9e73d",
		TokenId:           big.NewInt(673),
		PriceWei:          big.NewInt(1300000000000000000),
		ExpirationSeconds: big.NewInt(10000),
	}

	if rsp, err := client.OrdersCreateListings(&req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rsp)
	}

}