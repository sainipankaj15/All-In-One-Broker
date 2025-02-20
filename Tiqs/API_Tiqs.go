package tiqs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// PositionApi_Tiqs returns the response of the position API. It takes the UserID of the user as an argument and returns the response and an error.
func PositionApi_Tiqs(UserID_Tiqs string) (PositionAPIResp_Tiqs, error) {

	// Reading accessToken and APPID for fetching the APIs
	AccessToken, APPID, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while getting acces token from file")
		return PositionAPIResp_Tiqs{}, err
	}

	req, err := http.NewRequest("GET", positionUrl, nil)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while making request in Position API request")
		return PositionAPIResp_Tiqs{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("token", AccessToken)
	req.Header.Add("appId", APPID)

	// Make the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while making request in Position API")
		return PositionAPIResp_Tiqs{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while reading the body in byte array in Position API")
		return PositionAPIResp_Tiqs{}, err
	}

	// Converting into Response struct format
	var positionResp PositionAPIResp_Tiqs

	err = json.Unmarshal(body, &positionResp)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while Unmarshaling the data in Position API")
		return PositionAPIResp_Tiqs{}, err
	}

	// Return the response and nil error
	return positionResp, nil
}

// OrderPlaceMarket_Tiqs is used to place a market order for Tiqs Broker
// It takes exchange, token, quantity, TransSide, productType and UserID_Tiqs as parameters
// It returns an error if something goes wrong
func OrderPlaceMarket_Tiqs(exchange, token, quantity, TransSide, productType, UserID_Tiqs string) (placeOrderResp_Tiqs, error) {

	accessTokenofUser, appIdOfUser, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		return placeOrderResp_Tiqs{}, fmt.Errorf("error while getting access token from file for %v User: %w", UserID_Tiqs, err)
	}

	// Create the JSON to be sent in the body of the request
	values := map[string]string{
		"exchange":        exchange,
		"token":           token,
		"quantity":        quantity,
		"disclosedQty":    "0",
		"product":         productType,
		"transactionType": TransSide,
		"order":           "MKT",
		"price":           "0",
		"validity":        "DAY",
		"triggerPrice":    "0",
		"tags":            "Bhai_Kal_se_Pakka_trade_Nahi_karunga", // This is a tag that will be added to the order
		"symbol":          "it_doesn't_matter",                    // This is a symbol that will be added to the order
	}

	jsonParameters, err := json.Marshal(values)
	if err != nil {
		log.Println("Error marshaling JSON in orderPlacement_Tiqs()", err)
		return placeOrderResp_Tiqs{}, err
	}

	// Create HTTP client
	client := http.DefaultClient

	// Create HTTP request
	req, err := http.NewRequest("POST", placeOrderUrl, bytes.NewBuffer(jsonParameters))
	if err != nil {
		log.Println("Error while making request for Order Placement API ")
		return placeOrderResp_Tiqs{}, err
	}

	// Add the Session and token to the request header
	req.Header.Add("appId", appIdOfUser)
	req.Header.Add("token", accessTokenofUser)
	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while getting response in orderPlacement_Tiqs()")
		return placeOrderResp_Tiqs{}, err
	}
	defer resp.Body.Close()

	var response placeOrderResp_Tiqs
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return placeOrderResp_Tiqs{}, err
	}

	if response.Status != apiResponseStatus.SUCCESS {
		return placeOrderResp_Tiqs{}, fmt.Errorf("%w, response: %+v", ErrOrderPlacementFailed, response)
	}

	return response, nil
}

