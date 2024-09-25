package tiqs_greeks_socket

import (
	"github.com/alphadose/haxmap"
	tiqs "github.com/sainipankaj15/All-In-One-Broker/Tiqs"
)

type TiqsGreeksClient struct {
	appID                   string
	accessToken             string
	enableLog               bool
	priceMap                *haxmap.Map[int32, TickData]
	optionChain             map[string]map[string]tiqs.Symbol
	strikeToSyntheticFuture *haxmap.Map[int32, float64]
}

type TickData struct {
	LTP         int32
	Timestamp   int32
	StrikePrice int32
	OptionType  string
}
