package fyers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// PositionApi_Fyers fetches the current positions for a given user from the Fyers API.
// It takes the UserID of the user as a parameter and returns the positions and an error if any occurs.
func PositionApi_Fyers(UserID_Fyers string) (PositionAPIResp_Fyers, error) {

	// Retrieve the access token for the user
	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return PositionAPIResp_Fyers{}, err
	}

	// Define the URL for the positions API endpoint
	positionUrl := "https://api-t1.fyers.in/trade/v3/positions"

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", positionUrl, nil)
	if err != nil {
		log.Println("Error while making request in Position API request in Fyers")
		return PositionAPIResp_Fyers{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)

	// Make the request to the API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in Position API in Fyers")
		return PositionAPIResp_Fyers{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in Position API in Fyers")
		return PositionAPIResp_Fyers{}, err
	}

	// Log the direct response from the API
	jsonBody := string(body)
	log.Printf("Direct Response from Position API of Fyers %v", jsonBody)

	// Convert the response body into the PositionAPIResp_Fyers struct
	var positionResp PositionAPIResp_Fyers
	err = json.Unmarshal(body, &positionResp)
	if err != nil {
		log.Println("Error while Unmarshaling the data in Position API in Fyers")
		return PositionAPIResp_Fyers{}, err
	}

	// Return the positions and nil error
	return positionResp, nil
}

// ExitingAllPosition deletes all open positions for a given user in Fyers.
// It takes the list of sides, segments, product types and UserID_Fyers as parameters and returns an error if any occurs.
func ExitingAllPosition(Side, Segement []int, ProductType []string, UserID_Fyers string) error {

	// Retrieve the access token for the user
	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return err
	}

	// Define the URL for the positions API endpoint
	positionUrl := "https://api-t1.fyers.in/trade/v3/positions"

	// Define the data payload for the request
	dataPayload := map[string]interface{}{
		"exit_all":    0,
		"side":        Side,
		"segment":     Segement,
		"productType": ProductType,
	}

	// Marshal the data payload into JSON
	jsonData, err := json.Marshal(dataPayload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return err
	}

	// Create a new HTTP DELETE request
	req, err := http.NewRequest("DELETE", positionUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error while making request in Exit All Position API in Fyers")
		return err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in Exit All Position API in Fyers")
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in Exit All Position API")
		return err
	}

	// Log the direct response from the API
	msg := string(body)
	log.Printf("Direct Response for %v from Fyers API while Exiting all positions %v", UserID_Fyers, msg)

	// Log the final message
	finalMsg := fmt.Sprintf("For User : %v \nExit All API response %v", UserID_Fyers, msg)
	log.Println(finalMsg)
	return err
}

// ExitPositionByID_Fyers exits a position by its ID.
// It takes the UserID_Fyers and symbol name as parameters and returns an error if something goes wrong.
func ExitPositionByID_Fyers(UserID_Fyers string, symbolName string) error {

	// Retrieve the access token for the user
	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return err
	}

	// Define the URL for the positions API endpoint
	positionUrl := "https://api-t1.fyers.in/trade/v3/positions"

	// Define the data payload for the request
	dataPayload := map[string]interface{}{
		"id": symbolName,
	}

	// Marshal the data payload into JSON
	jsonData, err := json.Marshal(dataPayload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return err
	}

	// Create a new HTTP DELETE request
	req, err := http.NewRequest("DELETE", positionUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error while making request in Exiting via Position API")
		return err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in Position API")
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in Position API")
		return err
	}

	// Log the direct response from the API
	msg := string(body)
	log.Printf("Direct Response from Fyers API while Exiting position by ID %v", msg)

	// Create a new message that includes the user and the exit message
	newMsg := fmt.Sprintf("%v User \n\nExit Alert From Position\n\n", UserID_Fyers)
	newMsg += msg
	log.Println(newMsg)

	return nil
}

// MarketDepthAPI_Fyers fetches the market depth for a given symbol from the Fyers API.
// It takes the symbol name and UserID of the user as parameters and returns the market depth and an error if any occurs.
func MarketDepthAPI_Fyers(symbolName string, UserID_Fyers string) (MarketDepthAPIResp_Fyers, error) {

	// Retrieve the access token for the user
	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return MarketDepthAPIResp_Fyers{}, err
	}

	// Define the URL for the positions API endpoint
	url := fmt.Sprintf("https://api-t1.fyers.in/data/depth?symbol=%s&ohlcv_flag=1", symbolName)

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error while making request in marketDepthAPI")
		return MarketDepthAPIResp_Fyers{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in marketDepthAPI")
		return MarketDepthAPIResp_Fyers{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in marketDepthAPI")
		return MarketDepthAPIResp_Fyers{}, err
	}

	// Log the direct response from the API
	jsonBody := string(body)
	log.Printf("Direct Response from Market Depth API of fyers for %v is %v", symbolName, jsonBody)

	// Convert the response body into the MarketDepthAPIResp_Fyers struct
	var marketDepthResponse MarketDepthAPIResp_Fyers

	err = json.Unmarshal(body, &marketDepthResponse)
	if err != nil {
		log.Println("Error while Unmarshaling the data in marketDepthAPI")
		return MarketDepthAPIResp_Fyers{}, err
	}

	return marketDepthResponse, nil
}

