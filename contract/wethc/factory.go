package wethc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func NewWethInstance() *Weth {
	mainContractAddress := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	client, err := ethclient.Dial("https://geth.mytokenpocket.vip")
	if err != nil {
		log.Fatal(err)
	}
	instance, err := NewWeth(mainContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance
}
