package xts

const ADMIN_XTS string = "DK2100311"

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
	BUY:  "BUY",
	SELL: "SELL",
}

var PositionSide = struct {
	LONG, SHORT int
}{
	LONG:  1,
	SHORT: -1,
}

var PriceType = struct {
	LIMIT, MARKET, STOP, STOP_LIMIT string
}{
	LIMIT:      "LIMIT",
	MARKET:     "MARKET",
	STOP_LIMIT: "STOPLIMIT",
	STOP:       "STOPMARKET",
}

var Exchange = struct {
	NSECM, NSEFO, BSECM, BSEFO string
}{
	NSECM: "NSECM",
	NSEFO: "NSEFO",
	BSECM: "BSECM",
	BSEFO: "BSEFO",
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
	SUCCESS: "Ok",
	FAILURE: "failure",
}

// Tiqs Base URL
var baseURL = "https://jtrade.jainam.in"

var placeOrderUrl = baseURL + "/interactive/orders"

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
	REGULAR, BO, CO, AMO string
}{
	REGULAR: "Regular",
	BO:      "Bracket",
	CO:      "Cover",
	AMO:     "AMO",
}
