package tiqs

type readDataJsonTiqs struct {
	Date        string `json:"Date"`
	AccessToken string `json:"token"`
	Session     string `json:"session"`
	APPID       string `json:"appId"`
}

type netPosition_Tiqs struct {
	AvgPrice                 string `json:"avgPrice"`
	BreakEvenPrice           string `json:"breakEvenPrice"`
	CarryForwarAvgPrice      string `json:"carryForwarAvgPrice"`
	CarryForwardBuyAmount    string `json:"carryForwardBuyAmount"`
	CarryForwardBuyAvgPrice  string `json:"carryForwardBuyAvgPrice"`
	CarryForwardBuyQty       string `json:"carryForwardBuyQty"`
	CarryForwardSellAmount   string `json:"carryForwardSellAmount"`
	CarryForwardSellAvgPrice string `json:"carryForwardSellAvgPrice"`
	CarryForwardSellQty      string `json:"carryForwardSellQty"`
	DayBuyAmount             string `json:"dayBuyAmount"`
	DayBuyAvgPrice           string `json:"dayBuyAvgPrice"`
	DayBuyQty                string `json:"dayBuyQty"`
	DaySellAmount            string `json:"daySellAmount"`
	DaySellAvgPrice          string `json:"daySellAvgPrice"`
	DaySellQty               string `json:"daySellQty"`
	Exchange                 string `json:"exchange"`
	LotSize                  string `json:"lotSize"`
	Ltp                      string `json:"ltp"`
	Multiplier               string `json:"multiplier"`
	NetUploadPrice           string `json:"netUploadPrice"`
	OpenBuyAmount            string `json:"openBuyAmount"`
	OpenBuyAvgPrice          string `json:"openBuyAvgPrice"`
	OpenBuyQty               string `json:"openBuyQty"`
	OpenSellAmount           string `json:"openSellAmount"`
	OpenSellAvgPrice         string `json:"openSellAvgPrice"`
	OpenSellQty              string `json:"openSellQty"`
	PriceFactor              string `json:"priceFactor"`
	PricePrecision           string `json:"pricePrecision"`
	Product                  string `json:"product"`
	Qty                      string `json:"qty"`
	RealisedPnL              string `json:"realisedPnL"`
	UnRealisedPnL            string `json:"unRealisedPnl"`
	Symbol                   string `json:"symbol"`
	TickSize                 string `json:"tickSize"`
	Token                    string `json:"token"`
	UnrealisedMarkToMarket   string `json:"unrealisedMarkToMarket"`
	UploadPrice              string `json:"uploadPrice"`
}

type positionAPIResp_Tiqs struct {
	S                 string             `json:"status"`
	NetPosition_Tiqss []netPosition_Tiqs `json:"data"`
}

type orderBookAPIResp_Tiqs struct {
	OrderBook []struct {
		Status                   string `json:"status"`
		UserID                   string `json:"userID"`
		AccountID                string `json:"accountID"`
		Exchange                 string `json:"exchange"`
		Symbol                   string `json:"symbol"`
		ID                       string `json:"id"`
		RejectReason             string `json:"rejectReason"`
		Price                    string `json:"price"`
		Quantity                 string `json:"quantity"`
		MarketProtection         string `json:"marketProtection"`
		Product                  string `json:"product"`
		OrderStatus              string `json:"orderStatus"`
		TransactionType          string `json:"transactionType"`
		Order                    string `json:"order"`
		FillShares               string `json:"fillShares"`
		AveragePrice             string `json:"averagePrice"`
		ExchangeOrderID          string `json:"exchangeOrderID"`
		CancelQuantity           string `json:"cancelQuantity"`
		Tags                     string `json:"tags"`
		DisclosedQuantity        string `json:"disclosedQuantity"`
		OrderTriggerPrice        string `json:"orderTriggerPrice"`
		Retention                string `json:"retention"`
		BookProfitPrice          string `json:"bookProfitPrice"`
		BookLossPrice            string `json:"bookLossPrice"`
		TrailingPrice            string `json:"trailingPrice"`
		Amo                      string `json:"amo"`
		PricePrecision           string `json:"pricePrecision"`
		TickSize                 string `json:"tickSize"`
		LotSize                  string `json:"lotSize"`
		Token                    string `json:"token"`
		TimeStamp                string `json:"timeStamp"`
		OrderTime                string `json:"orderTime"`
		ExchangeUpdateTime       string `json:"exchangeUpdateTime"`
		SnoOrderDirection        string `json:"snoOrderDirection"`
		SnoOrderID               string `json:"snoOrderID"`
		PriceFactor              string `json:"priceFactor"`
		Multiplier               string `json:"multiplier"`
		DisplayName              string `json:"displayName"`
		RequiredQuantity         string `json:"requiredQuantity"`
		RequiredPrice            string `json:"requiredPrice"`
		RequiredTriggerPrice     string `json:"requiredTriggerPrice"`
		RequiredBookLossPrice    string `json:"requiredBookLossPrice"`
		RequiredOriginalQuantity string `json:"requiredOriginalQuantity"`
		RequiredOriginalPrice    string `json:"requiredOriginalPrice"`
		OriginalTriggerPrice     string `json:"originalTriggerPrice"`
		OriginalBookLossPrice    string `json:"originalBookLossPrice"`
	} `json:"data"`
	Status string `json:"status"`
}

type order_Tiqs struct {
	Orders   int `json:"orders"`
	Price    int `json:"price"`
	Quantity int `json:"quantity"`
}

type quotesData_Tiqs struct {
	Asks         []order_Tiqs `json:"asks"`
	AvgPrice     int          `json:"avgPrice"`
	Bids         []order_Tiqs `json:"bids"`
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
type quotesAPIResp_Tiqs struct {
	Data      []quotesData_Tiqs       `json:"data"`
	Status    string                  `json:"status"`
	TokenData map[int]quotesData_Tiqs `json:"-"`
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

type optionChainResp_Tiqs struct {
	Data   []Option `json:"data"`
	Status string   `json:"status"`
}

type ExpiryResp_Tiqs struct {
	Data   map[string][]string `json:"data"`
	Status string              `json:"status"`
}

type ltpofTokenResp_Tiqs struct {
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

// Define greeksData_Tiqs struct
type greeksData_Tiqs struct {
	Delta float64 `json:"delta"`
	Theta float64 `json:"theta"`
	Gamma float64 `json:"gamma"`
	Vega  float64 `json:"vega"`
	IV    float64 `json:"iv"`
}

// Define greeksResp_Tiqs struct
type greeksResp_Tiqs struct {
	Data   []greeksData_Tiqs `json:"data"`
	Status string            `json:"status"`
}

type holidaysData_Tiqs struct {
	Holidays map[string]string `json:"holidays"`
	// SpecialTradingDays map[string][]string `json:"specialTradingDays"` // Not requried as of now
}

type holidaysAPIResp_Tiqs struct {
	Data   holidaysData_Tiqs `json:"data"`
	Status string            `json:"status"`
}

type placeOrderResp_Tiqs struct {
	Message string `json:"message"`
	Data    struct {
		OrderNo     string `json:"orderNo"`
		RequestTime string `json:"requestTime"`
	} `json:"data"`
	Status string `json:"status"`
}
