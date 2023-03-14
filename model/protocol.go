package model

type Protocol struct {
	Parameters Parameters `opensea:"parameters" json:"parameters"`
	Signature  string     `opensea:"signature" json:"signature"`
}

type Parameters struct {
	Offerer                         string              `opensea:"offerer" json:"offerer"`
	Offer                           []OfferItem         `opensea:"offer" json:"offer"`
	Consideration                   []ConsiderationItem `opensea:"consideration" json:"consideration"`
	StartTime                       string              `opensea:"startTime" json:"startTime"`
	EndTime                         string              `opensea:"endTime" json:"endTime"`
	OrderType                       int                 `opensea:"orderType" json:"orderType"`
	Zone                            string              `opensea:"zone" json:"zone"`
	ZoneHash                        string              `opensea:"zoneHash" json:"zoneHash"`
	Salt                            string              `opensea:"salt" json:"salt"`
	ConduitKey                      string              `opensea:"conduitKey" json:"conduitKey"`
	TotalOriginalConsiderationItems int                 `opensea:"totalOriginalConsiderationItems" json:"totalOriginalConsiderationItems"`
	Counter                         interface{}         `opensea:"counter" json:"counter"`
}

type OfferItem struct {
	ItemType             uint8  `opensea:"itemType" json:"itemType"`
	Token                string `opensea:"token" json:"token"`
	IdentifierOrCriteria string `opensea:"identifierOrCriteria" json:"identifierOrCriteria"`
	StartAmount          string `opensea:"startAmount" json:"startAmount"`
	EndAmount            string `opensea:"endAmount" json:"endAmount"`
}

type ConsiderationItem struct {
	ItemType             uint8  `opensea:"itemType" json:"itemType"`
	Token                string `opensea:"token" json:"token"`
	IdentifierOrCriteria string `opensea:"identifierOrCriteria" json:"identifierOrCriteria"`
	StartAmount          string `opensea:"startAmount" json:"startAmount"`
	EndAmount            string `opensea:"endAmount" json:"endAmount"`
	Recipient            string `opensea:"recipient" json:"recipient"`
}

//type Protocol struct {
//	Parameters Parameters `opensea:"parameters" json:"parameters"`
//	Signature  string     `opensea:"signature" json:"signature"`
//}
//
//type Parameters struct {
//	Offerer                         string              `opensea:"offerer" json:"offerer"`
//	Offer                           []OfferItem         `opensea:"offer" json:"offer"`
//	Consideration                   []ConsiderationItem `opensea:"consideration" json:"consideration"`
//	StartTime                       int64               `opensea:"startTime" json:"startTime"`
//	EndTime                         int64               `opensea:"endTime" json:"endTime"`
//	OrderType                       int                 `opensea:"orderType" json:"orderType"`
//	Zone                            string              `opensea:"zone" json:"zone"`
//	ZoneHash                        string              `opensea:"zoneHash" json:"zoneHash"`
//	Salt                            int64               `opensea:"salt" json:"salt"`
//	ConduitKey                      string              `opensea:"conduitKey" json:"conduitKey"`
//	TotalOriginalConsiderationItems int                 `opensea:"totalOriginalConsiderationItems" json:"totalOriginalConsiderationItems"`
//	Counter                         int                 `opensea:"counter" json:"counter"`
//}
//
//type OfferItem struct {
//	ItemType             uint8  `opensea:"itemType" json:"itemType"`
//	Token                string `opensea:"token" json:"token"`
//	IdentifierOrCriteria int64  `opensea:"identifierOrCriteria" json:"identifierOrCriteria"`
//	StartAmount          int64  `opensea:"startAmount" json:"startAmount"`
//	EndAmount            int64  `opensea:"endAmount" json:"endAmount"`
//}
//
//type ConsiderationItem struct {
//	ItemType             uint8  `opensea:"itemType" json:"itemType"`
//	Token                string `opensea:"token" json:"token"`
//	IdentifierOrCriteria string `opensea:"identifierOrCriteria" json:"identifierOrCriteria"`
//	StartAmount          int64  `opensea:"startAmount" json:"startAmount"`
//	EndAmount            int64  `opensea:"endAmount" json:"endAmount"`
//	Recipient            string `opensea:"recipient" json:"recipient"`
//}
