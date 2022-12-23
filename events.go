package opensea

import (
	"github.com/z0uki/opensea-go/model"
)

// Events 获取Events
func (c *Client) Events(req *EventsRequest) (*EventsResponse, error) {
	var rsp, err = c.get("/api/v1/events", ObjectParams(req))
	if err != nil {
		return nil, err
	}

	var response EventsResponse
	if err := ParseRsp(rsp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type EventsRequest struct {
	// Restrict to events on OpenSea auctions. Can be true or false
	OnlyOpensea bool `structs:"only_opensea"`
	// The token's id to optionally filter by
	TokenID string `structs:"token_id"`
	// The NFT contract address for the assets
	AssetContractAddress string `structs:"asset_contract_address"`
	// Limit responses to events from a collection. Case sensitive and must match the collection slug exactly.
	// Will return all assets from all contracts in a collection.
	// For more information on collections, see our collections documentation.
	CollectionSlug   string `structs:"collection_slug"`
	CollectionEditor string `structs:"collection_editor"`
	// A user account's wallet address to filter for events on an account
	AccountAddress string `structs:"account_address"`
	// The event type to filter.
	EventType model.EventType `structs:"event_type"`
	// Filter by an auction type.
	AuctionType model.AuctionType `structs:"auction_type"`
	// Only show events listed before this timestamp. Seconds since the Unix epoch.
	// eg: 2017-07-21T17:32:28Z
	OccurredBefore string `structs:"occurred_before"`
	// Only show events listed after this timestamp. Seconds since the Unix epoch.
	// eg: 2017-07-21T17:32:28Z
	OccurredAfter string `structs:"occurred_after"`
	// A cursor pointing to the page to retrieve
	Cursor string `structs:"cursor"`
}

type EventsResponse struct {
	AssetEvents []*model.Event `json:"asset_events"`
	Next        string         `json:"next"`
	Previous    string         `json:"previous"`
}
