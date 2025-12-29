package zerodha

// PlaceOrderResponse represents the response from the place order API.
// It contains the status, data about the order, a message describing the result, and an error type if the order was unsuccessful.
type PlaceOrderResponse struct {
	Status    string    `json:"status"`
	Data      OrderData `json:"data"`       // data about the order
	Message   string    `json:"message"`    // message describing the result
	ErrorType string    `json:"error_type"` // error type if the order was unsuccessful
}

type OrderData struct {
	OrderID string `json:"order_id"`
}

type readDataJsonZerodha struct {
	Date        string `json:"Date"`
	ApiKey      string `json:"apiKey"`
	AccessToken string `json:"token"`
	UserID      string `json:"userID"`
	Name        string `json:"name"`
}

type FundsResponse struct {
	Status string    `json:"status"`
	Data   FundsData `json:"data"`
}

// FundsData represents the funds information for equity and commodity segments.
type FundsData struct {
	Equity    FundSegment `json:"equity"`
	Commodity FundSegment `json:"commodity"`
}

type FundSegment struct {
	Enabled   bool          `json:"enabled"`
	Net       float64       `json:"net"`
	Available FundAvailable `json:"available"`
	Utilised  FundUtilised  `json:"utilised"`
}

type FundAvailable struct {
	AdhocMargin    float64 `json:"adhoc_margin"`
	Cash           float64 `json:"cash"`
	OpeningBalance float64 `json:"opening_balance"`
	LiveBalance    float64 `json:"live_balance"`
	Collateral     float64 `json:"collateral"`
	IntradayPayin  float64 `json:"intraday_payin"`
}

type FundUtilised struct {
	Debits           float64 `json:"debits"`
	Exposure         float64 `json:"exposure"`
	M2MRealised      float64 `json:"m2m_realised"`
	M2MUnrealised    float64 `json:"m2m_unrealised"`
	OptionPremium    float64 `json:"option_premium"`
	Payout           float64 `json:"payout"`
	Span             float64 `json:"span"`
	HoldingSales     float64 `json:"holding_sales"`
	Turnover         float64 `json:"turnover"`
	LiquidCollateral float64 `json:"liquid_collateral"`
	StockCollateral  float64 `json:"stock_collateral"`
	Delivery         float64 `json:"delivery"`
}

// HoldingsResponse represents the response from the holdings API for Zerodha.
// It contains the status and a list of holding items.
type HoldingsResponse struct {
	// Status is the status of the API response.
	Status string `json:"status"`
	// Data is a list of holding items.
	Data []HoldingItem `json:"data"`
}

type HoldingItem struct {
	TradingSymbol       string      `json:"tradingsymbol"`
	Exchange            string      `json:"exchange"`
	InstrumentToken     uint32      `json:"instrument_token"`
	ISIN                string      `json:"isin"`
	Product             string      `json:"product"`
	Price               float64     `json:"price"`
	Quantity            int64       `json:"quantity"`
	UsedQuantity        int64       `json:"used_quantity"`
	T1Quantity          int64       `json:"t1_quantity"`
	RealisedQuantity    int64       `json:"realised_quantity"`
	AuthorisedQuantity  int64       `json:"authorised_quantity"`
	AuthorisedDate      string      `json:"authorised_date"`
	Authorisation       interface{} `json:"authorisation"`
	OpeningQuantity     int64       `json:"opening_quantity"`
	ShortQuantity       int64       `json:"short_quantity"`
	CollateralQuantity  int64       `json:"collateral_quantity"`
	CollateralType      string      `json:"collateral_type"`
	Discrepancy         bool        `json:"discrepancy"`
	AveragePrice        float64     `json:"average_price"`
	LastPrice           float64     `json:"last_price"`
	ClosePrice          float64     `json:"close_price"`
	PnL                 float64     `json:"pnl"`
	DayChange           float64     `json:"day_change"`
	DayChangePercentage float64     `json:"day_change_percentage"`
	MTF                 HoldingMTF  `json:"mtf"`
}

type HoldingMTF struct {
	Quantity      int64   `json:"quantity"`
	UsedQuantity  int64   `json:"used_quantity"`
	AveragePrice  float64 `json:"average_price"`
	Value         float64 `json:"value"`
	InitialMargin float64 `json:"initial_margin"`
}
