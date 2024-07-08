package main

import (
	"fmt"
	"log"

	fyers "github.com/sainipankaj15/All-In-One-Broker/Fyers"
	tiqs "github.com/sainipankaj15/All-In-One-Broker/tiqs"
)

func main() {
	fmt.Println("Hello World!")
	// time.Sleep(5 * time.Second)

	// allInOneBroker.ExitAllPosition_Tiqs("ABCD12")

	token := []int{1, 36605}

	resp, _ := tiqs.FetchQuotes_Tiqs(token, "FB5650")

	log.Println("Response is here")
	log.Println(resp.TokenData[36605].Ltp)
	log.Printf("Hi Hi \n\n\n\n")

	fyers_pos_resp, err := fyers.PositionApi_Fyers("XP03754")

	if err != nil {
		log.Println("Error while getting position API response")
		log.Println(err)
	} else {
		// Printing the response
		log.Println(fyers_pos_resp)
	}
}
