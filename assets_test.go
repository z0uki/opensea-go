package opensea

import (
	"fmt"
	"testing"
)

func TestAssets(t *testing.T) {
	assetsReq := AssetsRequest{
		Owner:         "0x57d38a1ECA9E2683978c481C81A24616E12Bdb48",
		Limit:         200,
		IncludeOrders: true,
	}

	if assets, err := client.Assets(&assetsReq); err != nil {
		t.Error(err)
	} else {
		for _, asset := range assets.Assets {
			fmt.Println(asset.SeaportSellOrders[0])
		}
	}
}
