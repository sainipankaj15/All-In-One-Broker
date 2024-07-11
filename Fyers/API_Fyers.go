package fyers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PositionApi_Fyers(UserID_Fyers string) (PositionAPIResp_Fyers, error) {

	AccessToken, err := readingAccessToken_Fyers(UserID_Fyers)
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

	AccessToken, err := readingAccessToken_Fyers(UserID_Fyers)
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

	AccessToken, err := readingAccessToken_Fyers(UserID_Fyers)
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

func MarketDepth_Fyers(symbolName string, UserID_Fyers string) (MarketDepthAPI_Fyers, error) {

	AccessToken, err := readingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return MarketDepthAPI_Fyers{}, err
	}

	url := fmt.Sprintf("https://api-t1.fyers.in/data/depth?symbol=%s&ohlcv_flag=1", symbolName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error while making request in marketDepthAPI")
		return MarketDepthAPI_Fyers{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in marketDepthAPI")
		return MarketDepthAPI_Fyers{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in marketDepthAPI")
		return MarketDepthAPI_Fyers{}, err
	}

	jsonBody := string(body)
	log.Printf("Direct Response from Market Depth API of fyers for %v is %v", symbolName, jsonBody)

	// Converting into Response struct format
	var marketDepthResponse MarketDepthAPI_Fyers

	err = json.Unmarshal(body, &marketDepthResponse)
	if err != nil {
		log.Println("Error while Unmarshaling the data in marketDepthAPI")
		return MarketDepthAPI_Fyers{}, err
	}

	return marketDepthResponse, nil
}
func LTP_Fyers(symbolName string, UserID_Fyers string) (float64, error) {

	AccessToken, err := readingAccessToken_Fyers(UserID_Fyers)
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
	var marketDepthResponse MarketDepthAPI_Fyers

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

func PlaceOrder_Fyers(symbolName string, LimitPriceForOrder float64, qty int, UserID_Fyers string, whichSide int) (bool, error) {

	AccessToken, err := readingAccessToken_Fyers(UserID_Fyers)
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
		"productType":  "MARGIN",
		"limitPrice":   0,
		"disclosedQty": 0,
		"stopPrice":    0,
		"validity":     "DAY",
		"offlineOrder": false,
		"stopLoss":     0,
		"takeProfit":   0,
		"orderTag":     "fiveEMAKaOrder",
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

func QuotesAPI_Fyers(symbolName, UserID_Fyers string) (QuoteAPI_Fyers, error) {

	AccessToken, err := readingAccessToken_Fyers(UserID_Fyers)
	if err != nil {
		log.Fatalf("Error while getting access token in Fyers")
		return QuoteAPI_Fyers{}, err
	}

	url := fmt.Sprintf("https://api-t1.fyers.in/data/quotes?symbols=%s", symbolName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error while making request in quotesFyersAPI")
		return QuoteAPI_Fyers{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("Authorization", AccessToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in quotesFyersAPI")
		return QuoteAPI_Fyers{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in quotesFyersAPI")
		return QuoteAPI_Fyers{}, err
	}

	jsonBody := string(body)
	log.Printf("Direct Response from Quotes API of fyers for %v is %v", symbolName, jsonBody)

	// Converting into Response struct format
	var qpr QuoteAPI_Fyers

	err = json.Unmarshal(body, &qpr)
	if err != nil {
		log.Println("Error while Unmarshaling the data in quotes Fyers API")
		return QuoteAPI_Fyers{}, err
	}

	return qpr, nil
}

func SymbolNameToExchToken(symbolName, UserID_Fyers string) (string, error) {

	AccessToken, err := readingAccessToken_Fyers(UserID_Fyers)
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
	var qpr QuoteAPI_Fyers

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
		return "", err
	}

	// Remove first 10 characters
	// For more info read this : https://myapi.fyers.in/docsv3#tag/Appendix/Fytoken
	exchangeToken := fytoken[10:]
	return exchangeToken, nil
}
