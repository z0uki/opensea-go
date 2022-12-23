package opensea

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/pkg/crypto/eip712"
)

const (
	WethAddress         = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	OpenSeaFeeRecipient = "0x0000a26b00c1F0DF003000390027140000fAa719"
	Zone                = "0x004C00500000aD104D7DBd00e3ae0A5C00560C00"
	ZoneHash            = "0x0000000000000000000000000000000000000000000000000000000000000000"
	ConduitKey          = "0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000"
	ZeroAddress         = "0x0000000000000000000000000000000000000000"
)

const (
	ItemType_NATIVE                uint8 = 0
	ItemType_ERC20                 uint8 = 1
	ItemType_ERC721                uint8 = 2
	ItemType_ERC1155               uint8 = 3
	ItemType_ERC721_WITH_CRITERIA  uint8 = 4
	ItemType_ERC1155_WITH_CRITERIA uint8 = 5
)

type SellType int64

type Wallet struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

var TYPES = eip712.Types{
	"EIP712Domain": {
		{Name: "name", Type: "string"},
		{Name: "version", Type: "string"},
		{Name: "chainId", Type: "uint256"},
		{Name: "verifyingContract", Type: "address"},
	},
	"OrderComponents": {
		{
			Name: "offerer",
			Type: "address",
		},
		{
			Name: "zone",
			Type: "address",
		},
		{
			Name: "offer",
			Type: "OfferItem[]",
		},
		{
			Name: "consideration",
			Type: "ConsiderationItem[]",
		},
		{
			Name: "orderType",
			Type: "uint8",
		},
		{
			Name: "startTime",
			Type: "uint256",
		},
		{
			Name: "endTime",
			Type: "uint256",
		},
		{
			Name: "zoneHash",
			Type: "bytes32",
		},
		{
			Name: "salt",
			Type: "uint256",
		},
		{
			Name: "conduitKey",
			Type: "bytes32",
		},
		{
			Name: "counter",
			Type: "uint256",
		},
	},
	"OfferItem": {
		{
			Name: "itemType",
			Type: "uint8",
		},
		{
			Name: "token",
			Type: "address",
		},
		{
			Name: "identifierOrCriteria",
			Type: "uint256",
		},
		{
			Name: "startAmount",
			Type: "uint256",
		},
		{
			Name: "endAmount",
			Type: "uint256",
		},
	},
	"ConsiderationItem": {
		{
			Name: "itemType",
			Type: "uint8",
		},
		{
			Name: "token",
			Type: "address",
		},
		{
			Name: "identifierOrCriteria",
			Type: "uint256",
		},
		{
			Name: "startAmount",
			Type: "uint256",
		},
		{
			Name: "endAmount",
			Type: "uint256",
		},
		{
			Name: "recipient",
			Type: "address",
		},
	},
}
