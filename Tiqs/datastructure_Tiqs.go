package tiqs

type ReadDataJsonTiqs struct {
	Date        string `json:"Date"`
	AccessToken string `json:"token"`
	Session     string `json:"session"`
	APPID       string `json:"appId"`
}

type NetPosition_Tiqs struct {
	// AvgPrice               string `json:"avgPrice"`
	// BreakEvenPrice         string `json:"breakEvenPrice"`
	// CarryForwarAvgPrice    string `json:"carryForwarAvgPrice"`
	// CarryForwardBuyAmount  string `json:"carryForwardBuyAmount"`
	// CarryForwardBuyAvgPrice string `json:"carryForwardBuyAvgPrice"`
	// CarryForwardBuyQty     string `json:"carryForwardBuyQty"`
	// CarryForwardSellAmount string `json:"carryForwardSellAmount"`
	// CarryForwardSellAvgPrice string `json:"carryForwardSellAvgPrice"`
	// CarryForwardSellQty    string `json:"carryForwardSellQty"`
	DayBuyAmount    string `json:"dayBuyAmount"`
	DayBuyAvgPrice  string `json:"dayBuyAvgPrice"`
	DayBuyQty       string `json:"dayBuyQty"`
	DaySellAmount   string `json:"daySellAmount"`
	DaySellAvgPrice string `json:"daySellAvgPrice"`
	DaySellQty      string `json:"daySellQty"`
	Exchange        string `json:"exchange"`
	LotSize         string `json:"lotSize"`
	Ltp             string `json:"ltp"`
	// Multiplier             string `json:"multiplier"`
	// NetUploadPrice         string `json:"netUploadPrice"`
	// OpenBuyAmount          string `json:"openBuyAmount"`
	// OpenBuyAvgPrice        string `json:"openBuyAvgPrice"`
	// OpenBuyQty             string `json:"openBuyQty"`
	// OpenSellAmount         string `json:"openSellAmount"`
	// OpenSellAvgPrice       string `json:"openSellAvgPrice"`
	// OpenSellQty            string `json:"openSellQty"`
	PriceFactor string `json:"priceFactor"`
	// PricePrecision         string `json:"pricePrecision"`
	Product     string `json:"product"`
	Qty         string `json:"qty"`
	RealisedPnL string `json:"realisedPnL"`
	Symbol      string `json:"symbol"`
	// TickSize               string `json:"tickSize"`
	Token                  string `json:"token"`
	UnrealisedMarkToMarket string `json:"unrealisedMarkToMarket"`
	// UploadPrice            string `json:"uploadPrice"`
}

type PositionAPIResp_Tiqs struct {
	S                 string             `json:"status"`
	NetPosition_Tiqss []NetPosition_Tiqs `json:"data"`
}

type Order_Tiqs struct {
	Orders   int `json:"orders"`
	Price    int `json:"price"`
	Quantity int `json:"quantity"`
}

type QuotesData_Tiqs struct {
	Asks         []Order_Tiqs `json:"asks"`
	AvgPrice     int          `json:"avgPrice"`
	Bids         []Order_Tiqs `json:"bids"`
	Close        int          `json:"close"`
	High         int          `json:"high"`
	Low          int          `json:"low"`
	LowerLimit   int          `json:"lowerLimit"`
	Ltp          int          `json:"ltp"`
	Ltq          int          `json:"ltq"`
	Ltt          int          `json:"ltt"`
	NetChange    int          `json:"netChange"`
	NetChangeInd int          `json:"netChangeIndicator"`
	Oi           int          `json:"oi"`
	OiDayHigh    int          `json:"oiDayHigh"`
	OiDayLow     int          `json:"oiDayLow"`
	Open         int          `json:"open"`
	OpeningOi    int          `json:"openingOI"`
	Time         int          `json:"time"`
	Token        int          `json:"token"`
	TotalBuyQty  int          `json:"totalBuyQty"`
	TotalSellQty int          `json:"totalSellQty"`
	UpperLimit   int          `json:"upperLimit"`
	Volume       int          `json:"volume"`
}
type QuotesAPIResp_Tiqs struct {
	Data      []QuotesData_Tiqs       `json:"data"`
	Status    string                  `json:"status"`
	TokenData map[int]QuotesData_Tiqs `json:"-"`
}

type Option struct {
	Exchange       string `json:"exchange"`
	Symbol         string `json:"symbol"`
	Token          string `json:"token"`
	OptionType     string `json:"optionType"`
	StrikePrice    string `json:"strikePrice"`
	PricePrecision string `json:"pricePrecision"`
	TickSize       string `json:"tickSize"`
	LotSize        string `json:"lotSize"`
	OpeningOI      string `json:"openingOI"`
}

type OptionChainResp_Tiqs struct {
	Data   []Option `json:"data"`
	Status string   `json:"status"`
}

type ExpiryResp_Tiqs struct {
	Data   map[string][]string `json:"data"`
	Status string              `json:"status"`
}

type LTPofTokenResp_Tiqs struct {
	Data struct {
		Close int `json:"close"`
		LTP   int `json:"ltp"`
		Token int `json:"token"`
	} `json:"data"`
	Status string `json:"status"`
}

type Symbol struct {
	Name  string
	Token string
}
