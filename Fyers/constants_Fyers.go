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
	INTRADAY, MARGIN, CNC, BO,CO string
}{
	INTRADAY: "INTRADAY",
	MARGIN:   "MARGIN",
	CNC:      "CNC",
	BO:       "BO",
	CO:       "CO",
}