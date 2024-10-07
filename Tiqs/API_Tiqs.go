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

// This API returns the position API response
func PositionApi_Tiqs(UserID_Tiqs string) (PositionAPIResp_Tiqs, error) {

	// Reading accessToken and APPID for fetching the APIs
	AccessToken, APPID, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		log.Println("Error while getting acces token from file")
		panic(err)
	}

	positionUrl := "https://api.tiqs.in/oms/user/positions"

	req, err := http.NewRequest("GET", positionUrl, nil)
	if err != nil {
		log.Println("Error while making request in Position API request")
		return PositionAPIResp_Tiqs{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("token", AccessToken)
	req.Header.Add("appId", APPID)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in Position API")
		return PositionAPIResp_Tiqs{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in Position API")
		return PositionAPIResp_Tiqs{}, err
	}

	jsonBody := string(body)
	log.Printf("Direct Response from Position API of Tiqs %v", jsonBody)

	// Converting into Response struct format
	var positionResp PositionAPIResp_Tiqs

	err = json.Unmarshal(body, &positionResp)
	if err != nil {
		log.Println("Error while Unmarshaling the data in Position API")
		return PositionAPIResp_Tiqs{}, err
	}

	return positionResp, nil
}

// This API is used to place market for Tiqs Broker
func OrderPlaceMarket_Tiqs(exchange, token, quantity, TransSide, productType, UserID_Tiqs string) error {

	accessTokenofUser, appIdOfUser, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		msg := fmt.Sprintf("Error while getting access token from file for %v User ", UserID_Tiqs)
		log.Println(msg)
		return err
	}

	// orderPlacment URL
	url := "https://api.tiqs.in/oms/order/regular"

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
		"tags":            "Bhai_Kal_se_Pakka_trade_Nahi_karunga",
		"symbol":          "it_doesn't_matter",
	}

	jsonParameters, err := json.Marshal(values)
	if err != nil {
		log.Println("Error marshaling JSON in orderPlacement_Tiqs()", err)
		return err
	}

	// Create HTTP client
	client := &http.Client{}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonParameters))
	if err != nil {
		log.Println("Error while making request for Order Placement API ")
		return err
	}

	// Add the Session and token to the request header
	req.Header.Add("appId", appIdOfUser)
	req.Header.Add("token", accessTokenofUser)
	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while getting response in orderPlacement_Tiqs()")
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in orderPlacement_Tiqs()")
		return err
	}

	jsonBody := string(body)

	msg := fmt.Sprintf("For %v user , response status is %v and response is %v", UserID_Tiqs, resp.Status, jsonBody)
	log.Println(msg)
	return nil
}

// fetchQuotes sends a POST request with the specified data and userID, then returns the response body as a string and any error encountered.
func FetchQuotes_Tiqs(tokenSlice []int, UserID_Tiqs string) (QuotesAPIResp_Tiqs, error) {

	quotesUrl := "https://api.tiqs.trading/info/quotes/full"

	accessTokenofUser, appIdOfUser, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		msg := fmt.Sprintf("Error while getting access token from file for %v User ", UserID_Tiqs)
		log.Println(msg)
		return QuotesAPIResp_Tiqs{}, err
	}

	// Convert []int to []byte
	jsonData, err := json.Marshal(tokenSlice)
	if err != nil {
		log.Println("Error while Marshalling token slice", err)
		return QuotesAPIResp_Tiqs{}, err
	}

	// Create a new request using http
	req, err := http.NewRequest("POST", quotesUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return QuotesAPIResp_Tiqs{}, fmt.Errorf("error creating request: %v", err)
	}

	// Add the Session and token to the request header
	req.Header.Add("appId", appIdOfUser)
	req.Header.Add("token", accessTokenofUser)
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return QuotesAPIResp_Tiqs{}, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return QuotesAPIResp_Tiqs{}, fmt.Errorf("error reading response body: %v", err)
	}

	log.Println(string(body))

	var apiResponse QuotesAPIResp_Tiqs
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error while unmarshaling JSON of Tiqs Quotes API", err)
		return QuotesAPIResp_Tiqs{}, err
	}

	// Populate the map
	apiResponse.TokenData = make(map[int]QuotesData_Tiqs)
	for _, data := range apiResponse.Data {
		apiResponse.TokenData[data.Token] = data
	}

	return apiResponse, nil
}

