package opensea

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
	"github.com/z0uki/opensea-go/model/stream"
	"os"
	"strconv"
	"testing"
	"time"
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
		parseInt, err := strconv.ParseInt(collectionOfferEvent.Payload.ProtocolData.Parameters.EndTime, 10, 64)
		if err != nil {
			return
		}
		fmt.Println("collection name:", (parseInt-time.Now().Unix()+60)/60)
	})

	select {}
}
