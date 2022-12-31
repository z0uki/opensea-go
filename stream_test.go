package opensea

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
	"github.com/z0uki/opensea-go/model/stream"
	"os"
	"testing"
)

func TestNewStreamClient(t *testing.T) {
	client := NewStreamClient(MAINNET, os.Getenv("API_KEY"), phx.LogInfo, func(err error) {
		fmt.Println("opensea.NewStreamClient err:", err)
	})
	if err := client.Connect(); err != nil {
		fmt.Println("client.Connect err:", err)
		return
	}

	client.OnCollectionOffer(func(response any) {
		var collectionOfferEvent stream.CollectionOfferEvent
		err := mapstructure.Decode(response, &collectionOfferEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Println("maker:", collectionOfferEvent.Payload.Maker.Address, "collection name:", collectionOfferEvent.Payload.Collection.Slug, "offer price:", collectionOfferEvent.Payload.BasePrice)
	})

	select {}
}
