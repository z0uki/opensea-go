package opensea

import (
	"fmt"
	"github.com/z0uki/opensea-go/model"
)

// Asset Used to fetch more in-depth information about an individual asset
func (c *Client) Asset(req *AssetRequest) (*model.Asset, error) {
	var params = AssetRequest{
		AccountAddress: req.AccountAddress,
		IncludeOrders:  req.IncludeOrders,
	}
	var rsp, err = c.get(fmt.Sprintf("/api/v1/asset/%s/%s/", req.AssetContractAddress, req.TokenID), ObjectParams(&params))
	if err != nil {
		return nil, err
	}

	var response model.Asset
	if err := ParseRsp(rsp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type AssetRequest struct {
	AssetContractAddress string
	TokenID              string
	AccountAddress       string `structs:"account_address"`
	IncludeOrders        bool   `structs:"include_orders"`
}