// fetchQuotes sends a POST request with the specified data and userID, then returns the response body as a string and any error encountered.
func FetchQuotes_Tiqs(tokenSlice []int, UserID_Tiqs string) (quotesAPIResp_Tiqs, error) {

	accessTokenofUser, appIdOfUser, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		msg := fmt.Sprintf("Error while getting access token from file for %v User ", UserID_Tiqs)
		log.Println(msg)
		return quotesAPIResp_Tiqs{}, err
	}

	// Convert []int to []byte
	jsonData, err := json.Marshal(tokenSlice)
	if err != nil {
		log.Println("Error while Marshalling token slice", err)
		return quotesAPIResp_Tiqs{}, err
	}

	// Create a new request using http
	req, err := http.NewRequest("POST", quotesUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return quotesAPIResp_Tiqs{}, fmt.Errorf("error creating request: %v", err)
	}

	// Add the Session and token to the request header
	req.Header.Add("appId", appIdOfUser)
	req.Header.Add("token", accessTokenofUser)
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return quotesAPIResp_Tiqs{}, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return quotesAPIResp_Tiqs{}, fmt.Errorf("error reading response body: %v", err)
	}

	var apiResponse quotesAPIResp_Tiqs
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error while unmarshaling JSON of Tiqs Quotes API", err)
		return quotesAPIResp_Tiqs{}, err
	}

	// Populate the map
	apiResponse.TokenData = make(map[int]quotesData_Tiqs)
	for _, data := range apiResponse.Data {
		apiResponse.TokenData[data.Token] = data
	}

	return apiResponse, nil
}

// GetOptionChain_Tiqs fetches the option chain for a given token number, option chain length, and expiry day.
// It returns the option chain response, the status code of the response, and any error encountered.
func GetOptionChain_Tiqs(IndexTokenNumber, OptionChainLength, expiryDay, UserID_Tiqs string) (optionChainResp_Tiqs, int, error) {

	log.Printf("GetOptionChain_Tiqs API for %v token , option chain length is %v and expiryDay is %v ", IndexTokenNumber, OptionChainLength, expiryDay)

	accessTokenofUser, appIdOfUser, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		msg := fmt.Sprintf("Error while getting access token from file for %v User ", UserID_Tiqs)
		log.Println(msg)
		return optionChainResp_Tiqs{}, 0, err
	}

	values := map[string]string{"token": IndexTokenNumber, "exchange": "INDEX", "count": OptionChainLength, "expiry": expiryDay}
	jsonParameters, err := json.Marshal(values)
	if err != nil {
		log.Println("Error marshaling JSON in GetOptionChain_Tiqs(): ", err)
		return optionChainResp_Tiqs{}, 0, err
	}

	req, err := http.NewRequest("POST", getOptionChainUrl, bytes.NewBuffer(jsonParameters))
	if err != nil {
		log.Println("Error while making request in GetOptionChain_Tiqs()")
		return optionChainResp_Tiqs{}, 0, err
	}

	// Add the Session and token to the request header : Here session will be APPID and Token will be token
	req.Header.Add("appId", appIdOfUser)
	req.Header.Add("token", accessTokenofUser)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while getting response in GetOptionChain_Tiqs()")
		return optionChainResp_Tiqs{}, resp.StatusCode, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in GetOptionChain_Tiqs()")
		return optionChainResp_Tiqs{}, resp.StatusCode, err
	}

	jsonBody := string(body)
	// Converting into Response struct format
	var Resp optionChainResp_Tiqs
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		log.Println("Error while Unmarshaling the data in GetOptionChain_Tiqs()", err)
		return optionChainResp_Tiqs{}, resp.StatusCode, err
	}

	msg := fmt.Sprintf("Direct Response in GetOptionChain_Tiqs() is %v", jsonBody)

	log.Println(msg)
	return Resp, resp.StatusCode, nil
}

