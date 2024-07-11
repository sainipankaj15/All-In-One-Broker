package fyers

type ReadDataJson_Fyers struct {
	Date                 string `json:"Date"`
	AccessToken          string `json:"access_token"`
	AccessTokenWithAppID string `json:"access_token_with_APPID"`
	AppId                string `json:"app_id"`
	AppIdWithType        string `json:"app_id_with_app_type"`
}

type NetPosition_Fyers struct {
	Symbol          string  `json:"symbol"`
	ID              string  `json:"id"`
	BuyAvg          float64 `json:"buy_avg"`
	BuyQty          int     `json:"buy_qty"`
	BuyVal          float64 `json:"buy_val"`
	SellAvg         float64 `json:"sell_avg"`
	SellQty         int     `json:"sell_qty"`
	SellVal         float64 `json:"sell_val"`
	NetAvg          float64 `json:"net_avg"`
	NetQty          int     `json:"net_qty"`
	TranSide        int     `json:"tran_side"`
	Qty             int     `json:"qty"`
	ProductType     string  `json:"product_type"`
	PLRealized      float64 `json:"pl_realized"`
	CrossCurrFlag   string  `json:"cross_curr_flag"`
	RBIRefRate      int     `json:"rbirefrate"`
	FYToken         string  `json:"fy_token"`
	SymbolDesc      string  `json:"symbol_desc"`
	SymbolExch      string  `json:"symbol_exch"`
	Exchange        int     `json:"exchange"`
	Segment         int     `json:"segment"`
	Instrument      int     `json:"instrument"`
	LotSize         int     `json:"lot_size"`
	TickSize        float64 `json:"tick_size"`
	DayBuyQty       int     `json:"day_buy_qty"`
	DayBuyAvg       float64 `json:"day_buy_avg"`
	DaySellQty      int     `json:"day_sell_qty"`
	DaySellAvg      float64 `json:"day_sell_avg"`
	DayNetQty       int     `json:"day_net_qty"`
	CFBuyQty        int     `json:"cf_buy_qty"`
	CFBuyAvg        float64 `json:"cf_buy_avg"`
	CFSellQty       int     `json:"cf_sell_qty"`
	CFSellAvg       float64 `json:"cf_sell_avg"`
	CFNetQty        int     `json:"cf_net_qty"`
	OMSFlag         string  `json:"oms_flag"`
	QtyMultiplier   int     `json:"qty_multiplier"`
	PriceMultiplier int     `json:"price_multiplier"`
	PLTotal         float64 `json:"pl_total"`
	PLUnrealized    float64 `json:"pl_unrealized"`
	LTPCh           float64 `json:"ltp_ch"`
	LTPChp          float64 `json:"ltp_chp"`
	LTP             float64 `json:"ltp"`
}

type Overall struct {
	CountTotal   int     `json:"count_total"`
	CountOpen    int     `json:"count_open"`
	PLTotal      float64 `json:"pl_total"`
	PLRealized   float64 `json:"pl_realized"`
	PLUnrealized float64 `json:"pl_unrealized"`
}

type PositionAPIResp_Fyers struct {
	S            string              `json:"s"`
	Code         int                 `json:"code"`
	Message      string              `json:"message"`
	NetPositions []NetPosition_Fyers `json:"netPositions"`
	Overall      Overall             `json:"overall"`
}

type MarketDepthAPI_Fyers struct {
	D       map[string]StockData `json:"d"`
	Message string               `json:"message"`
	Status  string               `json:"s"`
}

type StockData struct {
	H   float64 `json:"h"`
	L   float64 `json:"l"`
	Ltp float64 `json:"ltp"`
}

type QuoteAPI_Fyers struct {
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
