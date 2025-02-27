package jainam

type placeOrderResp_Jainam []struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		RequestTime string `json:"requestTime"`
		OrderNo     string `json:"orderNo"`
	} `json:"result"`
}

type placeOrderReq_Jainam struct {
	Exchange      string `json:"exchange"`
	Qty           string `json:"qty"`
	Price         string `json:"price"`
	Product       string `json:"product"`
	TransType     string `json:"transType"`
	PriceType     string `json:"priceType"`
	TriggerPrice  int    `json:"triggerPrice"`
	Ret           string `json:"ret"`
	DisclosedQty  int    `json:"disclosedQty"`
	MktProtection string `json:"mktProtection"`
	Target        int    `json:"target"`
	StopLoss      int    `json:"stopLoss"`
	OrderType     string `json:"orderType"`
	Token         string `json:"token"`
}

type readDataJsonJainam struct {
	Date        string `json:"Date"`
	AccessToken string `json:"token"`
	UserID      string `json:"userID"`
}
