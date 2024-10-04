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
		appID:                   appID,
		accessToken:             accessToken,
		enableLog:               enableLog,
		priceMap:                haxmap.New[int32, TickData](), // token to tick data
		strikeToSyntheticFuture: haxmap.New[int32, float64](),  // strike price to synthetic future price
		peTokenToCeToken:        haxmap.New[int32, int32](),    // PE token to corresponding CE token
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

// GetTickData retrieves the full TickData for a given token
func (t *TiqsGreeksClient) GetTickData(token int32) (TickData, error) {
	if tickData, ok := t.priceMap.Get(token); ok {
		return tickData, nil
	}
	
	// If not found in the map, fetch from API and update the map
	ltpInPaisa, err := tiqs.LTPInPaisa_Tiqs(int(token), tiqs.ADMIN_TIQS)
	if err != nil {
		return TickData{}, fmt.Errorf("failed to get price from API: %v", err)
	}

	now := int32(time.Now().Unix())
	newTickData := TickData{
		LTP:       int32(ltpInPaisa),
		Timestamp: now,
		// Note: Other fields like StrikePrice are not set here
		// as we don't have that information from just the LTP API call
	}

	t.priceMap.Set(token, newTickData)
	return newTickData, nil
}


// GetPriceMap returns the internal price map of the TiqsGreeksClient.
// This map contains the latest tick data for each instrument token.
func (t *TiqsGreeksClient) GetPriceMap() *haxmap.Map[int32, TickData] {
	return t.priceMap
}

// GetSyntheticFutureMap returns the internal map of synthetic future prices.
// This map contains the calculated synthetic future price for each strike price.
func (t *TiqsGreeksClient) GetSyntheticFutureMap() *haxmap.Map[int32, float64] {
	return t.strikeToSyntheticFuture
}

// GetPrice retrieves the latest price for a given instrument token.
// It first checks the internal cache (priceMap) for recent data.
// If recent data is not available, it falls back to fetching from the API.
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

	// setting the time to expiry in Days
	err := t.settingTimeToExpiry(TargetSymbol)
	if err != nil {
		return fmt.Errorf("error while setting time to expiry: %w", err)
	}

	tiqsWs, err := tiqsSocket.NewTiqsWebSocket(t.appID, t.accessToken, t.enableLog)
	if err != nil {
		return fmt.Errorf("error while connecting tiqs socket: %w", err)
	}

	dataChannel := tiqsWs.GetDataChannel()

	go func() {
		for tick := range dataChannel {
			if val, ok := t.priceMap.Get(tick.Token); ok {
				// If the token exists in the price map
				if val.StrikePrice != 0 {
					// If it's an option (has a strike price)
					var delta, theta, gamma, vega, impliedVol float64

					if val.OptionType == "CE" {
						// For Call options
						syntheticFuture, _ := t.strikeToSyntheticFuture.Get(val.StrikePrice)
						if syntheticFuture == 0 {
							// If synthetic future price is not available, set all Greeks to 0
							delta, theta, gamma, vega, impliedVol = 0, 0, 0, 0, 0
						} else {
							// Calculate Greeks using Black-76 model
							K := float64(val.StrikePrice)                    // Strike price
							T := calculateTimeToExpiry(t.timeToExpireInDays) // Time to expiration (in years)
							r := 0.00                                        // Risk-free interest rate
							price := tick.LTP                                // Option price
							impliedVol = black76ImpliedVol(syntheticFuture, K, T, r, float64(price)/100)
							
							if impliedVol == 0 {
								// If implied volatility calculation fails, fetch Greeks from API
								greeksData, err := tiqs.GetGreeks_Tiqs(int(tick.Token), tiqs.ADMIN_TIQS)
								if err != nil {
									log.Printf("Error fetching Greeks from API: %v", err)
									delta, theta, gamma, vega = 0, 0, 0, 0
								} else {
									// Use Greeks from API
									delta = greeksData.Delta
									theta = greeksData.Theta
									gamma = greeksData.Gamma
									vega = greeksData.Vega
									impliedVol = greeksData.IV
								}
							} else {
								// Calculate Greeks using Black-76 model
								delta, theta, gamma, vega = black76Greeks(syntheticFuture, K, T, r, impliedVol)
							}
						}
					} else if val.OptionType == "PE" {
						// For Put options
						if ceToken, ok := t.peTokenToCeToken.Get(tick.Token); ok {
							if ceTickData, ok := t.priceMap.Get(ceToken); ok {
								// Derive Put option Greeks from corresponding Call option
								delta = ceTickData.Delta - 1
								theta = ceTickData.Theta
								gamma = ceTickData.Gamma
								vega = ceTickData.Vega
								impliedVol = ceTickData.IV
							}
						}
					}

					// Update the price map with new tick data and calculated Greeks
					go t.priceMap.Set(tick.Token, TickData{LTP: tick.LTP, Timestamp: tick.Time, StrikePrice: val.StrikePrice, OptionType: val.OptionType, Delta: delta, Theta: theta, Vega: vega, Gamma: gamma, IV: impliedVol})
				} else {
					// If it's not an option (e.g., underlying asset), update only LTP and timestamp
					go t.priceMap.Set(tick.Token, TickData{LTP: tick.LTP, Timestamp: tick.Time})
				}
			} else {
				// If the token is not in the map, add it with basic tick data
				t.priceMap.Set(tick.Token, TickData{LTP: tick.LTP, Timestamp: tick.Time})
			}
		}
	}()

	// Subscribe to the target symbol token : Index Token
	tiqsWs.AddSubscription(TargetSymbolToken)

	// Subscribe to the option chain tokens : Strike Price Tokens
	optionChain, err := tiqs.GetOptionChainMap_Tiqs(TargetSymbol, strconv.Itoa(TargetSymbolToken), "20")
	if err != nil {
		return fmt.Errorf("error while getting option chain: %w", err)
	}

	// Setting the optionChain in the TiqsGreeksClient
	t.optionChain = optionChain

	// Subscribe to the option chain tokens
	counter := 1
	for strike, tokens := range optionChain {

		// Convert strike price to int32 and set into strikeToSyntheticFuture
		t.strikeToSyntheticFuture.Set(typeConversion.StringToInt32(strike), 0)

		var ceToken, peToken int32
		for optionType, symbol := range tokens {
			tokenInt := typeConversion.StringToInt(symbol.Token)
			if tokenInt == 0 {
				continue
			}

			token := typeConversion.StringToInt32(symbol.Token)
			t.priceMap.Set(token, TickData{LTP: 0, Timestamp: 0, StrikePrice: int32(typeConversion.StringToInt(strike)), OptionType: optionType})

			if optionType == "CE" {
				ceToken = token
			} else if optionType == "PE" {
				peToken = token
			}

			tiqsWs.AddSubscription(tokenInt)
			t.logger(fmt.Sprintf("Subscribing to Symbol: %s, Token: %d, Total Symbols: %d", symbol.Name, tokenInt, counter))
			counter++
		}

		// Store the mapping of PE token to CE token
		if peToken != 0 && ceToken != 0 {
			t.peTokenToCeToken.Set(peToken, ceToken)
		}
	}

	t.logger(fmt.Sprintf("Total symbols subscribed: %d", counter-1))

	t.logger("WebSocket started successfully")

	// Setting synthetic future for each strike price in separate go routine
	go t.settingSyntheticFuture()

	return nil
}

