package model

import (
	"github.com/shopspring/decimal"
	"math/big"
)

// Collection Collections are used to represent all the assets in a single (or multiple) contract addresses
// and help users group items from the same creator. They have one or more owners and are typically
// associated with important metadata such as creator royalties and descriptions.
// Visit it(https://docs.opensea.io/reference/collection-model) to learn anymore.
type Collection struct {
	// The collection name. Typically derived from the first contract imported to the collection but can be changed by the user
	Name string `opensea:"name" json:"name"`
	// Description for the model
	Description      string  `opensea:"description" json:"description"`
	ShortDescription *string `opensea:"short_description" json:"short_description"`
	// The collection slug that is used to link to the collection on OpenSea.
	// This value can change by the owner but must be unique across all collection slugs in OpenSea
	Slug string `opensea:"slug" json:"slug"`
	// External link to the original website for the collection
	ExternalURL string `opensea:"external_url" json:"external_url"`
	// An image for the collection. Note that this is the cached URL we store on our end.
	// The original image url is image_original_url
	ImageURL      string  `opensea:"image_url" json:"image_url"`
	LargeImageURL *string `opensea:"large_image_url" json:"large_image_url"`
	// Approved editors on this collection.
	Editors []string `opensea:"editors" json:"editors"`
	// The payment tokens accepted for this collection
	PaymentTokens []*PaymentToken `opensea:"payment_tokens" json:"payment_tokens"`
	// A list of the contracts that are associated with this collection
	PrimaryAssetContracts []*Contract `opensea:"primary_asset_contracts" json:"primary_asset_contracts"`
	// A dictionary listing all the trait types available within this collection
	Traits map[string]map[string]float64 `opensea:"traits" json:"traits"`
	// A dictionary containing some sales statistics related to this collection, including trade volume and floor prices
	Stats *CollectionStats `opensea:"stats" json:"stats"`
	// Image used in the horizontal top banner for the collection.
	BannerImageURL string `opensea:"banner_image_url" json:"banner_image_url"`
	// The payout address for the collection's royalties
	PayoutAddress string `opensea:"payout_address" json:"payout_address"`
	// The collector's fees that get paid out to them when sales are made for their collections
	DevSellerFeeBasisPoints string `opensea:"dev_seller_fee_basis_points" json:"dev_seller_fee_basis_points"`
	// The collection's approval status within OpenSea.
	// Can be not_requested (brand new collections), requested (collections that requested safelisting on our site),
	// approved (collections that are approved on our site and can be found in search results),
	// and verified (verified collections)
	SafelistRequestStatus       string      `opensea:"safelist_request_status" json:"safelist_request_status"`
	Fees                        Fees        `opensea:"fees" json:"fees"`
	IsCreatorFeesEnforced       bool        `opensea:"is_creator_fees_enforced" json:"is_creator_fees_enforced"`
	CreatedDate                 string      `opensea:"created_date" json:"created_date"`
	DefaultToFiat               bool        `opensea:"default_to_fiat" json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      string      `opensea:"dev_buyer_fee_basis_points" json:"dev_buyer_fee_basis_points"`
	DisplayData                 DisplayData `opensea:"display_data" json:"display_data"`
	Featured                    bool        `opensea:"featured" json:"featured"`
	FeaturedImageURL            *string     `opensea:"featured_image_url" json:"featured_image_url"`
	Hidden                      bool        `opensea:"hidden" json:"hidden"`
	IsSubjectToWhitelist        bool        `opensea:"is_subject_to_whitelist" json:"is_subject_to_whitelist"`
	OnlyProxiedTransfers        bool        `opensea:"only_proxied_transfers" json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string      `opensea:"opensea_buyer_fee_basis_points" json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `opensea:"opensea_seller_fee_basis_points" json:"opensea_seller_fee_basis_points"`
	RequireEmail                bool        `opensea:"require_email" json:"require_email"`

	TwitterUsername   *string `opensea:"twitter_username" json:"twitter_username"`
	InstagramUsername string  `opensea:"instagram_username" json:"instagram_username"`
	MediumUsername    *string `opensea:"medium_username" json:"medium_username"`

	TelegramURL *string `opensea:"telegram_url" json:"telegram_url"`
	DiscordURL  string  `opensea:"discord_url" json:"discord_url"`
	ChatURL     *string `opensea:"chat_url" json:"chat_url"`
	WikiURL     *string `opensea:"wiki_url" json:"wiki_url"`
}

func (c *Collection) CollectionSlug() string {
	if c == nil {
		return ""
	}

	return c.Slug
}

type Fee = map[string]*big.Int

type Fees struct {
	SellerFees  *Fee `opensea:"seller_fees" json:"seller_fees"`
	OpenseaFees *Fee `opensea:"opensea_fees" json:"opensea_fees"`
}

type PaymentToken struct {
	ID       int              `opensea:"id" json:"id"`
	Symbol   string           `opensea:"symbol" json:"symbol"`
	Address  string           `opensea:"address" json:"address"`
	ImageURL string           `opensea:"image_url" json:"image_url"`
	Name     string           `opensea:"name" json:"name"`
	Decimals int              `opensea:"decimals" json:"decimals"`
	EthPrice *decimal.Decimal `opensea:"eth_price" json:"eth_price"`
	UsdPrice *decimal.Decimal `opensea:"usd_price" json:"usd_price"`
}

type CollectionStats struct {
	OneDayVolume          float64 `opensea:"one_day_volume" json:"one_day_volume"`
	OneDayChange          float64 `opensea:"one_day_change" json:"one_day_change"`
	OneDaySales           float64 `opensea:"one_day_sales" json:"one_day_sales"`
	OneDayAveragePrice    float64 `opensea:"one_day_average_price" json:"one_day_average_price"`
	SevenDayVolume        float64 `opensea:"seven_day_volume" json:"seven_day_volume"`
	SevenDayChange        float64 `opensea:"seven_day_change" json:"seven_day_change"`
	SevenDaySales         float64 `opensea:"seven_day_sales" json:"seven_day_sales"`
	SevenDayAveragePrice  float64 `opensea:"seven_day_average_price" json:"seven_day_average_price"`
	ThirtyDayVolume       float64 `opensea:"thirty_day_volume" json:"thirty_day_volume"`
	ThirtyDayChange       float64 `opensea:"thirty_day_change" json:"thirty_day_change"`
	ThirtyDaySales        float64 `opensea:"thirty_day_sales" json:"thirty_day_sales"`
	ThirtyDayAveragePrice float64 `opensea:"thirty_day_average_price" json:"thirty_day_average_price"`
	TotalVolume           float64 `opensea:"total_volume" json:"total_volume"`
	TotalSales            float64 `opensea:"total_sales" json:"total_sales"`
	TotalSupply           float64 `opensea:"total_supply" json:"total_supply"`
	Count                 float64 `opensea:"count" json:"count"`
	NumOwners             int     `opensea:"num_owners" json:"num_owners"`
	AveragePrice          float64 `opensea:"average_price" json:"average_price"`
	NumReports            int     `opensea:"num_reports" json:"num_reports"`
	MarketCap             float64 `opensea:"market_cap" json:"market_cap"`
	FloorPrice            float64 `opensea:"floor_price" json:"floor_price"`
}

type DisplayData struct {
	CardDisplayStyle string `opensea:"card_display_style" json:"card_display_style"`
}
