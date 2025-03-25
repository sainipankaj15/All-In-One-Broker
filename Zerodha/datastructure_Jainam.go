package zerodha

type placeOrderResp_Zerodha struct {
	Status string `json:"status"`
	Data   struct {
		OrderID string `json:"order_id"`
	} `json:"data"`
	Message   string `json:"message"`
	ErrorType string `json:"error_type"`
}
type placeOrderReq_Zerodha struct {
	Tradingsymbol   string `json:"tradingsymbol"`
	Exchange        string `json:"exchange"`
	TransactionType string `json:"transaction_type"`
	OrderType       string `json:"order_type"`
	Quantity        string `json:"quantity"`
	Product         string `json:"product"`
	Validity        string `json:"validity"`
}

type readDataJsonZerodha struct {
	Date        string `json:"Date"`
	ApiKey      string `json:"apiKey"`
	AccessToken string `json:"token"`
	UserID      string `json:"userID"`
}
