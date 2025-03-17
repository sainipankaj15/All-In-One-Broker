package xts

type readDataJsonXTS struct {
	Date        string `json:"Date"`
	AccessToken string `json:"token"`
	UserID      string `json:"userID"`
}

type placeOrderResp_XTS struct {
	Type        string `json:"type"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Result      struct {
		AppOrderID            int    `json:"AppOrderID"`
		ClientID              string `json:"ClientID"`
		OrderUniqueIdentifier string `json:"OrderUniqueIdentifier"`
	} `json:"result"`
}

type placeOrderReq_XTS struct {
	ExchangeSegment       string  `json:"exchangeSegment"`
	ExchangeInstrumentID  int64   `json:"exchangeInstrumentID"`
	ProductType           string  `json:"productType"`
	OrderType             string  `json:"orderType"`
	OrderSide             string  `json:"orderSide"`
	TimeInForce           string  `json:"timeInForce"`
	DisclosedQuantity     int64   `json:"disclosedQuantity"`
	OrderQuantity         int64   `json:"orderQuantity"`
	LimitPrice            float64 `json:"limitPrice"`
	StopPrice             int64   `json:"stopPrice"`
	OrderUniqueIdentifier string  `json:"orderUniqueIdentifier"`
}