// GetExpiryList_Tiqs fetches the list of expiry dates for options from the Tiqs API.
// It takes a user ID as a parameter and returns the expiry response, status code, and any error encountered.
func GetExpiryList_Tiqs(UserID_Tiqs string) (ExpiryResp_Tiqs, error) {

	// Read access token and APPID for the user
	accessTokenofUser, appIdOfUser, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		log.Printf("Error while getting access token from file for %v User", UserID_Tiqs)
		return ExpiryResp_Tiqs{}, err
	}

	// Create an empty map to marshal into JSON for the request body
	values := map[string]string{}
	jsonParameters, err := json.Marshal(values)
	if err != nil {
		log.Println("Error marshaling JSON in GetExpiryList_Tiqs():", err)
		return ExpiryResp_Tiqs{}, err
	}

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", expiryDayListUrl, bytes.NewBuffer(jsonParameters))
	if err != nil {
		log.Println("Error while making request in GetExpiryList_Tiqs()")
		return ExpiryResp_Tiqs{}, err
	}

	// Add the Session and token to the request header: Here session will be APPID and Token will be token
	req.Header.Add("appId", appIdOfUser)
	req.Header.Add("token", accessTokenofUser)
	req.Header.Set("Content-Type", "application/json")

	// Send the request using the default HTTP client
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while getting response in GetExpiryList_Tiqs()")
		return ExpiryResp_Tiqs{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in GetExpiryList_Tiqs()")
		return ExpiryResp_Tiqs{}, err
	}

	// Convert the response body to a string for logging
	jsonBody := string(body)
	// Unmarshal the JSON response into the ExpiryResp_Tiqs struct
	var response ExpiryResp_Tiqs
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Error while Unmarshaling the data in GetExpiryList_Tiqs()", err)
		return ExpiryResp_Tiqs{}, err
	}

	// Check if the status code is not 200 (OK), indicating an error
	if resp.StatusCode != 200 {
		log.Printf("Error response in GetExpiryList_Tiqs(): %v", jsonBody)
		return ExpiryResp_Tiqs{}, errors.New("token expired")
	}

	// Return the unmarshaled response and status code
	return response, nil
}

// LTPInPaisa_Tiqs fetches the LTP of a given token number from the Tiqs API.
// It takes the token number and the UserID of the user as parameters.
// It returns the LTP of the token in paisa and an error if something goes wrong.
func LTPInPaisa_Tiqs(tokenNumber int, UserID_Tiqs string) (int, error) {

	// Reading accessToken and APPID for fetching the APIs
	AccessToken, APPID, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		log.Println("Error while getting acces token from file")
		return 0, err
	}

	// Create a map for the JSON data
	data := map[string]int{
		"token": tokenNumber,
	}

	// Convert the map to JSON
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", ltpUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error while making request in LTPOfToken_Tiqs request")
		return 0, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("token", AccessToken)
	req.Header.Add("appId", APPID)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in LTPOfToken_Tiqs API")
		return 0, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in LTPOfToken_Tiqs API")
		return 0, err
	}

	jsonBody := string(body)
	log.Printf("Direct Response from LTPOfToken_Tiqs API of Tiqs %v", jsonBody)

	// Converting into Response struct format
	var apiResp ltpofTokenResp_Tiqs

	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Println("Error while Unmarshaling the data in Position API")
		return 0, err
	}

	ltp := apiResp.Data.LTP

	return ltp, nil
}

// GetGreeks_Tiqs fetches the Greeks (Delta, Gamma, Theta, and Vega) for a given token number.
// It takes a token number and a UserID_Tiqs as parameters and returns the Greeks data and an error if something goes wrong.
func GetGreeks_Tiqs(tokenNumber int, UserID_Tiqs string) (greeksData_Tiqs, error) {
	// Reading accessToken and APPID for fetching the APIs
	AccessToken, APPID, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		log.Println("Error while getting access token from file")
		return greeksData_Tiqs{}, err
	}

	// Create a slice with the token number
	data := []int{tokenNumber}

	// Convert the slice to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error while marshaling data in GetGreeks_Tiqs")
		return greeksData_Tiqs{}, err
	}

	req, err := http.NewRequest("POST", greeksUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error while making request in GetGreeks_Tiqs")
		return greeksData_Tiqs{}, err
	}

	// Add the headers to the request
	req.Header.Add("appId", APPID)
	req.Header.Add("token", AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in GetGreeks_Tiqs API")

		return greeksData_Tiqs{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in GetGreeks_Tiqs API")
		return greeksData_Tiqs{}, err
	}

	jsonBody := string(body)
	log.Printf("Direct Response from GetGreeks_Tiqs API: %v", jsonBody)

	// Converting into Response struct format
	var apiResp greeksResp_Tiqs

	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Println("Error while Unmarshaling the data in GetGreeks_Tiqs API")
		return greeksData_Tiqs{}, err
	}

	if apiResp.Status != "success" {
		return greeksData_Tiqs{}, fmt.Errorf("API returned non-success status: %s", apiResp.Status)
	}

	if len(apiResp.Data) == 0 {
		return greeksData_Tiqs{}, fmt.Errorf("no data returned for the given token")
	}

	greeksData := apiResp.Data[0]
	greeksData.IV *= 100
	return greeksData, nil
}

