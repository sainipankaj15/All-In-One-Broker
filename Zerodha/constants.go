package zerodha

const ADMIN_TIQS string = "ZM7200"

var Index = struct {
	BANKNIFTY, NIFTY, MIDCPNIFTY, FINNIFTY, NIFTYNXT, SENSEX string
}{
	BANKNIFTY:  "NIFTY BANK",
	NIFTY:      "NIFTY 50",
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

var TransactionSide = struct {
	BUY, SELL string
}{
	BUY:  "BUY",
	SELL: "SELL",
}

var PositionSide = struct {
	LONG, SHORT int
}{
	LONG:  1,
	SHORT: -1,
}

var ExchangeToken = struct {
	BANKNIFTY, NIFTY50, MIDCPNIFTY, FINNIFTY, NIFTYNXT, SENSEX int
}{
	BANKNIFTY:  260105,
	NIFTY50:    256265,
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

// Zerodha Base URL
var baseURL = "https://api.kite.trade"

var placeOrderUrl = baseURL + "/orders/regular"
var positionsURL = baseURL + "/portfolio/positions"
var holdingsURL = baseURL + "/portfolio/holdings"
var marginsURL = baseURL + "/user/margins"

type orderStatus string

var OrderStatuses = struct {
	OPEN, CANCELED, COMPLETE, REJECTED, PENDING orderStatus
}{
	OPEN:     "OPEN",
	CANCELED: "CANCELED",
	COMPLETE: "COMPLETE",
	REJECTED: "REJECTED",
	PENDING:  "PENDING",
}

var OrderType = struct {
	LIMIT, MARKET, STOPLOSS, STOPLOSS_MARKET string
}{
	LIMIT:           "LIMIT",
	MARKET:          "MARKET",
	STOPLOSS:        "SL",
	STOPLOSS_MARKET: "SL-M",
}

var ProductType = struct {
	INTRADAY, NORMAL, CNC string
}{
	CNC:      "CNC",
	NORMAL:   "NRML", //Normal for futures and options
	INTRADAY: "MIS",  // Margin Intraday Squareoff for futures and options
}

var Exchange = struct {
	NSE, BSE, NFO, CDS, BCD, MCX string
}{
	NSE: "NSE",
	BSE: "BSE",
	NFO: "NFO",
	CDS: "CDS",
	BCD: "BCD",
	MCX: "MCX",
}
