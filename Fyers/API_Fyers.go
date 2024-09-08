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

func PositionApi_Fyers(UserID_Fyers string) (PositionAPIResp_Fyers, error) {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return PositionAPIResp_Fyers{}, err
	}

	positionUrl := "https://api-t1.fyers.in/trade/v3/positions"

	req, err := http.NewRequest("GET", positionUrl, nil)
	if err != nil {
		log.Println("Error while making request in Position API request in Fyers")
		return PositionAPIResp_Fyers{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)

	// Make the request
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

	jsonBody := string(body)
	log.Printf("Direct Response from Position API of Fyers %v", jsonBody)

	// Converting into Response struct format
	var positionResp PositionAPIResp_Fyers

	err = json.Unmarshal(body, &positionResp)
	if err != nil {
		log.Println("Error while Unmarshaling the data in Position API in Fyers")
		return PositionAPIResp_Fyers{}, err
	}

	return positionResp, nil
}

func ExitingAllPosition(Side, Segement []int, ProductType []string, UserID_Fyers string) error {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return err
	}

	positionUrl := "https://api-t1.fyers.in/trade/v3/positions"

	dataPayload := map[string]interface{}{
		"exit_all":    0,
		"side":        Side,
		"segment":     Segement,
		"productType": ProductType,
	}

	jsonData, err := json.Marshal(dataPayload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return err
	}

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

	msg := string(body)
	log.Printf("Direct Response for %v from Fyers API while Exiting all positions %v", UserID_Fyers, msg)

	finalMsg := fmt.Sprintf("For User : %v \nExit All API response %v", UserID_Fyers, msg)

	log.Println(finalMsg)
	return err
}

func ExitPositionByID_Fyers(UserID_Fyers string, symbolName string) error {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return err
	}

	positionUrl := "https://api-t1.fyers.in/trade/v3/positions"

	dataPayload := map[string]interface{}{
		"id": symbolName,
	}

	jsonData, err := json.Marshal(dataPayload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return err
	}

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

	msg := string(body)
	log.Printf("Direct Response from Fyers API while Exiting position by ID %v", msg)

	newMsg := fmt.Sprintf("%v User \n\nExit Alert From Position\n\n", UserID_Fyers)
	newMsg += msg
	log.Println(newMsg)

	return nil
}

func MarketDepthAPI_Fyers(symbolName string, UserID_Fyers string) (MarketDepthAPIResp_Fyers, error) {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return MarketDepthAPIResp_Fyers{}, err
	}

	url := fmt.Sprintf("https://api-t1.fyers.in/data/depth?symbol=%s&ohlcv_flag=1", symbolName)

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

	jsonBody := string(body)
	log.Printf("Direct Response from Market Depth API of fyers for %v is %v", symbolName, jsonBody)

	// Converting into Response struct format
	var marketDepthResponse MarketDepthAPIResp_Fyers

	err = json.Unmarshal(body, &marketDepthResponse)
	if err != nil {
		log.Println("Error while Unmarshaling the data in marketDepthAPI")
		return MarketDepthAPIResp_Fyers{}, err
	}

	return marketDepthResponse, nil
}
func LTP_Fyers(symbolName string, UserID_Fyers string) (float64, error) {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return 0.0, err
	}

	url := fmt.Sprintf("https://api-t1.fyers.in/data/depth?symbol=%s&ohlcv_flag=1", symbolName)

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

	jsonBody := string(body)
	log.Printf("Direct Response from Market Depth API of fyers for %v is %v", symbolName, jsonBody)

	// Converting into Response struct format
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

