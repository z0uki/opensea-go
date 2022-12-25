package model

// Asset The primary object in the OpenSea API is the asset,
// which represents a unique digital item whose ownership is managed by the blockchain.
// The below CryptoSaga hero is an example of an asset shown on OpenSea.
type Asset struct {
	// OpenSea NFT ID
	ID int `opensea:"id" json:"id"`
	// The token ID of the NFT
	TokenID  string `opensea:"token_id" json:"token_id"`
	NumSales int    `opensea:"num_sales" json:"num_sales"`
	// The background color to be displayed with the item
	BackgroundColor *string `opensea:"background_color" json:"background_color"`
	// An image for the item. Note that this is the cached URL we store on our end. The original image url is image_original_url
	ImageURL             string  `opensea:"image_url" json:"image_url"`
	ImagePreviewURL      string  `opensea:"image_preview_url" json:"image_preview_url"`
	ImageThumbnailURL    string  `opensea:"image_thumbnail_url" json:"image_thumbnail_url"`
	ImageOriginalURL     string  `opensea:"image_original_url" json:"image_original_url"`
	AnimationURL         *string `opensea:"animation_url" json:"animation_url"`
	AnimationOriginalURL *string `opensea:"animation_original_url" json:"animation_original_url"`
	// Name of the item
	Name        string  `opensea:"name" json:"name"`
	Description *string `opensea:"description" json:"description"`
	// External link to the original website for the item
	ExternalLink *string `opensea:"external_link" json:"external_link"`
	// Dictionary of data on the contract itself (see asset contract section)
	AssetContract *Contract   `opensea:"asset_contract" json:"asset_contract"`
	Permalink     string      `opensea:"permalink" json:"permalink"`
	Collection    *Collection `opensea:"collection" json:"collection"`
	Decimals      int         `opensea:"decimals" json:"decimals"`
	TokenMetadata string      `opensea:"token_metadata" json:"token_metadata"`
	// Dictionary of data on the owner (see account section)
	Owner             *Account     `opensea:"owner" json:"owner"`
	SellOrders        []*SellOrder `opensea:"sell_orders,omitempty" json:"sell_orders,omitempty"`
	SeaportSellOrders []*SellOrder `opensea:"seaport_sell_orders,omitempty" json:"seaport_sell_orders,omitempty"`
	Creator           *Account     `opensea:"creator" json:"creator"`
	// A list of traits associated with the item (see traits section)
	Traits []*Trait `opensea:"traits" json:"traits"`
	// When this item was last sold (null if there was no last sale)
	LastSale                *LastSale       `opensea:"last_sale" json:"last_sale"`
	TopBid                  interface{}     `opensea:"top_bid" json:"top_bid"`
	ListingDate             *string         `opensea:"listing_date" json:"listing_date"`
	IsPresale               bool            `opensea:"is_presale" json:"is_presale"`
	TransferFeePaymentToken *PaymentToken   `opensea:"transfer_fee_payment_token" json:"transfer_fee_payment_token"`
	TransferFee             *string         `opensea:"transfer_fee" json:"transfer_fee"`
	RelatedAssets           []*Asset        `opensea:"related_assets" json:"related_assets"`
	Orders                  []*Order        `opensea:"orders" json:"orders"`
	Auctions                []interface{}   `opensea:"auctions" json:"auctions"`
	SupportsWyvern          bool            `opensea:"supports_wyvern" json:"supports_wyvern"`
	TopOwnerships           []*TopOwnership `opensea:"top_ownerships" json:"top_ownerships"`
	Ownership               interface{}     `opensea:"ownership" json:"ownership"`
	HighestBuyerCommitment  interface{}     `opensea:"highest_buyer_commitment" json:"highest_buyer_commitment"`
}

type Transaction struct {
	BlockHash        string   `opensea:"block_hash" json:"block_hash"`
	BlockNumber      string   `opensea:"block_number" json:"block_number"`
	FromAccount      *Account `opensea:"from_account" json:"from_account"`
	ID               int      `opensea:"id" json:"id"`
	Timestamp        string   `opensea:"timestamp" json:"timestamp"`
	ToAccount        *Account `opensea:"to_account" json:"to_account"`
	TransactionHash  string   `opensea:"transaction_hash" json:"transaction_hash"`
	TransactionIndex string   `opensea:"transaction_index" json:"transaction_index"`
}

type LastSale struct {
	Asset          *AssetToken   `opensea:"asset" json:"asset"`
	AssetBundle    *Bundle       `opensea:"asset_bundle" json:"asset_bundle"`
	EventType      string        `opensea:"event_type" json:"event_type"`
	EventTimestamp string        `opensea:"event_timestamp" json:"event_timestamp"`
	AuctionType    *AuctionType  `opensea:"auction_type" json:"auction_type"`
	TotalPrice     string        `opensea:"total_price" json:"total_price"`
	PaymentToken   *PaymentToken `opensea:"payment_token" json:"payment_token"`
	Transaction    *Transaction  `opensea:"transaction" json:"transaction"`
	CreatedDate    string        `opensea:"created_date" json:"created_date"`
	Quantity       string        `opensea:"quantity" json:"quantity"`
}

type TopOwnership struct {
	Owner    *Account `opensea:"owner" json:"owner"`
	Quantity string   `opensea:"quantity" json:"quantity"`
}

type AssetAddress struct {
	ID      string `opensea:"id" json:"id"`
	Address string `opensea:"address" json:"address"`
}

type AssetToken struct {
	TokenID  string `opensea:"token_id" json:"token_id"`
	Decimals int    `opensea:"decimals" json:"decimals"`
}
