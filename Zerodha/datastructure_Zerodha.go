package zerodha

type PlaceOrderResponse struct {
	Status    string    `json:"status"`
	Data      OrderData `json:"data"`
	Message   string    `json:"message"`
	ErrorType string    `json:"error_type"`
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