// GetHolidays_Tiqs fetches the holidays from the Tiqs API.
// It takes a UserID_Tiqs as a parameter and returns the response and an error if something goes wrong.
func GetHolidays_Tiqs(UserID_Tiqs string) (holidaysAPIResp_Tiqs, error) {
	// Reading accessToken and APPID for fetching the APIs
	AccessToken, APPID, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		log.Println("Error while getting access token from file")
		return holidaysAPIResp_Tiqs{}, err
	}

	// Create a new request using http
	req, err := http.NewRequest("GET", holidaysUrl, nil)
	if err != nil {
		log.Println("Error while making request in GetHolidays_Tiqs")
		return holidaysAPIResp_Tiqs{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("token", AccessToken)
	req.Header.Add("appId", APPID)

	// Make the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in GetHolidays_Tiqs")
		return holidaysAPIResp_Tiqs{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in GetHolidays_Tiqs")
		return holidaysAPIResp_Tiqs{}, err
	}

	// Converting into Response struct format
	var holidaysResp holidaysAPIResp_Tiqs
	err = json.Unmarshal(body, &holidaysResp)
	if err != nil {
		log.Println("Error while Unmarshaling the data in GetHolidays_Tiqs")
		return holidaysAPIResp_Tiqs{}, err
	}

	return holidaysResp, nil
}

// GetOrderStatus_Tiqs fetches the status of an order from the Tiqs API.
// It takes the order ID and the UserID of the user as parameters.
// It returns the status of the order and an error if something goes wrong.
func GetOrderStatus_Tiqs(orderID string, UserID_Tiqs string) (string, error) {

	// Reading accessToken and APPID for fetching the APIs
	AccessToken, APPID, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while getting acces token from file")
		return "", err
	}

	// Create the URL for the API request
	url := orderBookURL + "/" + orderID

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while making request in GetOrderStatus_Tiqs request")
		return "", err
	}

	// Add the Bearer token to the request header
	req.Header.Add("token", AccessToken)
	req.Header.Add("appId", APPID)

	// Make the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while making request in GetOrderStatus_Tiqs API")
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while reading the body in byte array in GetOrderStatus_Tiqs API")
		return "", err
	}

	// Converting into Response struct format
	var orderBookResp OrderBookAPIResp_Tiqs

	err = json.Unmarshal(body, &orderBookResp)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while Unmarshaling the data in GetOrderStatus_Tiqs API")
		return "", err
	}

	if orderBookResp.Status != apiResponseStatus.SUCCESS {
		// If the API returned non-success status, return an error
		return "", errors.New("API returned non-success status")
	}

	// Return the status of the order
	status := orderBookResp.OrderBook[0].OrderStatus
	return status, nil
}

// OrderBookApi_Tiqs returns the response of the OrderBook API. It takes the UserID of the user as an argument and returns the response and an error.
func OrderBookApi_Tiqs(UserID_Tiqs string) (OrderBookAPIResp_Tiqs, error) {

	// Reading accessToken and APPID for fetching the APIs
	AccessToken, APPID, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while getting acces token from file")
		return OrderBookAPIResp_Tiqs{}, err
	}

	req, err := http.NewRequest("GET", orderBookURL, nil)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while making request in OrderBook API request")
		return OrderBookAPIResp_Tiqs{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("token", AccessToken)
	req.Header.Add("appId", APPID)

	// Make the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while making request in OrderBook API")
		return OrderBookAPIResp_Tiqs{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while reading the body in byte array in OrderBook API")
		return OrderBookAPIResp_Tiqs{}, err
	}

	// Converting into Response struct format
	var orderBookResp OrderBookAPIResp_Tiqs

	err = json.Unmarshal(body, &orderBookResp)
	if err != nil {
		// Log the error and return an error
		log.Println("Error while Unmarshaling the data in Position API")
		return OrderBookAPIResp_Tiqs{}, err
	}

	// Return the response and nil error
	return orderBookResp, nil
}