// LTP_Fyers fetches the last traded price (LTP) of a symbol from Fyers using the Market Depth API.
func LTP_Fyers(symbolName string, UserID_Fyers string) (float64, error) {

	// Retrieve the access token for the user
	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return 0.0, err
	}

	// Construct the URL for the Market Depth API request
	url := fmt.Sprintf("https://api-t1.fyers.in/data/depth?symbol=%s&ohlcv_flag=1", symbolName)

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error while making request in marketDepthAPI")
		return 0.0, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in marketDepthAPI")
		return 0.0, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in marketDepthAPI")
		return 0.0, err
	}

	// Log the direct response from the API
	jsonBody := string(body)
	log.Printf("Direct Response from Market Depth API of fyers for %v is %v", symbolName, jsonBody)

	// Convert the response body into the MarketDepthAPIResp_Fyers struct
	var marketDepthResponse MarketDepthAPIResp_Fyers

	err = json.Unmarshal(body, &marketDepthResponse)
	if err != nil {
		log.Println("Error while Unmarshaling the data in marketDepthAPI")
		return 0.0, err
	}

	// ltp is the last traded price
	ltp := marketDepthResponse.D[symbolName].Ltp
	final := fmt.Sprintf("\nFor %v LTP is %v", symbolName, ltp)
	log.Println(final)

	return ltp, nil
}

// PlaceOrder_Fyers places a limit order for the given symbol using the Fyers API.
// It takes the symbol name, limit price, quantity, side, product type and user ID as parameters.
// Returns a boolean indicating success and an error if any occurs.
func PlaceOrder_Fyers(symbolName string, LimitPriceForOrder float64, qty int, whichSide int, productType string, UserID_Fyers string) (bool, error) {

	// Retrieve the access token for the user
	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return false, err
	}

	// Log the order details
	msg := fmt.Sprintf("Placing Order for %v and Price is %v and total qty is %v and Client Name is %v", symbolName, LimitPriceForOrder, qty, UserID_Fyers)
	log.Println(msg)

	// Define the URL for the order placement API endpoint
	url := "https://api-t1.fyers.in/api/v3/orders/sync"

	// Prepare the payload for the order
	dataPayload := map[string]interface{}{
		"symbol":       symbolName,
		"qty":          qty,
		"type":         1, // Assuming 1 represents a limit order
		"side":         whichSide,
		"productType":  productType,
		"limitPrice":   LimitPriceForOrder,
		"disclosedQty": 0,
		"stopPrice":    0,
		"validity":     "DAY",
		"offlineOrder": false,
		"stopLoss":     0,
		"takeProfit":   0,
		"orderTag":     "orderFromSDK",
	}

	// Marshal the payload into JSON
	jsonData, err := json.Marshal(dataPayload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return false, err
	}

	// Create a new HTTP POST request with the JSON payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error creating new HTTP request:", err)
		return false, err
	}

	// Add headers to the request
	req.Header.Add("Authorization", AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// Log the payload being sent
	msg = fmt.Sprintf("Order Placement Payload is %v", dataPayload)
	log.Println(msg)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending HTTP request:", err)
		return false, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return false, err
	}

	// Log the response status and body
	msg = fmt.Sprintf("Order Place Status for : %v is %v  \n\nAfter order placement response is %v ", symbolName, resp.Status, string(body))
	log.Println(msg)

	return true, nil
}
// PlaceMktOrder_Fyers places a market order for the given symbol using the Fyers API.
// It takes the symbol name, quantity, side, product type and user ID as parameters.
// Returns a boolean indicating success and an error if any occurs.
func PlaceMktOrder_Fyers(symbolName string, qty int, whichSide int, productType string, UserID_Fyers string) (bool, error) {

	// Retrieve the access token for the user
	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return false, err
	}

	// Log the order details
	msg := fmt.Sprintf("Placing Market Order for %v and total qty is %v and Client Name is %v", symbolName, qty, UserID_Fyers)
	log.Println(msg)
	// TelegramSend(msg)

	// Define the URL for the order placement API endpoint
	url := "https://api-t1.fyers.in/api/v3/orders/sync"

	// Prepare the payload for the order
	dataPayload := map[string]interface{}{
		"symbol":       "NSE:ITC-EQ",
		"qty":          1,
		"type":         2,
		"side":         whichSide,
		"productType":  productType,
		"limitPrice":   0,
		"disclosedQty": 0,
		"stopPrice":    0,
		"validity":     "DAY",
		"offlineOrder": false,
		"stopLoss":     0,
		"takeProfit":   0,
		"orderTag":     "orderFromSDK",
	}

	// Change the name of symbol and price in the payload
	dataPayload["symbol"] = symbolName
	dataPayload["qty"] = qty

	// Marshal the payload into JSON
	jsonData, err := json.Marshal(dataPayload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return false, err
	}

	// Create a new HTTP POST request with the JSON payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return false, err
	}

	// Add headers to the request
	req.Header.Add("Authorization", AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// Log the payload being sent
	msg = fmt.Sprintf("Order Placement Payload is %v", dataPayload)
	log.Println(msg)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer resp.Body.Close()

	// Print the response status and body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false, err
	}
	msg = fmt.Sprintf("Order Place Status for : %v is %v  \n\nAfter order placement response is %v ", symbolName, resp.Status, string(body))
	log.Println(msg)

	return true, nil
}

