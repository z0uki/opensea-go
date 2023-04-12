package opensea

import (
	"fmt"
	"log"
	"net/url"

	"github.com/nshafer/phx"
)

type EventType string

const (
	ItemMetadataUpdated EventType = "item_metadata_updated"
	ItemListed          EventType = "item_listed"
	ItemSold            EventType = "item_sold"
	ItemTransferred     EventType = "item_transferred"
	ItemReceivedOffer   EventType = "item_received_offer"
	ItemReceivedBid     EventType = "item_received_bid"
	ItemCancelled       EventType = "item_cancelled"
	CollectionOffer     EventType = "collection_offer"
)

type Network string

const (
	MAINNET Network = "mainnet"
	TESTNET Network = "testnet"
)

type StreamClient struct {
	socket   *phx.Socket
	channels map[EventType]*phx.Channel
}

func NewStreamClient(network Network, token string, logLevel phx.LoggerLevel, onError func(error)) *StreamClient {
	m := map[Network]string{
		MAINNET: "wss://stream.openseabeta.com/socket",
		TESTNET: "wss://testnets-stream.openseabeta.com/socket",
	}
	socketUrl := fmt.Sprintf("%s?token=%s", m[network], token)

	endPoint, _ := url.Parse(socketUrl)
	socket := phx.NewSocket(endPoint)

	socket.OnError(onError)
	socket.OnClose(func() {
		err := socket.Reconnect()
		if err != nil {
			onError(err)
		}
	})
	socket.Logger = phx.NewSimpleLogger(logLevel)
	return &StreamClient{
		socket:   socket,
		channels: make(map[EventType]*phx.Channel),
	}
}

func (s StreamClient) Connect() error {
	fmt.Println("Connecting to socket")
	return s.socket.Connect()
}
func (s *StreamClient) Disconnect() error {
	//s.socket.OnError()
	fmt.Println("Succesfully disconnected from socket")
	s.channels = make(map[EventType]*phx.Channel)
	return s.socket.Disconnect()
}
func (s *StreamClient) createChannel(topic string, eventType EventType) (channel *phx.Channel) {
	channel = s.socket.Channel(topic, nil)
	join, err := channel.Join()
	if err != nil {
		fmt.Println(err)
		return
	}
	join.Receive("ok", func(response any) {
		log.Println("Joined channel:", channel.Topic(), response)
	})
	join.Receive("error", func(response any) {
		log.Println("failed 2 joined channel:", channel.Topic(), response)
	})
	s.channels[eventType] = channel
	return
}

func (s StreamClient) getChannel(topic string, eventType EventType) (channel *phx.Channel) {
	var ok bool
	channel, ok = s.channels[eventType]
	if !ok {
		channel = s.createChannel(topic, eventType)
	}
	return channel
}

func (s StreamClient) on(eventType EventType, collectionSlug string, callback func(payload any)) func() {
	topic := collectionTopic(collectionSlug)
	fmt.Printf("Fetching channel %s\n", topic)
	channel := s.getChannel(topic, eventType)
	fmt.Printf("Subscribing to %s events on %s\n", eventType, topic)
	channel.On(string(eventType), callback)
	return func() {
		fmt.Printf("Unsubscribing from %s events on %s\n", eventType, topic)
		leave, err := channel.Leave()
		if err != nil {
			fmt.Println("channel.Leave err:", err)
		}
		leave.Receive("ok", func(response any) {
			delete(s.channels, eventType)
			fmt.Printf("Succesfully left channel %s listening for %s\n", topic, eventType)
		})
	}

}

func collectionTopic(slug string) string {
	return fmt.Sprintf("collection:%s", slug)
}
func (s StreamClient) OnItemListed(collectionSlug string, Callback func(itemListedEvent any)) func() {
	return s.on(ItemListed, collectionSlug, Callback)
}

func (s StreamClient) OnItemSold(collectionSlug string, Callback func(itemSoldEvent any)) func() {
	return s.on(ItemSold, collectionSlug, Callback)
}
func (s StreamClient) OnItemTransferred(collectionSlug string, Callback func(itemTransferredEvent any)) func() {
	return s.on(ItemTransferred, collectionSlug, Callback)
}
func (s StreamClient) OnItemCancelled(collectionSlug string, Callback func(itemCancelledEvent any)) func() {
	return s.on(ItemCancelled, collectionSlug, Callback)
}
func (s StreamClient) OnItemReceivedBid(collectionSlug string, Callback func(itemReceivedBidEvent any)) func() {
	return s.on(ItemReceivedBid, collectionSlug, Callback)
}
func (s StreamClient) OnItemReceivedOffer(collectionSlug string, Callback func(itemReceivedOfferEvent any)) func() {
	return s.on(ItemReceivedOffer, collectionSlug, Callback)
}
func (s StreamClient) OnItemMetadataUpdated(collectionSlug string, Callback func(itemMetadataUpdatedEvent any)) func() {
	return s.on(ItemMetadataUpdated, collectionSlug, Callback)
}
func (s StreamClient) OnCollectionOffer(Callback func(collectionOfferEvent any)) func() {
	return s.on(CollectionOffer, "*", Callback)
}
