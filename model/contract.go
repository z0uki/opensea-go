package model

// ContractType Types of asset contracts
// Given by the asset_contract_type in the OpenSea API
type ContractType string

const (
	Fungible     ContractType = "fungible"
	SemiFungible ContractType = "semi-fungible"
	NonFungible  ContractType = "non-fungible"
	Unknown      ContractType = "unknown"
)

// Contract Asset contracts contain data about the contract itself,
// such as the CryptoKitties contract or the CoolCats contract.
type Contract struct {
	// Address of the asset contract
	Address           string       `opensea:"address" json:"address"`
	AssetContractType ContractType `opensea:"asset_contract_type" json:"asset_contract_type"`
	CreatedDate       string       `opensea:"created_date" json:"created_date"`
	// Name of the asset contract
	Name           string  `opensea:"name" json:"name"`
	NftVersion     string  `opensea:"nft_version" json:"nft_version"`
	OpenseaVersion *string `opensea:"opensea_version" json:"opensea_version"`
	Owner          int     `opensea:"owner" json:"owner"`
	SchemaName     string  `opensea:"schema_name" json:"schema_name"`
	// Symbol, such as CKITTY
	Symbol      string `opensea:"symbol" json:"symbol"`
	TotalSupply string `opensea:"total_supply" json:"total_supply"`
	// Description of the asset contract
	Description string `opensea:"description" json:"description"`
	// Link to the original website for this contract
	ExternalLink string `opensea:"external_link" json:"external_link"`
	// Image associated with the asset contract
	ImageURL                    string      `opensea:"image_url" json:"image_url"`
	DefaultToFiat               bool        `opensea:"default_to_fiat" json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `opensea:"dev_buyer_fee_basis_points" json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `opensea:"dev_seller_fee_basis_points" json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `opensea:"only_proxied_transfers" json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int         `opensea:"opensea_buyer_fee_basis_points" json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `opensea:"opensea_seller_fee_basis_points" json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `opensea:"buyer_fee_basis_points" json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `opensea:"seller_fee_basis_points" json:"seller_fee_basis_points"`
	PayoutAddress               string      `opensea:"payout_address" json:"payout_address"`
	Collection                  *Collection `opensea:"collection,omitempty" json:"collection"`
}

func (c *Contract) ContractAddress() string {
	if c == nil {
		return ""
	}

	return c.Address
}
