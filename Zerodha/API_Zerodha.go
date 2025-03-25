package zerodha

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func OrderPlaceMarket_Zerodha(exchange, tradingSymbol, quantity, orderType, transSide, productType, userID_Zerodha string) (placeOrderResp_Zerodha, error) {
	// Get access token
	_, apiKey, accessTokenofUser, err := ReadingAccessToken_Zerodha(userID_Zerodha)
	if err != nil {
		fmt.Println("Error while reading access token in OrderPlaceMarket_Zerodha()")
		return placeOrderResp_Zerodha{}, err
	}
	fmt.Println("Access Token is ", accessTokenofUser)

	// Create form data
	formData := url.Values{}
	formData.Set("tradingsymbol", tradingSymbol)
	formData.Set("exchange", exchange)
	formData.Set("transaction_type", transSide)
	formData.Set("order_type", orderType)
	formData.Set("quantity", quantity)
	formData.Set("product", productType)
	formData.Set("validity", "DAY")

	fmt.Println("Form data is ", formData)

	// Create HTTP client
	client := &http.Client{}

	// Updated URL from the example
	url := placeOrderUrl

	// Create HTTP request with form encoded data
	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println("Error while making request for Order Placement API")
		return placeOrderResp_Zerodha{}, err
	}

	// Add headers to the request
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "token "+apiKey+":"+accessTokenofUser)
	req.Header.Set("X-Kite-Version", "3")

	fmt.Println("Request is ", req)

	// Send HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while getting response in OrderPlaceMarket_Zerodha()")
		return placeOrderResp_Zerodha{}, err
	}
	fmt.Println("resp", resp)
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return placeOrderResp_Zerodha{}, err
	}

	// Parse the response
	var response placeOrderResp_Zerodha
	err = json.Unmarshal(body, &response)
	if err != nil {
		return placeOrderResp_Zerodha{}, fmt.Errorf("error decoding response: %w", err)
	}

	if response.Status != apiResponseStatus.SUCCESS {
		return response, fmt.Errorf("order placement failed: %s", response)
	}

	return response, nil
}
