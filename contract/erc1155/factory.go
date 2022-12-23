package erc1155

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func NewErc1155Instance(contractAddress string) *Contract {
	client, err := ethclient.Dial("https://geth.mytokenpocket.vip")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(contractAddress)
	instance, err := NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance
}
