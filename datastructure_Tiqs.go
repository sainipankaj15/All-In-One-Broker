package main

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
