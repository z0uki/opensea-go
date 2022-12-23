package erc721

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"testing"
)

func TestErc721(t *testing.T) {
	client, err := ethclient.Dial("https://geth.mytokenpocket.vip")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x2822b8a0d457209a3146986bfac5d675260f2c7e")
	instance, err := NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}

	all, err := instance.IsApprovedForAll(nil, common.HexToAddress("0x8fE3e0C9b2B86D7cB33bfC4C09c6e748A37F82de"), common.HexToAddress("0xf849de01b080adc3a814fabe1e2087475cf2e354"))
	if err != nil {
		return
	}

	fmt.Println(all)
}
