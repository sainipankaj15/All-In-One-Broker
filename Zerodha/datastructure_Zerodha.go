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
