package fyers

type ReadDataJson_Fyers struct {
	Date                 string `json:"Date"`
	AccessToken          string `json:"access_token"`
	AccessTokenWithAppID string `json:"access_token_with_APPID"`
	AppId                string `json:"app_id"`
	AppIdWithType        string `json:"app_id_with_app_type"`
}

type PositionResponse struct {
	S            string          `json:"s"`
	Code         int             `json:"code"`
	Message      string          `json:"message"`
	NetPositions []NetPosition   `json:"netPositions"`
	Overall      PositionOverall `json:"overall"`
}

type PositionOverall struct {
	CountTotal   int     `json:"count_total"`
	CountOpen    int     `json:"count_open"`
	PLTotal      float64 `json:"pl_total"`
	PLRealized   float64 `json:"pl_realized"`
	PLUnrealized float64 `json:"pl_unrealized"`
}

type NetPosition struct {
	Symbol string `json:"symbol"`
	ID     string `json:"id"`

	BuyQty int     `json:"buyQty"`
	BuyAvg float64 `json:"buyAvg"`
	BuyVal float64 `json:"buyVal"`

	SellQty int     `json:"sellQty"`
	SellAvg float64 `json:"sellAvg"`
	SellVal float64 `json:"sellVal"`

	NetQty   int     `json:"netQty"`
	Qty      int     `json:"qty"`
	AvgPrice float64 `json:"avgPrice"`
	NetAvg   float64 `json:"netAvg"`
	Side     int     `json:"side"`

	ProductType string `json:"productType"`

	PL           float64 `json:"pl"`
	PLRealized   float64 `json:"realized_profit"`
	PLUnrealized float64 `json:"unrealized_profit"`

	LTP float64 `json:"ltp"`

	FYToken       string  `json:"fyToken"`
	CrossCurrency string  `json:"crossCurrency"`
	RBIRefRate    float64 `json:"rbiRefRate"`
	QtyMultiplier float64 `json:"qtyMulti_com"`

	Segment  int `json:"segment"`
	Exchange int `json:"exchange"`
	SlNo     int `json:"slNo"`

	CFBuyQty   int `json:"cfBuyQty"`
	CFSellQty  int `json:"cfSellQty"`
	DayBuyQty  int `json:"dayBuyQty"`
	DaySellQty int `json:"daySellQty"`
}

type MarketDepthAPIResp_Fyers struct {
	D       map[string]StockData `json:"d"`
	Message string               `json:"message"`
	Status  string               `json:"s"`
}

type StockData struct {
	H   float64 `json:"h"`
	L   float64 `json:"l"`
	Ltp float64 `json:"ltp"`
}

type QuoteAPIResp_Fyers struct {
	Code int          `json:"code"`
	S    string       `json:"s"`
	D    []StockEntry `json:"d"`
}

type StockEntry struct {
	N string     `json:"n"`
	S string     `json:"s"`
	V StockValue `json:"v"`
}

type StockValue struct {
	Ch             float64 `json:"ch"`
	Chp            float64 `json:"chp"`
	Lp             float64 `json:"lp"`
	Spread         float64 `json:"spread"`
	Ask            float64 `json:"ask"`
	Bid            float64 `json:"bid"`
	OpenPrice      float64 `json:"open_price"`
	HighPrice      float64 `json:"high_price"`
	LowPrice       float64 `json:"low_price"`
	PrevClosePrice float64 `json:"prev_close_price"`
	Volume         int     `json:"volume"`
	ShortName      string  `json:"short_name"`
	Exchange       string  `json:"exchange"`
	Description    string  `json:"description"`
	OriginalName   string  `json:"original_name"`
	Symbol         string  `json:"symbol"`
	FyToken        string  `json:"fyToken"`
	Tt             string  `json:"tt"`
}

type MarginAPIResp_Fyers struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		MarginAvail float64 `json:"margin_avail"`
		MarginTotal float64 `json:"margin_total"`
	} `json:"data"`
	S string `json:"s"`
}

type OptionChainAPIResponse struct {
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
	S       string `json:"s"`
}

type Data struct {
	CallOi       int64         `json:"callOi"`
	ExpiryData   []Expiry      `json:"expiryData"`
	IndiaVixData IndiaVix      `json:"indiavixData"`
	OptionsChain []OptionChain `json:"optionsChain"`
	PutOi        int64         `json:"putOi"`
}

type Expiry struct {
	Date   string `json:"date"`
	Expiry string `json:"expiry"`
}