// QuotesAPI_Fyers fetches the quote data for a given symbol from the Fyers API.
// It takes the symbol name and UserID of the user as parameters and returns the quote response and an error if any occurs.
func QuotesAPI_Fyers(symbolName, UserID_Fyers string) (QuoteAPIResp_Fyers, error) {

	// Retrieve the access token for the user
	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return QuoteAPIResp_Fyers{}, err
	}

	// Construct the URL for the Quotes API request
	url := fmt.Sprintf("https://api-t1.fyers.in/data/quotes?symbols=%s", symbolName)

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error while making request in quotesFyersAPI")
		return QuoteAPIResp_Fyers{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)

	// Make the request to the API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in quotesFyersAPI")
		return QuoteAPIResp_Fyers{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in quotesFyersAPI")
		return QuoteAPIResp_Fyers{}, err
	}

	// Log the direct response from the API
	jsonBody := string(body)
	log.Printf("Direct Response from Quotes API of fyers for %v is %v", symbolName, jsonBody)

	// Convert the response body into the QuoteAPIResp_Fyers struct
	var qpr QuoteAPIResp_Fyers
	err = json.Unmarshal(body, &qpr)
	if err != nil {
		log.Println("Error while Unmarshaling the data in quotes Fyers API")
		return QuoteAPIResp_Fyers{}, err
	}

	return qpr, nil
}

// SymbolNameToExchToken retrieves the exchange token for a given symbol name from the Fyers API.
// It takes the symbol name and UserID of the user as parameters and returns the exchange token and an error if any occurs.
func SymbolNameToExchToken(symbolName, UserID_Fyers string) (string, error) {

	// Retrieve the access token for the user
	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return "", err
	}

	// Construct the URL for the Quotes API request
	url := fmt.Sprintf("https://api-t1.fyers.in/data/quotes?symbols=%s", symbolName)

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error while making request in quotesFyersAPI")
		return "", err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)

	// Make the request to the API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in quotesFyersAPI")
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in quotesFyersAPI")
		return "", err
	}

	// Log the direct response from the API
	jsonBody := string(body)
	log.Printf("Direct Response from Quotes API of fyers for %v is %v", symbolName, jsonBody)

	// Convert the response body into the QuoteAPIResp_Fyers struct
	var qpr QuoteAPIResp_Fyers

	err = json.Unmarshal(body, &qpr)
	if err != nil {
		log.Println("Error while Unmarshaling the data in marketDepthAPI")
		return "", err
	}

	// Extract the Fyers token
	fytoken := qpr.D[0].V.FyToken
	final := fmt.Sprintf("\nFor %v Fyers token is %v", symbolName, fytoken)
	log.Println(final)

	// Check if the token meets the minimum length requirement
	if len(fytoken) < 12 {
		err := errors.New("fyers token(FyToken) is less than 12 ")
		return "", err
	}

	// Remove the first 10 characters to get the exchange token
	// For more info read this : https://myapi.fyers.in/docsv3#tag/Appendix/Fytoken
	exchangeToken := fytoken[10:]
	return exchangeToken, nil
}

