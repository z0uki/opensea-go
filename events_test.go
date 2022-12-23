package opensea

import (
	"fmt"
	"testing"
)

func TestEvents(t *testing.T) {
	req := EventsRequest{
		AccountAddress: "0x57d38a1ECA9E2683978c481C81A24616E12Bdb48",
		EventType:      "collection_offer",
	}

	events, err := client.Events(&req)
	if err != nil {
		return
	}

	fmt.Println(events.AssetEvents[0])
}
