package opensea

import "testing"

func TestCollection(t *testing.T) {
	req := CollectionRequest{
		CollectionSlug: "dragonkitty-v3",
	}

	if collection, err := client.Collection(&req); err != nil {
		t.Error(err)
	} else {
		t.Log(collection.PrimaryAssetContracts[0].SchemaName)
	}
}
