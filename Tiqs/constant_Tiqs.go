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
	LIMIT:      "LMT",
	MARKET:     "MKT",
	STOP:       "SL-MKT",
	STOP_LIMIT: "SL-LMT",
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
var baseURL = "https://api.tiqs.trading"

var quotesUrl = baseURL + "/info/quotes/full"
var positionUrl = baseURL + "/user/positions"
var orderBookURL = baseURL + "/order"
var placeOrderUrl = baseURL + "/order/regular"
var getOptionChainUrl = baseURL + "/info/option-chain"
var expiryDayListUrl = baseURL + "/info/option-chain-symbols"
var ltpUrl = baseURL + "/info/quote/ltp"
var greeksUrl = baseURL + "/info/greeks"
var holidaysUrl = baseURL + "/info/holidays"

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
