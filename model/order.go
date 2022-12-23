package model

import "github.com/shopspring/decimal"

type Side uint8

const (
	SideBuy Side = iota
	SideSell
)

// SaleKind the kind of sell order
// use only_english=true for filtering for only English Auctions
type SaleKind uint8

const (
	// SaleKindFixedOrMinBidPrice fixed-price sales or min-bid auctions
	SaleKindFixedOrMinBidPrice SaleKind = iota
	// SaleKindDecliningPrice declining-price Dutch Auctions
	SaleKindDecliningPrice
)

type HowToCall uint8

const (
	Call HowToCall = iota
	DelegateCall
)

type FeeMethod uint8

const (
	ProtocolFee FeeMethod = iota
	SplitFee
)

type AccountFees struct {
	Account     *Account `opensea:"account" json:"account"`
	BasisPoints string   `opensea:"basis_points" json:"basis_points"`
}

type MakerAssetBundle struct {
	Assets []*Asset `opensea:"assets" json:"assets"`
}

type Order struct {
	CreatedDate      string           `opensea:"created_date" json:"created_date"`
	ClosingDate      string           `opensea:"closing_date" json:"closing_date"`
	ListingTime      int              `opensea:"listing_time" json:"listing_time"`
	ExpirationTime   int              `opensea:"expiration_time" json:"expiration_time"`
	OrderHash        string           `opensea:"order_hash" json:"order_hash"`
	ProtocolData     *Protocol        `opensea:"protocol_data" json:"protocol_data"`
	ProtocolAddress  string           `opensea:"protocol_address" json:"protocol_address"`
	Maker            *Account         `opensea:"maker" json:"maker"`
	Taker            *Account         `opensea:"taker" json:"taker"`
	CurrentPrice     *decimal.Decimal `opensea:"current_price" json:"current_price"`
	MakerFees        []*AccountFees   `opensea:"maker_fees" json:"maker_fees"`
	TakerFees        []*AccountFees   `opensea:"taker_fees" json:"taker_fees"`
	Side             string           `opensea:"side" json:"side"`
	OrderType        string           `opensea:"order_type" json:"order_type"`
	Cancelled        bool             `opensea:"cancelled" json:"cancelled"`
	Finalized        bool             `opensea:"finalized" json:"finalized"`
	MarkedInvalid    bool             `opensea:"marked_invalid" json:"marked_invalid"`
	ClientSignature  string           `opensea:"client_signature" json:"client_signature"`
	RelayId          string           `opensea:"relay_id" json:"relay_id"`
	CriteriaProof    *string          `opensea:"criteria_proof" json:"criteria_proof"`
	MakerAssetBundle *Bundle          `opensea:"maker_asset_bundle" json:"maker_asset_bundle"`
	TakerAssetBundle *Bundle          `opensea:"taker_asset_bundle" json:"taker_asset_bundle"`
}

func (o *Order) Hash() string {
	if o == nil {
		return ""
	}

	return o.OrderHash
}

type Metadata struct {
	Asset           *AssetAddress `opensea:"asset" json:"asset"`
	Schema          string        `opensea:"schema" json:"schema"`
	ReferrerAddress string        `opensea:"referrerAddress" json:"referrerAddress"`
}
