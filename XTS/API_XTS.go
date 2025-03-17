package xts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func OrderPlaceMarket_XTS(exchange, token, quantity, orderType, transSide, productType, userid_XTS string) (placeOrderResp_XTS, error) {
	// Get access token
	_, accessTokenofUser, err := ReadingAccessToken_XTS(userid_XTS)
	if err != nil {
		fmt.Println("Error while reading access token in OrderPlaceMarket_Jainam()")
		return placeOrderResp_XTS{}, err
	}

	// Create the order request object
	order := placeOrderReq_XTS{
		ExchangeSegment:       exchange,
		ExchangeInstrumentID:  StringToInt(token),
		ProductType:           productType,
		OrderType:             orderType,
		OrderSide:             transSide,
		TimeInForce:           "DAY",
		DisclosedQuantity:     0,
		OrderQuantity:         StringToInt(quantity),
		LimitPrice:            0,
		StopPrice:             0,
		OrderUniqueIdentifier: "XTS_Jainam|TRNKR",
	}

	jsonParameters, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error marshaling JSON in OrderPlaceMarket_Jainam()", err)
		return placeOrderResp_XTS{}, err
	}
	// Create a slice with the order as its only element

	// Create HTTP client
	client := &http.Client{}

	// Updated URL from the example
	url := placeOrderUrl

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonParameters))
	if err != nil {
		fmt.Println("Error while making request for Order Placement API")
		return placeOrderResp_XTS{}, err
	}

	// Add headers to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", accessTokenofUser)

	fmt.Println("Request:", req)
	// Send HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while getting response in OrderPlaceMarket_Jainam()")
		return placeOrderResp_XTS{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return placeOrderResp_XTS{}, err
	}

	// Parse the response
	var response placeOrderResp_XTS
	err = json.Unmarshal(body, &response)
	if err != nil {
		return placeOrderResp_XTS{}, fmt.Errorf("error decoding response: %w", err)
	}

	if response.Code != "200" {
		return response, fmt.Errorf("order placement failed: %v", response)
	}

	return response, nil
}
