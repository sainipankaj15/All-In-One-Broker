package tiqs_greeks_socket

import (
	"github.com/alphadose/haxmap"
	tiqs "github.com/sainipankaj15/All-In-One-Broker/Tiqs"
)

type TiqsGreeksClient struct {
	appID                   string
	accessToken             string
	enableLog               bool
	timeToExpireInDays      int
	priceMap                *haxmap.Map[int32, TickData]
	optionChain             map[string]map[string]tiqs.Symbol
	strikeToSyntheticFuture *haxmap.Map[int32, float64]
	peTokenToCeToken        *haxmap.Map[int32, int32]
}

type TickData struct {
	LTP         int32
	Timestamp   int32
	StrikePrice int32
	OptionType  string
	Delta       float64
	Theta       float64
	Vega        float64
	Gamma       float64
	IV          float64
}
