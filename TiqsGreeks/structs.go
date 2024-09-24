package tiqs_greeks_socket

import (
	"github.com/alphadose/haxmap"
)

type TiqsGreeksClient struct {
	appID       string
	accessToken string
	enableLog   bool
	priceMap    *haxmap.Map[int32, TickData]
}

type TickData struct {
	LTP       int32
	Timestamp int32
}
