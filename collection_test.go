package opensea

import (
	"testing"
)

func TestCollection(t *testing.T) {
	req := CollectionRequest{
		CollectionSlug: "playerone-1p",
	}

	if collection, err := client.Collection(&req); err != nil {
		t.Error(err)
	} else {
		t.Log(int64(collection.Stats.FloorPrice * 1000000000000000000))
	}
}
