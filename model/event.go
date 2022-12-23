package model

import "github.com/shopspring/decimal"

// EventType Describes the event type
type EventType = string

const (
	// ETCreated new auctions
	ETCreated EventType = "created"
	// ETSuccessful sales
	ETSuccessful EventType = "successful"
	// ETCancelled cancelled auctions
	ETCancelled    EventType = "cancelled"
	ETBidEntered   EventType = "bid_entered"
	ETBidWithdrawn EventType = "bid_withdrawn"
	ETTransfer     EventType = "transfer"
	ETOfferEntered EventType = "offer_entered"
	ETApprove      EventType = "approve"
)

// AuctionType Auction type
type AuctionType = string

const (
	// ATEnglish English Auctions
	ATEnglish AuctionType = "english"
	// ATDutch fixed-price and declining-price sell orders (Dutch Auctions)
	ATDutch AuctionType = "dutch"
	// ATMinPrice CryptoPunks bidding auctions
	ATMinPrice AuctionType = "min-price"
)

// Event Asset events represent state changes that occur for assets.
// This includes putting them on sale, bidding on them, selling them, cancelling sales, transferring them, and more.
type Event struct {
	ApprovedAccount *Account `opensea:"approved_account" json:"approved_account"`
	// A subfield containing a simplified version of the Asset or Asset Bundle on which this event happened
	Asset *Asset `opensea:"asset" json:"asset"`
	// Ditto
	AssetBundle     *Bundle      `opensea:"asset_bundle" json:"asset_bundle"`
	AuctionType     *AuctionType `opensea:"auction_type" json:"auction_type"`
	BidAmount       string       `opensea:"bid_amount" json:"bid_amount"`
	CollectionSlug  string       `opensea:"collection_slug" json:"collection_slug"`
	ContractAddress string       `opensea:"contract_address" json:"contract_address"`
	// When the event was recorded
	CreatedDate             string           `opensea:"created_date" json:"created_date"`
	CustomEventName         *string          `opensea:"custom_event_name" json:"custom_event_name"`
	DevFeePaymentEvent      interface{}      `opensea:"dev_fee_payment_event" json:"dev_fee_payment_event"`
	DevSellerFeeBasisPoints int              `opensea:"dev_seller_fee_basis_points" json:"dev_seller_fee_basis_points"`
	Duration                *int32           `opensea:"duration,string" json:"duration,string"`
	EndingPrice             *decimal.Decimal `opensea:"ending_price" json:"ending_price"`
	// Describes the event type
	EventType EventType `opensea:"event_type" json:"event_type"`
	// The accounts associated with this event.
	FromAccount *Account `opensea:"from_account" json:"from_account"`
	// Ditto
	ToAccount *Account `opensea:"to_account" json:"to_account"`
	ID        int64    `opensea:"id" json:"id"`
	// A boolean value that is true if the sale event was a private sale
	IsPrivate    *bool    `opensea:"is_private" json:"is_private"`
	OwnerAccount *Account `opensea:"owner_account" json:"owner_account"`
	// The payment asset used in this transaction, such as ETH, WETH or DAI
	PaymentToken *PaymentToken `opensea:"payment_token" json:"payment_token"`
	// The amount of the item that was sold. Applicable for semi-fungible assets
	Quantity      string           `opensea:"quantity" json:"quantity"`
	Seller        *Account         `opensea:"seller" json:"seller"`
	StartingPrice *decimal.Decimal `opensea:"starting_price" json:"starting_price"`
	// The total price that the asset was bought for. This includes any royalties that might have been collected
	TotalPrice    *decimal.Decimal `opensea:"total_price" json:"total_price"`
	Transaction   *Transaction     `opensea:"transaction" json:"transaction"`
	WinnerAccount *Account         `opensea:"winner_account" json:"winner_account"`
	// eg: 2017-07-21T17:32:28Z
	ListingTime *string `opensea:"listing_time" json:"listing_time"`
}
