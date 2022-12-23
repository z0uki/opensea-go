package opensea

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/z0uki/opensea-go/contract/erc721"
)

var (
	SEAPORT_CONTRACT_ERC721 = common.HexToAddress("0x1E0049783F008A0085193E00003D00cd54003c71")
)

// IsApproved check if the contract is approved
func (c *Client) IsApproved(operator common.Address) (bool, error) {
	erc721Instance := erc721.NewErc721Instance(operator)
	return erc721Instance.IsApprovedForAll(nil, c.wallet.Address, SEAPORT_CONTRACT_ERC721)
}

// SetApprovalForAll set approval for all
func (c *Client) SetApprovalForAll(operator common.Address) (*types.Receipt, error) {
	erc721Instance := erc721.NewErc721Instance(operator)

	Opts, err := bind.NewKeyedTransactorWithChainID(c.wallet.PrivateKey, big.NewInt(1))
	if err != nil {
		return nil, err
	}

	tx, err := erc721Instance.SetApprovalForAll(Opts, SEAPORT_CONTRACT_ERC721, true)
	if err != nil {
		return nil, err
	}

	return c.WaitMined(tx)
}

// WaitMined watch tx status
func (c *Client) WaitMined(tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(context.Background(), c.eclient, tx)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}