type IndiaVix struct {
	Ask         float64 `json:"ask"`
	Bid         float64 `json:"bid"`
	Description string  `json:"description"`
	ExSymbol    string  `json:"ex_symbol"`
	Exchange    string  `json:"exchange"`
	FyToken     string  `json:"fyToken"`
	Ltp         float64 `json:"ltp"`
	LtpCh       float64 `json:"ltpch"`
	LtpChp      float64 `json:"ltpchp"`
	OptionType  string  `json:"option_type"`
	StrikePrice int64   `json:"strike_price"`
	Symbol      string  `json:"symbol"`
}

type OptionChain struct {
	Ask         float64 `json:"ask"`
	Bid         float64 `json:"bid"`
	Description string  `json:"description,omitempty"`
	ExSymbol    string  `json:"ex_symbol,omitempty"`
	Exchange    string  `json:"exchange,omitempty"`
	Fp          float64 `json:"fp,omitempty"`
	Fpch        float64 `json:"fpch,omitempty"`
	Fpchp       float64 `json:"fpchp,omitempty"`
	FyToken     string  `json:"fyToken"`
	Ltp         float64 `json:"ltp"`
	LtpCh       float64 `json:"ltpch"`
	LtpChp      float64 `json:"ltpchp"`
	Oi          int64   `json:"oi,omitempty"`
	Oich        int64   `json:"oich,omitempty"`
	Oichp       float64 `json:"oichp,omitempty"`
	OptionType  string  `json:"option_type"`
	PrevOi      int64   `json:"prev_oi,omitempty"`
	StrikePrice int64   `json:"strike_price"`
	Symbol      string  `json:"symbol"`
	Volume      int64   `json:"volume,omitempty"`
}

type Symbol struct {
	Name          string // Name of the symbol Including Exchange
	FyToken       string // Fyers Token
	TradingSymbol string // Trading Symbol Excluding Exchange
}

// Struct to represent the History API response
type HistoricalDataAPI_Resp struct {
	Status  string          `json:"s"`
	Candles [][]interface{} `json:"candles"`
}

// Struct to represent individual candle data
type Candle struct {
	EpochTime int64   // Current epoch time
	Open      float64 // Open Value
	High      float64 // Highest Value
	Low       float64 // Lowest Value
	Close     float64 // Close Value
	Volume    int64   // Total traded quantity (volume)
}

type PlaceOrderResponse struct {
	Status  string `json:"s"` // ok / error
	Code    int    `json:"code"`
	Message string `json:"message"` // status message
	ID      string `json:"id"`      // order id
}

// HoldingsResponse represents the response from the holding API.
// It contains the status, a list of holdings items and overall details.
type HoldingsResponse struct {
	S        string          `json:"s"`        // status of the response ok/error
	Code     int             `json:"code"`     // code of the response
	Message  string          `json:"message"`  // message of the response
	Holdings []Holding       `json:"holdings"` // list of holding items
	Overall  HoldingsOverall `json:"overall"`  // overall details of the holdings
}

type Holding struct {
	HoldingType             string  `json:"holdingType"`
	Quantity                int     `json:"quantity"`
	CostPrice               float64 `json:"costPrice"`
	MarketVal               float64 `json:"marketVal"`
	RemainingQuantity       int     `json:"remainingQuantity"`
	PL                      float64 `json:"pl"`
	LTP                     float64 `json:"ltp"`
	ID                      int     `json:"id"`
	FyToken                 string  `json:"fyToken"`
	Exchange                int     `json:"exchange"`
	Symbol                  string  `json:"symbol"`
	Segment                 int     `json:"segment"`
	ISIN                    string  `json:"isin"`
	QtyT1                   int     `json:"qty_t1"`
	RemainingPledgeQuantity int     `json:"remainingPledgeQuantity"`
	CollateralQuantity      int     `json:"collateralQuantity"`
}

type HoldingsOverall struct {
	CountTotal        int     `json:"count_total"`
	TotalInvestment   float64 `json:"total_investment"`
	TotalCurrentValue float64 `json:"total_current_value"`
	TotalPL           float64 `json:"total_pl"`
	PnLPercent        float64 `json:"pnl_perc"`
}

// FundsResponse represents the response from the funds API.
// It contains the code, message, status and list of fund items.
type FundsResponse struct {
	Code      int        `json:"code"`
	Message   string     `json:"message"`
	S         string     `json:"s"`
	FundLimit []FundItem `json:"fund_limit"` // list of fund items
}

type FundItem struct {
	ID              int     `json:"id"`
	Title           string  `json:"title"`
	EquityAmount    float64 `json:"equityAmount"`
	CommodityAmount float64 `json:"commodityAmount"`
}