func PlaceOrder_Fyers(symbolName string, LimitPriceForOrder float64, qty int, whichSide int, productType string, UserID_Fyers string) (bool, error) {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return false, err
	}

	// I will use ltp as a limit price to avoid big loss while executing order basically market spread
	msg := fmt.Sprintf("Placing Order for %v and Price is %v and total qty is %v and Client Name is %v", symbolName, LimitPriceForOrder, qty, UserID_Fyers)
	log.Println(msg)
	// TelegramSend(msg)

	url := "https://api-t1.fyers.in/api/v3/orders/sync"

	dataPayload := map[string]interface{}{
		"symbol":       "NSE:ITC-EQ",
		"qty":          1,
		"type":         1,
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

	// Now will change the name of symbol and price
	dataPayload["symbol"] = symbolName
	dataPayload["limitPrice"] = LimitPriceForOrder
	dataPayload["qty"] = qty

	jsonData, err := json.Marshal(dataPayload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return false, err
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return false, err
	}

	// Add headers to the request
	req.Header.Add("Authorization", AccessToken)
	req.Header.Set("Content-Type", "application/json")

	msg = fmt.Sprintf("Order Placement Payload is %v", dataPayload)
	log.Println(msg)

	// Make the request
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
func PlaceMktOrder_Fyers(symbolName string, qty int, whichSide int, productType string, UserID_Fyers string) (bool, error) {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return false, err
	}

	// I will use ltp as a limit price to avoid big loss while executing order basically market spread
	msg := fmt.Sprintf("Placing Market Order for %v and total qty is %v and Client Name is %v", symbolName, qty, UserID_Fyers)
	log.Println(msg)
	// TelegramSend(msg)

	url := "https://api-t1.fyers.in/api/v3/orders/sync"

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

	// Now will change the name of symbol and price
	dataPayload["symbol"] = symbolName
	dataPayload["qty"] = qty

	jsonData, err := json.Marshal(dataPayload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return false, err
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return false, err
	}

	// Add headers to the request
	req.Header.Add("Authorization", AccessToken)
	req.Header.Set("Content-Type", "application/json")

	msg = fmt.Sprintf("Order Placement Payload is %v", dataPayload)
	log.Println(msg)

	// Make the request
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

func QuotesAPI_Fyers(symbolName, UserID_Fyers string) (QuoteAPIResp_Fyers, error) {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return QuoteAPIResp_Fyers{}, err
	}

	url := fmt.Sprintf("https://api-t1.fyers.in/data/quotes?symbols=%s", symbolName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error while making request in quotesFyersAPI")
		return QuoteAPIResp_Fyers{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)

	// Make the request
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

	jsonBody := string(body)
	log.Printf("Direct Response from Quotes API of fyers for %v is %v", symbolName, jsonBody)

	// Converting into Response struct format
	var qpr QuoteAPIResp_Fyers

	err = json.Unmarshal(body, &qpr)
	if err != nil {
		log.Println("Error while Unmarshaling the data in quotes Fyers API")
		return QuoteAPIResp_Fyers{}, err
	}

	return qpr, nil
}

func SymbolNameToExchToken(symbolName, UserID_Fyers string) (string, error) {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return "", err
	}

	url := fmt.Sprintf("https://api-t1.fyers.in/data/quotes?symbols=%s", symbolName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error while making request in quotesFyersAPI")
		return "", err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)

	// Make the request
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

	jsonBody := string(body)
	log.Printf("Direct Response from Quotes API of fyers for %v is %v", symbolName, jsonBody)

	// Converting into Response struct format
	var qpr QuoteAPIResp_Fyers

	err = json.Unmarshal(body, &qpr)
	if err != nil {
		log.Println("Error while Unmarshaling the data in marketDepthAPI")
		return "", err
	}

	// Fyers token
	fytoken := qpr.D[0].V.FyToken
	final := fmt.Sprintf("\nFor %v Fyers token is %v", symbolName, fytoken)
	log.Println(final)

	if len(fytoken) < 12 {
		// Minimum size requirement not met
		err := errors.New("fyers token(FyToken) is less than 12 ")
		return "", err
	}

	// Remove first 10 characters
	// For more info read this : https://myapi.fyers.in/docsv3#tag/Appendix/Fytoken
	exchangeToken := fytoken[10:]
	return exchangeToken, nil
}

func MarginMktOrder_Fyers(symbolName string, qty int, whichSide int, productType string, UserID_Fyers string) (MarginAPIResp_Fyers, error) {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return MarginAPIResp_Fyers{}, err
	}

	// I will use ltp as a limit price to avoid big loss while executing order basically market spread
	msg := fmt.Sprintf("Margin Market Order for %v and total qty is %v and Client Name is %v", symbolName, qty, UserID_Fyers)
	log.Println(msg)

	url := "https://api-t1.fyers.in/trade/v3/margin"

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

	// Now will change the name of symbol and price
	dataPayload["symbol"] = symbolName
	dataPayload["qty"] = qty

	jsonData, err := json.Marshal(dataPayload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return MarginAPIResp_Fyers{}, err
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return MarginAPIResp_Fyers{}, err
	}

	// Add headers to the request
	req.Header.Add("Authorization", AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return MarginAPIResp_Fyers{}, err
	}
	defer resp.Body.Close()

	// Print the response status and body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return MarginAPIResp_Fyers{}, err
	}

	var response MarginAPIResp_Fyers
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Failed to unmarshal response in Margin API in Fyers", err)
		return MarginAPIResp_Fyers{}, err
	}

	return response, nil
}

func GetOptionChain_Fyers(Symbol string, StrikeCount int, UserID_Fyers string) (OptionChainAPIResponse, error) {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return OptionChainAPIResponse{}, err
	}

	url := fmt.Sprintf("https://api-t1.fyers.in/data/options-chain-v3?symbol=%s&strikecount=%d", Symbol, StrikeCount)

	// Create a new HTTP POST request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return OptionChainAPIResponse{}, err
	}

	// Add headers to the request
	req.Header.Add("Authorization", AccessToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return OptionChainAPIResponse{}, err
	}
	defer resp.Body.Close()

	// Print the response status and body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return OptionChainAPIResponse{}, err
	}

	var response OptionChainAPIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Failed to unmarshal response in Option Chain API in Fyers", err)
		return OptionChainAPIResponse{}, err
	}
	return response, nil
}

func GetHistoricalData_Fyers(symbol, resolution, dateFormat, rangeFrom, rangeTo, UserID_Fyers string) ([]Candle, error) {

	AccessToken, err := ReadingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return []Candle{}, err
	}
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
