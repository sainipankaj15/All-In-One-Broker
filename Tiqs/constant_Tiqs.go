package tiqs

const ADMIN_TIQS string = "FB5650"

type ProductType string

var Product = struct {
	INTRADAY, MARGIN, CNC ProductType
}{
	INTRADAY: "I",
	MARGIN:   "M",
	CNC:      "C",
}

var Index = struct {
	BANKNIFTY, NIFTY, MIDCPNIFTY, FINNIFTY, NIFTYNXT string
}{
	BANKNIFTY:  "BANKNIFTY",
	NIFTY:      "NIFTY",
	MIDCPNIFTY: "MIDCPNIFTY",
	FINNIFTY:   "FINNIFTY",
	NIFTYNXT:   "NIFTYNXT50",
}

var OptionTypes = struct {
	CALL, PUT string
}{
	CALL: "CE",
	PUT:  "PE",
}

var OrderSide = struct {
	BUY, SELL string
}{
	BUY:  "B",
	SELL: "S",
}

var PositionSide = struct {
	LONG, SHORT int
}{
	LONG:  1,
	SHORT: -1,
}

var OrderType = struct {
	LIMIT, MARKET, STOP, STOP_LIMIT string
}{
	LIMIT:      "LMT",
	MARKET:     "MKT",
	STOP:       "SL-MKT",
	STOP_LIMIT: "SL-LMT",
}

var Exchange = struct {
	NSE, NFO string
}{
	NSE: "NSE",
	NFO: "NFO",
}

var ExchangeToken = struct {
	BANKNIFTY, NIFTY50, MIDCPNIFTY, FINNIFTY, NIFTYNXT int
}{
	BANKNIFTY:  26009,
	NIFTY50:    26000,
	MIDCPNIFTY: 26074,
	FINNIFTY:   26037,
	NIFTYNXT:   26013,
}

var Lotsize = struct {
	BANKNIFTY, NIFTY50, MIDCPNIFTY, FINNIFTY int
}{
	BANKNIFTY:  15,
	NIFTY50:    25,
	MIDCPNIFTY: 50,
	FINNIFTY:   25,
}

var StrikeGap = struct {
	BANKNIFTY, NIFTY50, MIDCPNIFTY, FINNIFTY int
}{
	BANKNIFTY:  100,
	NIFTY50:    50,
	MIDCPNIFTY: 25,
	FINNIFTY:   50,
}

var apiResponseStatus = struct {
	SUCCESS, FAILURE string
}{
	SUCCESS: "success",
	FAILURE: "failure",
}
