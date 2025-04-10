package fyers

const ADMIN_FYERS string = "XP03754"

var OptionTypes = struct {
	CALL, PUT string
}{
	CALL: "CE",
	PUT:  "PE",
}

var OrderSide = struct {
	BUY, SELL int
}{
	BUY:  1,
	SELL: -1,
}

var PositionSide = struct {
	LONG, SHORT int
}{
	LONG:  1,
	SHORT: -1,
}

var OrderType = struct {
	LIMIT, MARKET, STOP, STOP_LIMIT int
}{
	LIMIT:      1,
	MARKET:     2,
	STOP:       3,
	STOP_LIMIT: 4,
}

var ProductType = struct {
	INTRADAY, MARGIN, CNC, BO, CO string
}{
	INTRADAY: "INTRADAY",
	MARGIN:   "MARGIN",
	CNC:      "CNC",
	BO:       "BO",
	CO:       "CO",
}

var Lotsize = struct {
	BANKNIFTY, NIFTY50, MIDCPNIFTY, FINNIFTY int
}{
	BANKNIFTY:  30,
	NIFTY50:    75,
	MIDCPNIFTY: 120,
	FINNIFTY:   65,
}

var StrikeGap = struct {
	BANKNIFTY, NIFTY50, MIDCPNIFTY, FINNIFTY int
}{
	BANKNIFTY:  100,
	NIFTY50:    50,
	MIDCPNIFTY: 25,
	FINNIFTY:   50,
}

var Index = struct {
	BANKNIFTY, NIFTY, MIDCPNIFTY, FINNIFTY, NIFTYNXT, SENSEX string
}{
	BANKNIFTY:  "NSE:NIFTYBANK-INDEX",
	NIFTY:      "NSE:NIFTY50-INDEX",
	MIDCPNIFTY: "NSE:MIDCPNIFTY-INDEX",
	FINNIFTY:   "NSE:FINNIFTY-INDEX",
	NIFTYNXT:   "NSE:NIFTYNXT50-INDEX",
	SENSEX:     "BSE:SENSEX-INDEX",
}
