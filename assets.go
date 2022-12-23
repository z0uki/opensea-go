package opensea

import "github.com/z0uki/opensea-go/model"

// Assets To retrieve assets from our API, call the /assets endpoint with the desired filter parameters.
func (c *Client) Assets(req *AssetsRequest) (*AssetsResponse, error) {
	var rsp, err = c.get("/api/v1/assets", ObjectParams(req))
	if err != nil {
		return nil, err
	}

	var response AssetsResponse
	if err = ParseRsp(rsp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type AssetsRequest struct {
	// The address of the owner of the assets
	Owner string `structs:"owner"`
	// An array of token IDs to search for (e.g. ?token_ids=1&token_ids=209).
	// Will return a list of assets with token_id matching any of the IDs in this array.
	TokenIDs []string `structs:"token_ids"`
	// The NFT contract address for the assets
	AssetContractAddress string `structs:"asset_contract_address"`
	// An array of contract addresses to search for (e.g. ?asset_contract_addresses=0x1...&asset_contract_addresses=0x2...).
	// Will return a list of assets with contracts matching any of the addresses in this array.
	// If "token_ids" is also specified, then it will only return assets that match each (address, token_id) pairing, respecting order.
	AssetContractAddresses []string `structs:"asset_contract_addresses"`
	// How to order the assets returned. By default, the API returns the fastest ordering.
	// Options you can set are sale_date (the last sale's transaction's timestamp),
	// sale_count (number of sales), and sale_price (the last sale's total_price)
	OrderBy string `structs:"order_by"`
	// Can be asc for ascending or desc for descending
	OrderDirection string `structs:"order_direction,required"`
	// Offset
	Offset int32 `structs:"offset"`
	// Limit
	Limit int32 `structs:"limit,required"`

	//Cursor A cursor pointing to the page to retrieve
	Cursor string `structs:"cursor"`

	// Limit responses to members of a collection.
	// Case-sensitive and must match the collection slug exactly.
	// Will return all assets from all contracts in a collection.
	// For more information on collections, see our collections documentation.
	Collection string `structs:"collection"`

	// Limit responses to members of a collection.
	// Case sensitive and must match the collection slug exactly.
	// Will return all assets from all contracts in a collection.
	// For more information on collections,
	CollectionSlug string `structs:"collection_slug"`

	// CollectionEditor
	CollectionEditor string `structs:"collection_editor"`

	// A flag determining if order information should be included in the response.
	IncludeOrders bool `structs:"include_orders"`
}

type AssetsResponse struct {
	Assets   []*model.Asset `json:"assets"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
}