// MarginMktOrder_Fyers retrieves the margin information for a market order for a given symbol using the Fyers API.
// It takes the symbol name, quantity, side, product type, and user ID as parameters.
// Returns a MarginAPIResp_Fyers struct and an error if any occurs.
func MarginMktOrder_Fyers(symbolName string, qty int, whichSide int, productType string, UserID_Fyers string) (MarginAPIResp_Fyers, error) {

	// Retrieve the access token for the user
	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return MarginAPIResp_Fyers{}, err
	}

	// Log the order details
	msg := fmt.Sprintf("Margin Market Order for %v and total qty is %v and Client Name is %v", symbolName, qty, UserID_Fyers)
	log.Println(msg)

	// Define the URL for the margin API endpoint
	url := "https://api-t1.fyers.in/trade/v3/margin"

	// Prepare the payload for the margin request
	dataPayload := map[string]interface{}{
		"symbol":       symbolName,
		"qty":          qty,
		"type":         2, // Assuming 2 represents a market order
		"side":         whichSide,
		"productType":  productType,
		"limitPrice":   0, // Limit price is set to 0 for market order
		"disclosedQty": 0,
		"stopPrice":    0,
		"validity":     "DAY",
		"offlineOrder": false,
		"stopLoss":     0,
		"takeProfit":   0,
		"orderTag":     "orderFromSDK",
	}

	// Marshal the payload into JSON
	jsonData, err := json.Marshal(dataPayload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return MarginAPIResp_Fyers{}, err
	}

	// Create a new HTTP POST request with the JSON payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return MarginAPIResp_Fyers{}, err
	}

	// Add headers to the request
	req.Header.Add("Authorization", AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// Make the request to the API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return MarginAPIResp_Fyers{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return MarginAPIResp_Fyers{}, err
	}

	// Unmarshal the response into the MarginAPIResp_Fyers struct
	var response MarginAPIResp_Fyers
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Failed to unmarshal response in Margin API in Fyers", err)
		return MarginAPIResp_Fyers{}, err
	}

	return response, nil
}

// GetOptionChain_Fyers fetches the option chain for a given symbol and strike count from the Fyers API.
// It takes the symbol name, strike count and UserID of the user as parameters and returns the option chain response and an error if any occurs.
func GetOptionChain_Fyers(Symbol string, StrikeCount int, UserID_Fyers string) (OptionChainAPIResponse, error) {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return OptionChainAPIResponse{}, err
	}

	// Construct the URL for the option chain API request
	url := fmt.Sprintf("https://api-t1.fyers.in/data/options-chain-v3?symbol=%s&strikecount=%d", Symbol, StrikeCount)

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return OptionChainAPIResponse{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return OptionChainAPIResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return OptionChainAPIResponse{}, err
	}

	// Log the direct response from the API
	jsonBody := string(body)
	log.Printf("Direct Response from Option Chain API of fyers for %v is %v", Symbol, jsonBody)

	// Convert the response body into the OptionChainAPIResponse struct
	var response OptionChainAPIResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Error while Unmarshaling the data in Option Chain API")
		return OptionChainAPIResponse{}, err
	}
	return response, nil
}

// GetHistoricalData_Fyers fetches the historical data for a given symbol from the Fyers API.
// It takes the symbol name, resolution, date format, range from and range to as parameters and returns the historical data and an error if any occurs.
func GetHistoricalData_Fyers(symbol, resolution, dateFormat, rangeFrom, rangeTo, UserID_Fyers string) ([]Candle, error) {

	// Retrieve the access token for the user
	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return []Candle{}, err
	}

	// Construct the URL for the historical data API request
	baseURL := "https://api-t1.fyers.in/data/history"

	// Build the URL with query parameters
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	// Add query parameters
	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("resolution", resolution)
	params.Add("date_format", dateFormat)
	params.Add("range_from", rangeFrom)
	params.Add("range_to", rangeTo)
	params.Add("cont_flag", "")

	apiURL.RawQuery = params.Encode()

	// Create HTTP client and set timeout
	client := &http.Client{Timeout: 10 * time.Second}

	// Prepare the GET request
	req, err := http.NewRequest("GET", apiURL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add headers to the request
	req.Header.Add("Authorization", AccessToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check if response status is 200 OK
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	// Unmarshal the response into the temporary structure
	var rawData HistoricalDataAPI_Resp
	err = json.Unmarshal(body, &rawData)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Convert the raw candle data into the []Candle format
	candles, err := convertToCandles(rawData.Candles)
	if err != nil {
		log.Fatalf("Failed to convert candle data: %v", err)
	}

	// Return the candles
	return candles, nil
}

// Convert [][]interface{} to []Candle
func convertToCandles(rawCandles [][]interface{}) ([]Candle, error) {
	var candles []Candle
	for _, rawCandle := range rawCandles {
		if len(rawCandle) != 6 {
			return nil, fmt.Errorf("invalid candle data length: expected 6, got %d", len(rawCandle))
		}

		candle := Candle{
			EpochTime: int64(rawCandle[0].(float64)),
			Open:      rawCandle[1].(float64),
			High:      rawCandle[2].(float64),
			Low:       rawCandle[3].(float64),
			Close:     rawCandle[4].(float64),
			Volume:    int64(rawCandle[5].(float64)),
		}
		candles = append(candles, candle)
	}
	return candles, nil
}
