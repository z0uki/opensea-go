package opensea

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
	"github.com/z0uki/opensea-go/model/stream"
	"log"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

var (
	ItemListedCount        uint64 = 0
	CollectionOfferCount   uint64 = 0
	ItemReceivedBidCount   uint64 = 0
	ItemReceivedOfferCount uint64 = 0
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
			return
		}
		atomic.AddUint64(&CollectionOfferCount, 1)
		//fmt.Println("collection offer:", collectionOfferEvent.Payload.Quantity)
	})

	client.OnItemReceivedBid("*", func(response any) {
		var event stream.ItemReceivedBidEvent
		err := mapstructure.Decode(response, &event)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
			return
		}
		atomic.AddUint64(&ItemReceivedBidCount, 1)
		//fmt.Println("bid:", event.Payload.Item.NftId)
	})

	client.OnItemReceivedOffer("*", func(response any) {
		var event stream.ItemReceivedOfferEvent
		err := mapstructure.Decode(response, &event)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
			return
		}
		atomic.AddUint64(&ItemReceivedOfferCount, 1)
		//fmt.Println("offer:", event.Payload.Item.NftId)
	})

	client.OnItemListed("*", func(response any) {
		var event stream.ItemListedEvent
		err := mapstructure.Decode(response, &event)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
			return
		}
		atomic.AddUint64(&ItemListedCount, 1)
		//fmt.Println("list:", event.Payload.Item.NftId)
	})

	startTime := time.Now()

	for {
		log.Printf("挂单数量: %d, 收到bid数量: %d, 收到collection_offer数量: %d 运行时间: %fs\n", ItemListedCount, ItemReceivedBidCount, CollectionOfferCount, time.Now().Sub(startTime).Seconds())
		time.Sleep(time.Second * 1)
	}

	//select {}
}
