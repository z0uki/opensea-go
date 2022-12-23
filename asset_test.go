package opensea

import "testing"

func TestAsset(t *testing.T) {
	req := AssetRequest{
		AssetContractAddress: "0x7831fd9527f142d95c1445387870d0eb6be55180",
		TokenID:              "8",
		AccountAddress:       "",
		IncludeOrders:        false,
	}
	if asset, err := client.Asset(&req); err != nil {
		t.Error(err)
	} else {
		t.Log(asset.Creator)
	}
}
