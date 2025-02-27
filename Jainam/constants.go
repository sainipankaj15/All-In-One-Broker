package jainam

const ADMIN_TIQS string = "DK2100311"

type ProductType string

var Product = struct {
	INTRADAY, MARGIN, CNC ProductType
}{
	INTRADAY: "MIS",
	MARGIN:   "NRML",
	CNC:      "CNC",
}

var Index = struct {
	BANKNIFTY, NIFTY, MIDCPNIFTY, FINNIFTY, NIFTYNXT, SENSEX string
}{
	BANKNIFTY:  "BANKNIFTY",
	NIFTY:      "NIFTY",
	MIDCPNIFTY: "MIDCPNIFTY",
	FINNIFTY:   "FINNIFTY",
	NIFTYNXT:   "NIFTYNXT50",
	SENSEX:     "SENSEX",
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
	LIMIT:      "L",
	MARKET:     "MKT",
	STOP:       "SL-M",
	STOP_LIMIT: "SL",
}

var Exchange = struct {
	NSE, NFO, BFO string
}{
	NSE: "NSE",
	NFO: "NFO",
	BFO: "BFO",
}

var ExchangeToken = struct {
	BANKNIFTY, NIFTY50, MIDCPNIFTY, FINNIFTY, NIFTYNXT, SENSEX int
}{
	BANKNIFTY:  26009,
	NIFTY50:    26000,
	MIDCPNIFTY: 26074,
	FINNIFTY:   26037,
	NIFTYNXT:   26013,
	SENSEX:     999001,
}

var Lotsize = struct {
	BANKNIFTY, NIFTY50, MIDCPNIFTY, FINNIFTY, SENSEX int
}{
	BANKNIFTY:  30,
	NIFTY50:    75,
	MIDCPNIFTY: 120,
	FINNIFTY:   65,
	SENSEX:     20,
}

var StrikeGap = struct {
	BANKNIFTY, NIFTY50, MIDCPNIFTY, FINNIFTY, SENSEX int
}{
	BANKNIFTY:  100,
	NIFTY50:    50,
	MIDCPNIFTY: 25,
	FINNIFTY:   50,
	SENSEX:     100,
}

var apiResponseStatus = struct {
	SUCCESS, FAILURE string
}{
	SUCCESS: "success",
	FAILURE: "failure",
}

// Tiqs Base URL
var baseURL = "https://protrade.jainam.in"

var positionUrl = baseURL + "/api/po-rest/positions	"
var orderBookURL = baseURL + "/api/od-rest/info/orderbook	"
var tradeBookURL = baseURL + "/api/od-rest/info/tradebook	"
var placeOrderUrl = baseURL + "/api/od-rest/orders/execute"

type orderStatus string

var OrderStatues = struct {
	OPEN, CANCELED, COMPLETE, REJECTED, PENDING orderStatus
}{
	OPEN:     "OPEN",
	CANCELED: "CANCELED",
	COMPLETE: "COMPLETE",
	REJECTED: "REJECTED",
	PENDING:  "PENDING",
}
