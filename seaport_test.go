package opensea

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestSeaport(t *testing.T) {
	if approved, err := client.IsApproved(common.HexToAddress("0xdb3a7d80e41279981a4b40c9db34e584c0cac35f")); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(approved)
	}
}
