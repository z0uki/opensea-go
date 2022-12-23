package opensea

import (
	"fmt"

	"github.com/z0uki/opensea-go/model"
)

// Contract Used to fetch more in-depth information about an contract asset
func (c *Client) Contract(req *ContractRequest) (*model.Contract, error) {
	var rsp, err = c.get(fmt.Sprintf("/api/v1/asset_contract/%s", req.AssetContractAddress), nil)
	if err != nil {
		return nil, err
	}

	var response model.Contract
	if err := ParseRsp(rsp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type ContractRequest struct {
	// Address of the contract
	AssetContractAddress string `json:"asset_contract_address,required"`
}
