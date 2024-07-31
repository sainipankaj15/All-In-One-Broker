package main

import (
	"fmt"
	"time"

	fyers "github.com/sainipankaj15/All-In-One-Broker/Fyers"
)

func main() {
	fmt.Println("Hello World!")
	// time.Sleep(5 * time.Second)

	// allInOneBroker.ExitAllPosition_Tiqs("ABCD12")

	// token := []int{1, 36605}

	// resp, _ := tiqs.FetchQuotes_Tiqs(token, "FB5650")

	// log.Println("Response is here")
	// log.Println(resp.TokenData[36605].Ltp)
	// log.Printf("Hi Hi \n\n\n\n")

	// fyers_pos_resp, err := fyers.PositionApi_Fyers("XP03754")

	// if err != nil {
	// 	log.Println("Error while getting position API response")
	// 	log.Println(err)
	// } else {
	// 	// Printing the response
	// 	log.Println(fyers_pos_resp)
	// }

	// _ = fyers.ExitingAllPosition([]int{1, -1}, []int{10, 11, 12, 20}, []string{"INTRADAY", "MARGIN", "CO", "BO"}, "XP03754")
	// log.Println(resp)

	// exchnToken, err := fyers.SymbolNameToExchToken("ITC-EQ", "XP03754")

	// if err!= nil {
	// 	log.Println(err)
	// }
	// log.Println(exchnToken)

	// a, _ := fyers.QuotesAPI_Fyers("NSE:ITC-EQ", "XP03754")
	// a , _ := fyers.LTP_Fyers("NSE:ITC-EQ", "XP03754")
	// a, _ := fyers.MarketDepthAPI_Fyers("NSE:ITC-EQ", "XP03754")

	// _, _ = tiqs.ExitAllPosition_Tiqs("FB5650")

	// _ = tiqs.ExitByPositionID_Tiqs("MIDCPNIFTY22JUL24C12700", "M", "FB5650")

	// a, _, err := tiqs.GetOptionChain_Tiqs("26009", "1", "31-JUL-2024", "FB5650")
	// a , _ , err := tiqs.GetExpiryList_Tiqs("FB5650")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// b := a.Data["BANKNIFTY"]
	// c := b[0]
	// log.Println(c)

	// a, _ := tiqs.ClosestExpiryDate_Tiqs(tiqs.Index.FINNIFTY, tiqs.ADMIN_TIQS)
	// fmt.Println(a)

	a, _ := fyers.MarginMktOrder_Fyers("NSE:ITC-EQ", 1, 1, fyers.ProductType.INTRADAY, "XP03754")
	fmt.Println(a.Data.MarginTotal)

	time.Sleep(50000 * time.Second)

}
