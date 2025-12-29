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
