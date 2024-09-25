package main

import (
	"fmt"
	"time"

	tiqs "github.com/sainipankaj15/All-In-One-Broker/Tiqs"
	tiqsGreeksSocket "github.com/sainipankaj15/All-In-One-Broker/TiqsGreeks"
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

	// a, _ := fyers.MarginMktOrder_Fyers("NSE:ITC-EQ", 1, 1, fyers.ProductType.INTRADAY, "XP03754")
	// fmt.Println(a.Data.MarginTotal)

	// a, _ := tiqs.LTPInPaisa_Tiqs(26009, "FB5650")
	// fmt.Println(a)

	// b, err := fyers.GetOptionChain_Fyers("NSE:NIFTYBANK-INDEX", 2, "XP03754")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v", b.Data.OptionsChain)

	// c, err := fyers.GetOptionChainMap_Fyers("NSE:NIFTYBANK-INDEX", 10, "XP03754")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v", c)
	// printOptionChainMap(c)
	// fyers.PrintOptionChainMap(c)

	// c, err := fyers.QuotesAPI_Fyers("NSE:SEACOAST-EQ", "XP03754")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v", c)

	// c, err := fyers.GetHistoricalData_Fyers("NSE:ITC-EQ", "1D", "1", "2021-01-01", "2021-01-10", "XP03754")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("%+v", c)

	// time.Sleep(50000 * time.Second)

	// d, err := tiqs.ClosestExpiryDate_Tiqs(tiqs.Index.BANKNIFTY, tiqs.ADMIN_TIQS)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Date is %+v", d)

	// d, err := tiqs.NextExpiryDateOnExpiry_Tiqs(tiqs.Index.NIFTY, tiqs.ADMIN_TIQS)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf(" date is %+v", d)

	// c := utils.RoundOff(52126 , 50)

	// fmt.Println("c is ", c)

	tokenId, appId, _ := tiqs.ReadingAccessToken_Tiqs("FB5650")

	gs, err := tiqsGreeksSocket.NewTiqsGreeksSocket(appId, tokenId, true)

	if err != nil {
		fmt.Println(err)
	}

	gs.StartWebSocket(tiqs.Index.MIDCPNIFTY, tiqs.ExchangeToken.MIDCPNIFTY)

	time.Sleep(5 * time.Second)

	fmt.Printf("gs is %+v", gs)
	time.Sleep(500 * time.Second)

	gs.PrintSyntheticFutureMap()

	// gs.PrintPriceMap()
	time.Sleep(15 * time.Second)

	gs.PrintSyntheticFutureMap()

	time.Sleep(15 * time.Second)

	gs.PrintSyntheticFutureMap()
	select {}

	// gs.PrintPriceMap()

	// for {
	// 	time.Sleep(1 * time.Second)

	// 	price, err := gs.GetPrice(40508)
	// 	if err != nil {
	// 		log.Printf("Error getting price: %v\n", err)
	// 	} else {
	// 		log.Printf("Price for token 40508: %.2f\n", price)
	// 	}
	// 	// fmt.Println("Hello World!")
	// }

	// select {}
}
