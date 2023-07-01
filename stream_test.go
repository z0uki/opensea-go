package opensea

import (
	"fmt"
	"os"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
	"github.com/z0uki/opensea-go/model/stream"
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
	//client.OnItemCancelled("*", func(response any) {
	//	var event stream.ItemCancelledEvent
	//	err := mapstructure.Decode(response, &event)
	//	if err != nil {
	//		log.Println("mapstructure.Decode err:", err)
	//		return
	//	}
	//
	//	marshal, err := json.Marshal(event)
	//	if err != nil {
	//		return
	//	}
	//
	//	fmt.Println(string(marshal))
	//})

	client.OnCollectionOffer(func(response any) {
		var event stream.CollectionOfferEvent
		err := mapstructure.Decode(response, &event)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
			return
		}
		//atomic.AddUint64(&CollectionOfferCount, 1)

		fmt.Println(event.EventType)

		//for _, c := range event.Payload.ProtocolData.Parameters.Consideration {
		//	if c.ItemType > 2 {
		//		in, _ := new(big.Int).SetString(c.IdentifierOrCriteria, 10)
		//		//fmt.Println(in)
		//		fmt.Println("list:" + strings.ToLower(c.Token) + ":" + common.BigToHash(in).String())
		//	}
		//}

	})

	client.OnItemReceivedBid("*", func(response any) {
		var event stream.ItemReceivedBidEvent
		err := mapstructure.Decode(response, &event)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
			return
		}
		fmt.Println(event.EventType)
		//fmt.Println("bid:", event.Payload.Item.NftId)
	})

	//client.OnItemReceivedOffer("*", func(response any) {
	//	var event stream.ItemReceivedOfferEvent
	//	err := mapstructure.Decode(response, &event)
	//	if err != nil {
	//		fmt.Println("mapstructure.Decode err:", err)
	//		return
	//	}
	//	atomic.AddUint64(&ItemReceivedOfferCount, 1)
	//	//fmt.Println("offer:", event.Payload.Item.NftId)
	//})
	//
	client.OnItemListed("*", func(response any) {
		var event stream.ItemListedEvent
		err := mapstructure.Decode(response, &event)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
			return
		}
		//atomic.AddUint64(&ItemListedCount, 1)
		fmt.Println("list:", event.Payload.Item.NftId)
	})

	//startTime := time.Now()
	//
	//for {
	//	log.Printf("挂单数量: %d, 收到bid数量: %d, 收到collection_offer数量: %d 运行时间: %fs\n", ItemListedCount, ItemReceivedBidCount, CollectionOfferCount, time.Now().Sub(startTime).Seconds())
	//	time.Sleep(time.Second * 1)
	//}

	select {}
}

func TestName(t *testing.T) {
	client := NewStreamClient(MAINNET, os.Getenv("API_KEY"), phx.LogInfo, func(err error) {
		fmt.Println("opensea.NewStreamClient err:", err)
	})
	if err := client.Connect(); err != nil {
		fmt.Println("client.Connect err:", err)
		return
	}

	client.OnItemInvalidation("*", func(response any) {
		var event stream.ItemInvalidationEvent
		err := mapstructure.Decode(response, &event)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
			return
		}
		fmt.Println(event.Payload.OrderHash, event.Payload.ProtocolAddress)
	})

	select {}
}
