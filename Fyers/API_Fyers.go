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
