package main

import (
	"fmt"
	"log"

	allInOneBroker "github.com/sainipankaj15/All-In-One-Broker"
)

func main() {
	fmt.Println("Hello World!")
	// time.Sleep(5 * time.Second)

	// allInOneBroker.ExitAllPosition_Tiqs("ABCD12")

	token := []int{1, 36605}

	resp, _ := allInOneBroker.FetchQuotes_Tiqs(token, "FB5650")

	log.Println("Response is here")
	log.Println(resp.TokenData[36605].Ltp)
}
