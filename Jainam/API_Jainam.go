package jainam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func OrderPlaceMarket_Jainam(exchange, token, quantity, priceType, orderType, transSide, productType, userID_Jainam string) (placeOrderResp_Jainam, error) {
	// Get access token
	_, accessTokenofUser, err := ReadingAccessToken_Jainam(userID_Jainam)
	if err != nil {
		fmt.Println("Error while reading access token in OrderPlaceMarket_Jainam()")
		return placeOrderResp_Jainam{}, err
	}

	// Create the order request object
	order := placeOrderReq_Jainam{
		Exchange:      exchange,
		Qty:           quantity,
		Price:         "",
		Product:       productType,
		TransType:     transSide,
		PriceType:     priceType, // Changed from OrderType.MARKET to "MKT"
		TriggerPrice:  0,
		Ret:           "DAY",
		DisclosedQty:  0,
		MktProtection: "",
		Target:        0,
		StopLoss:      0,
		OrderType:     orderType,
		Token:         token,
	}

	// Create a slice with the order as its only element
	orderSlice := []placeOrderReq_Jainam{order}

	jsonParameters, err := json.Marshal(orderSlice)
	if err != nil {
		fmt.Println("Error marshaling JSON in OrderPlaceMarket_Jainam()", err)
		return placeOrderResp_Jainam{}, err
	}

	// Create HTTP client
	client := &http.Client{}

	// Updated URL from the example
	url := placeOrderUrl

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonParameters))
	if err != nil {
		fmt.Println("Error while making request for Order Placement API")
		return placeOrderResp_Jainam{}, err
	}

	// Add headers to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessTokenofUser)

	// Send HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while getting response in OrderPlaceMarket_Jainam()")
		return placeOrderResp_Jainam{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return placeOrderResp_Jainam{}, err
	}

	// Parse the response
	var response placeOrderResp_Jainam
	err = json.Unmarshal(body, &response)
	if err != nil {
		return placeOrderResp_Jainam{}, fmt.Errorf("error decoding response: %w", err)
	}

	if response.Status != apiResponseStatus.SUCCESS {
		return response, fmt.Errorf("order placement failed: %s", response)
	}

	return response, nil
}
