package erc721

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func NewErc721Instance(contractAddress common.Address) *Contract {
	client, err := ethclient.Dial("https://geth.mytokenpocket.vip")
	if err != nil {
		log.Fatal(err)
	}
	instance, err := NewContract(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance
}