func GetOptionChain_Tiqs(IndexTokenNumber, OptionChainLength, expiryDay, UserID_Tiqs string) (OptionChainResp_Tiqs, int, error) {

	getOptionChainUrl := "https://api.tiqs.trading/info/option-chain"

	log.Printf("GetOptionChain_Tiqs API for %v token , option chain length is %v and expiryDay is %v ", IndexTokenNumber, OptionChainLength, expiryDay)

	accessTokenofUser, appIdOfUser, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		msg := fmt.Sprintf("Error while getting access token from file for %v User ", UserID_Tiqs)
		log.Println(msg)
		return OptionChainResp_Tiqs{}, 0, err
	}

	values := map[string]string{"token": IndexTokenNumber, "exchange": "INDEX", "count": OptionChainLength, "expiry": expiryDay}
	jsonParameters, err := json.Marshal(values)
	if err != nil {
		log.Println("Error marshaling JSON in GetOptionChain_Tiqs(): ", err)
		return OptionChainResp_Tiqs{}, 0, err
	}

	req, err := http.NewRequest("POST", getOptionChainUrl, bytes.NewBuffer(jsonParameters))
	if err != nil {
		log.Println("Error while making request in GetOptionChain_Tiqs()")
		return OptionChainResp_Tiqs{}, 0, err
	}

	// Add the Session and token to the request header : Here session will be APPID and Token will be token
	req.Header.Add("appId", appIdOfUser)
	req.Header.Add("token", accessTokenofUser)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while getting response in GetOptionChain_Tiqs()")
		return OptionChainResp_Tiqs{}, resp.StatusCode, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in GetOptionChain_Tiqs()")
		return OptionChainResp_Tiqs{}, resp.StatusCode, err
	}

	jsonBody := string(body)
	// Converting into Response struct format
	var Resp OptionChainResp_Tiqs
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		log.Println("Error while Unmarshaling the data in GetOptionChain_Tiqs()", err)
		return OptionChainResp_Tiqs{}, resp.StatusCode, err
	}

	msg := fmt.Sprintf("Direct Response in GetOptionChain_Tiqs() is %v", jsonBody)

	log.Println(msg)
	return Resp, resp.StatusCode, nil
}
func GetExpiryList_Tiqs(UserID_Tiqs string) (ExpiryResp_Tiqs, int, error) {

	expiryDayListUrl := "https://api.tiqs.trading/info/option-chain-symbols"

	accessTokenofUser, appIdOfUser, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		msg := fmt.Sprintf("Error while getting access token from file for %v User ", UserID_Tiqs)
		log.Println(msg)
		return ExpiryResp_Tiqs{}, 0, err
	}

	values := map[string]string{}
	jsonParameters, err := json.Marshal(values)
	if err != nil {
		log.Println("Error marshaling JSON in GetExpiryList_Tiqs(): ", err)
		return ExpiryResp_Tiqs{}, 0, err
	}

	req, err := http.NewRequest("GET", expiryDayListUrl, bytes.NewBuffer(jsonParameters))
	if err != nil {
		log.Println("Error while making request in GetExpiryList_Tiqs()")
		return ExpiryResp_Tiqs{}, 0, err
	}

	// Add the Session and token to the request header : Here session will be APPID and Token will be token
	req.Header.Add("appId", appIdOfUser)
	req.Header.Add("token", accessTokenofUser)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while getting response in GetExpiryList_Tiqs()")
		return ExpiryResp_Tiqs{}, resp.StatusCode, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in byte array in GetExpiryList_Tiqs()")
		return ExpiryResp_Tiqs{}, resp.StatusCode, err
	}

	jsonBody := string(body)
	// Converting into Response struct format
	var Resp ExpiryResp_Tiqs
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		log.Println("Error while Unmarshaling the data in GetExpiryList_Tiqs()", err)
		return ExpiryResp_Tiqs{}, resp.StatusCode, err
	}

	msg := fmt.Sprintf("Direct Response in GetExpiryList_Tiqs() is %v", jsonBody)

	if resp.StatusCode != 200 {
		log.Println(msg)
		err := errors.New("token Expired")
		return ExpiryResp_Tiqs{}, resp.StatusCode, err
	}
	log.Println(msg)
	return Resp, resp.StatusCode, nil
}

