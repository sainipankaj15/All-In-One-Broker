package tiqs_greeks_socket

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/alphadose/haxmap"
	tiqs "github.com/sainipankaj15/All-In-One-Broker/Tiqs"
	tiqsSocket "github.com/sainipankaj15/All-In-One-Broker/TiqsWS"
	typeConversion "github.com/sainipankaj15/data-type-conversion"
)

func NewTiqsGreeksSocket(appID string, accessToken string, enableLog bool) (*TiqsGreeksClient, error) {
	client := &TiqsGreeksClient{
		appID:       appID,
		accessToken: accessToken,
		enableLog:   enableLog,
		priceMap:    haxmap.New[int32, TickData](),
	}

	if enableLog {
		log.Printf("Initializing NewTiqsGreeksSocket with appID: %s and accessToken: %s", appID, accessToken)
		client.logger("TiqsGreeksClient created successfully")
	}

	return client, nil
}

func (t *TiqsGreeksClient) logger(msg ...any) {
	if t.enableLog {
		log.Println(msg...)
	}
}

func (t *TiqsGreeksClient) GetPriceMap() *haxmap.Map[int32, TickData] {
	return t.priceMap
}

func (t *TiqsGreeksClient) GetPrice(instrumentToken int32) (float64, error) {
	now := int32(time.Now().Unix())

	if tickData, ok := t.priceMap.Get(instrumentToken); ok {
		if now-tickData.Timestamp < 10 {
			return float64(tickData.LTP) / 100, nil
		}
	}

	// Fallback to API if recent price not available
	ltpInInt, err := tiqs.LTPInPaisa_Tiqs(int(instrumentToken), tiqs.ADMIN_TIQS)
	if err != nil {
		return 0, fmt.Errorf("failed to get price from API: %v", err)
	}

	ltp := int32(ltpInInt)
	t.priceMap.Set(instrumentToken, TickData{LTP: ltp, Timestamp: now})

	return float64(ltp) / 100, nil
}

func (t *TiqsGreeksClient) StartWebSocket(TargetSymbol string, TargetSymbolToken int) error {
	tiqsWs, err := tiqsSocket.NewTiqsWebSocket(t.appID, t.accessToken, t.enableLog)
	if err != nil {
		return fmt.Errorf("error while connecting tiqs socket: %w", err)
	}

	dataChannel := tiqsWs.GetDataChannel()

	go func() {
		for tick := range dataChannel {
			go t.priceMap.Set(tick.Token, TickData{LTP: tick.LTP, Timestamp: tick.Time})
		}
	}()

	// Subscribe to the target symbol token : Index Token
	tiqsWs.AddSubscription(TargetSymbolToken)

	// Subscribe to the option chain tokens : Strike Price Tokens
	optionChain, err := tiqs.GetOptionChainMap_Tiqs(TargetSymbol, strconv.Itoa(TargetSymbolToken), "25")
	if err != nil {
		return fmt.Errorf("error while getting option chain: %w", err)
	}

	fmt.Println("optionChain: ", optionChain)

	// Subscribe to the option chain tokens
	counter := 1
	for _, strikes := range optionChain {
		for _, symbol := range strikes {
			tokenInt := typeConversion.StringToInt(symbol.Token)
			if tokenInt == 0 {
				continue
			}

			tiqsWs.AddSubscription(tokenInt)
			t.logger(fmt.Sprintf("Subscribing to Symbol: %s, Token: %d, Total Symbols: %d", symbol.Name, tokenInt, counter))
			counter++
		}
	}

	t.logger(fmt.Sprintf("Total symbols subscribed: %d", counter-1))

	t.logger("WebSocket started successfully")
	return nil
}
