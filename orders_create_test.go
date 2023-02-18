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
		CollectionSlug:    "autominter-pro-pass",
		Quantity:          1,
		PriceWei:          big.NewInt(7040000000000000),
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
		TokenAddress:      "0x71f67624a8b2f6310e0bc538ff78f843e1d9ee6e",
		TokenId:           big.NewInt(158),
		PriceWei:          big.NewInt(110000000000000000),
		ExpirationSeconds: big.NewInt(3600),
	}

	if rsp, err := client.OrdersCreateListings(&req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rsp)
	}

}