func LTPInPaisa_Tiqs(tokenNumber int, UserID_Tiqs string) (int, error) {

	// Reading accessToken and APPID for fetching the APIs
	AccessToken, APPID, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		log.Println("Error while getting acces token from file")
		panic(err)
	}

	url := "https://api.tiqs.trading/info/quote/ltp"

	// Create a map for the JSON data
	data := map[string]int{
		"token": tokenNumber,
	}

	// Convert the map to JSON
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error while making request in LTPOfToken_Tiqs request")
		return 0, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("token", AccessToken)
	req.Header.Add("appId", APPID)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
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
	var apiResp LTPofTokenResp_Tiqs

	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Println("Error while Unmarshaling the data in Position API")
		return 0, err
	}

	ltp := apiResp.Data.LTP

	return ltp, nil
}

func GetGreeks_Tiqs(tokenNumber int, UserID_Tiqs string) (GreeksData_Tiqs, error) {
	// Reading accessToken and APPID for fetching the APIs
	AccessToken, APPID, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		log.Println("Error while getting access token from file")
		return GreeksData_Tiqs{}, err
	}

	url := "https://api.tiqs.trading/info/greeks"

	// Create a slice with the token number
	data := []int{tokenNumber}

	// Convert the slice to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error while marshaling data in GetGreeks_Tiqs")
		return GreeksData_Tiqs{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error while making request in GetGreeks_Tiqs")
		return GreeksData_Tiqs{}, err
	}

	// Add the headers to the request
	req.Header.Add("appId", APPID)
	req.Header.Add("token", AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in GetGreeks_Tiqs API")

		return GreeksData_Tiqs{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in GetGreeks_Tiqs API")
		return GreeksData_Tiqs{}, err
	}

	jsonBody := string(body)
	log.Printf("Direct Response from GetGreeks_Tiqs API: %v", jsonBody)

	// Converting into Response struct format
	var apiResp GreeksResp_Tiqs

	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Println("Error while Unmarshaling the data in GetGreeks_Tiqs API")
		return GreeksData_Tiqs{}, err
	}

	if apiResp.Status != "success" {
		return GreeksData_Tiqs{}, fmt.Errorf("API returned non-success status: %s", apiResp.Status)
	}

	if len(apiResp.Data) == 0 {
		return GreeksData_Tiqs{}, fmt.Errorf("no data returned for the given token")
	}

	greeksData := apiResp.Data[0]
	greeksData.IV *= 100
	return greeksData, nil
}

func GetHolidays_Tiqs(UserID_Tiqs string) (HolidaysAPIResp_Tiqs, error) {
	// Reading accessToken and APPID for fetching the APIs
	AccessToken, APPID, err := ReadingAccessToken_Tiqs(UserID_Tiqs)
	if err != nil {
		log.Println("Error while getting access token from file")
		return HolidaysAPIResp_Tiqs{}, err
	}

	holidaysUrl := "https://api.tiqs.trading/info/holidays"

	// Create a new request using http
	req, err := http.NewRequest("GET", holidaysUrl, nil)
	if err != nil {
		log.Println("Error while making request in GetHolidays_Tiqs")
		return HolidaysAPIResp_Tiqs{}, err
	}

	// Add the Bearer token to the request header
	req.Header.Add("token", AccessToken)
	req.Header.Add("appId", APPID)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while making request in GetHolidays_Tiqs")
		return HolidaysAPIResp_Tiqs{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the body in GetHolidays_Tiqs")
		return HolidaysAPIResp_Tiqs{}, err
	}

	// Converting into Response struct format
	var holidaysResp HolidaysAPIResp_Tiqs
	err = json.Unmarshal(body, &holidaysResp)
	if err != nil {
		log.Println("Error while Unmarshaling the data in GetHolidays_Tiqs")
		return HolidaysAPIResp_Tiqs{}, err
	}

	return holidaysResp, nil
}
