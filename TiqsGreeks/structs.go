package tiqs_greeks_socket

import (
	"github.com/alphadose/haxmap"
)

type TiqsGreeksClient struct {
	appID                   string
	accessToken             string
	enableLog               bool
	priceMap                *haxmap.Map[int32, TickData]
	strikeToSyntheticFuture *haxmap.Map[int32, float64]
}

type TickData struct {
	LTP         int32
	Timestamp   int32
	StrikePrice int32
	OptionType  string
}
