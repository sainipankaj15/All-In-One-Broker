package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// This API returns the position API response
func PositionApi_Tiqs(UserID_Tiqs string) (PositionAPIResp_Tiqs, error) {

	// Reading accessToken and APPID for fetching the APIs
	AccessToken, APPID, err := readingAccessToken_Tiqs(UserID_Tiqs)
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

	accessTokenofUser, appIdOfUser, err := readingAccessToken_Tiqs(UserID_Tiqs)
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
