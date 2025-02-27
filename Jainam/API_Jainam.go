package jainam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func OrderPlaceMarket_Jainam(exchange, token, quantity, TransSide, productType, UserID_Tiqs string) (placeOrderResp_Jainam, error) {

	accessTokenofUser, appIdOfUser, err := ReadingAccessToken_Jainam(UserID_Tiqs)
	if err != nil {
		return placeOrderResp_Jainam{}, err
	}

	// Create the JSON to be sent in the body of the request
	values := placeOrderReq_Jainam{
		Exchange:      exchange,
		Qty:           quantity,
		Price:         "1",
		Product:       productType,
		TransType:     TransSide,
		PriceType:     OrderType.MARKET,
		TriggerPrice:  0,
		Ret:           "DAY",
		DisclosedQty:  0,
		MktProtection: "",
		Target:        1,
		StopLoss:      1,
		OrderType:     "Regular",
		Token:         token,
	}

	jsonParameters, err := json.Marshal(values)
	if err != nil {
		log.Println("Error marshaling JSON in orderPlacement_Tiqs()", err)
		return placeOrderResp_Jainam{}, err
	}

	// Create HTTP client
	client := http.DefaultClient

	// Create HTTP request
	req, err := http.NewRequest("POST", placeOrderUrl, bytes.NewBuffer(jsonParameters))
	if err != nil {
		log.Println("Error while making request for Order Placement API ")
		return placeOrderResp_Jainam{}, err
	}

	// Add the Session and token to the request header
	req.Header.Add("appId", appIdOfUser)
	req.Header.Add("token", accessTokenofUser)
	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while getting response in orderPlacement_Tiqs()")
		return placeOrderResp_Jainam{}, err
	}
	defer resp.Body.Close()

	var response placeOrderResp_Jainam
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return placeOrderResp_Jainam{}, err
	}

	if response[0].Status != "Success" {
		return response, fmt.Errorf("order Placement Failed: %s", response[0].Message)
	}

	return response, nil
}