func (t *TiqsGreeksClient) PrintPriceMap() {
	fmt.Println("PriceMap Contents:")
	t.priceMap.ForEach(func(key int32, value TickData) bool {
		fmt.Printf("Token: %d\n", key)
		fmt.Printf("  LTP: %d\n", value.LTP)
		fmt.Printf("  Timestamp: %d\n", value.Timestamp)
		fmt.Printf("  Strike Price: %d\n", value.StrikePrice)
		fmt.Printf("  Option Type: %s\n", value.OptionType)
		fmt.Println("--------------------")
		return true
	})
}

func (t *TiqsGreeksClient) settingSyntheticFuture() {

	// Setting synthetic future for each strike price
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			t.strikeToSyntheticFuture.ForEach(func(key int32, value float64) bool {
				strike := typeConversion.Int32ToString(key)
				if _, ok := t.optionChain[strike]["CE"]; !ok {
					return true
				}
				callToken := t.optionChain[strike]["CE"].Token
				putToken := t.optionChain[strike]["PE"].Token

				callPrice, err := t.GetPrice(typeConversion.StringToInt32(callToken))
				if err != nil {
					t.logger(fmt.Sprintf("Error in getting call price for strike: %s: %v", strike, err))
					return true
				}

				putPrice, err := t.GetPrice(typeConversion.StringToInt32(putToken))
				if err != nil {
					t.logger(fmt.Sprintf("Error in getting put price for strike: %s: %v", strike, err))
					return true
				}

				syntheticFuture := float64(key) + callPrice - putPrice
				t.strikeToSyntheticFuture.Set(key, syntheticFuture)
				return true
			})
		}
	}()
}

func (t *TiqsGreeksClient) PrintSyntheticFutureMap() {
	fmt.Println("Synthetic Future Map Contents:")
	t.strikeToSyntheticFuture.ForEach(func(key int32, value float64) bool {
		log.Printf("Strike Price: %d\n", key)
		log.Printf("  Synthetic Future: %f\n", value)
		log.Println("--------------------")
		return true
	})
}

func (t *TiqsGreeksClient) settingTimeToExpiry(TargetSymbol string) error {
	// First we will fetch closest expiry for that Index
	closestExpiry, err := tiqs.ClosestExpiryDate_Tiqs(TargetSymbol, tiqs.ADMIN_TIQS)
	if err != nil {
		return fmt.Errorf("error while getting closest expiry: %w", err)
	}

	// Get today's date in the same format
	today := time.Now().Format("2-Jan-2006")

	// Parse both dates
	expiryDate, err := time.Parse("2-Jan-2006", closestExpiry)
	if err != nil {
		return fmt.Errorf("error while parsing closestExpiry: %w", err)
	}

	todayDate, err := time.Parse("2-Jan-2006", today)
	if err != nil {
		return fmt.Errorf("error while parsing today's date: %w", err)
	}

	// Calculate the difference in days and add 1
	diffDays := expiryDate.Sub(todayDate).Hours() / 24
	t.timeToExpireInDays = int(diffDays) + 1

	return nil
}

// CalculateTimeToExpiry calculates the time to expiry in years based on the number of days
// so if today is expiry then give daysToExpiry is 1
func calculateTimeToExpiry(daysToExpiry int) float64 {
	// Get current time
	now := time.Now()

	// Set the expiry time to 15:30 on the expiry day
	expiry := time.Date(now.Year(), now.Month(), now.Day(), 15, 30, 0, 0, now.Location())

	// Add the number of days to expiry
	expiry = expiry.AddDate(0, 0, daysToExpiry-1)

	// Calculate the time difference
	timeDiff := expiry.Sub(now)

	// Convert the time difference to days
	daysFraction := timeDiff.Hours() / 24

	// Convert days to years
	years := daysFraction / 365

	return years
}
