package tiqs_greeks_socket

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/alphadose/haxmap"
	tiqs "github.com/sainipankaj15/All-In-One-Broker/Tiqs"
	tiqsSocket "github.com/sainipankaj15/All-In-One-Broker/TiqsWS"
	typeConversion "github.com/sainipankaj15/data-type-conversion"
)

// NewTiqsGreeksSocket initializes and returns a new TiqsGreeksClient
func NewTiqsGreeksSocket(appID string, accessToken string, enableLog bool) (*TiqsGreeksClient, error) {
	client := &TiqsGreeksClient{
		appID:                   appID,
		accessToken:             accessToken,
		enableLog:               enableLog,
		priceMap:                haxmap.New[int32, TickData](),           // token to tick data
		strikeToSyntheticFuture: haxmap.New[int32, float64](),            // strike price to synthetic future price
		peTokenToCeToken:        haxmap.New[int32, int32](),              // PE token to corresponding CE token
		optionChain:             make(map[string]map[string]tiqs.Symbol), // option chain data
	}

	if enableLog {
		log.Printf("Initializing NewTiqsGreeksSocket with appID: %s and accessToken: %s", appID, accessToken)
		client.logger("TiqsGreeksClient created successfully")
	}

	return client, nil
}

// logger is a helper method to log messages if logging is enabled
func (t *TiqsGreeksClient) logger(msg ...any) {
	if t.enableLog {
		log.Println(msg...)
	}
}

// GetTickData retrieves the full TickData for a given token
func (t *TiqsGreeksClient) GetTickData(token int32) (TickData, error) {
	// Check if the data is already in the cache
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

	// Check if we have recent data in the cache
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

// GetOptionChainMap returns the internal map of option chain data.
// This map contains the mapping of strike price to its corresponding CE and PE symbols.
func (t *TiqsGreeksClient) GetOptionChainMap() map[string]map[string]tiqs.Symbol {
	return t.optionChain
}

// StartWebSocket initializes and starts the WebSocket connection
func (t *TiqsGreeksClient) StartWebSocket(TargetSymbol string, TargetSymbolToken int) error {
	// Setting the time to expiry in Days
	err := t.settingTimeToExpiry(TargetSymbol)
	if err != nil {
		return fmt.Errorf("error while setting time to expiry: %w", err)
	}

	// Initialize the WebSocket connection
	tiqsWs, err := tiqsSocket.NewTiqsWebSocket(t.appID, t.accessToken, t.enableLog)
	if err != nil {
		return fmt.Errorf("error while connecting tiqs socket: %w", err)
	}

	dataChannel := tiqsWs.GetDataChannel()

	// Start a goroutine to handle incoming WebSocket data
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
	optionChain, err := tiqs.GetOptionChainMap_Tiqs(TargetSymbol, strconv.Itoa(TargetSymbolToken), "18")
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

// PrintPriceMap prints the contents of the price map for debugging purposes
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

// settingSyntheticFuture calculates and updates the synthetic future prices
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

// PrintSyntheticFutureMap prints the contents of the synthetic future map for debugging purposes
func (t *TiqsGreeksClient) PrintSyntheticFutureMap() {
	fmt.Println("Synthetic Future Map Contents:")
	t.strikeToSyntheticFuture.ForEach(func(key int32, value float64) bool {
		log.Printf("Strike Price: %d\n", key)
		log.Printf("  Synthetic Future: %f\n", value)
		log.Println("--------------------")
		return true
	})
}

// settingTimeToExpiry calculates and sets the time to expiry for the options
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

// calculateTimeToExpiry calculates the time to expiry in years based on the number of days
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

// GetNearestCallToken returns the token number of the Call option with the nearest delta value
func (t *TiqsGreeksClient) GetNearestCallToken(delta float64) (int32, error) {
	nearestToken := int32(0)
	nearestDeltaDiff := float64(1e9) // Start with a large number

	t.priceMap.ForEach(func(key int32, value TickData) bool {
		if value.OptionType == "CE" { // Check if it's a Call option
			deltaDiff := math.Abs(float64(value.Delta) - delta)
			if deltaDiff < nearestDeltaDiff {
				nearestDeltaDiff = deltaDiff
				nearestToken = key
			}
		}
		return true
	})

	if nearestToken == 0 {
		return 0, fmt.Errorf("no Call option found with a delta close to %f", delta)
	}
	return nearestToken, nil
}

// GetNearestPutToken returns the token number of the Put option with the nearest delta value
func (t *TiqsGreeksClient) GetNearestPutToken(delta float64) (int32, error) {
	nearestToken := int32(0)
	nearestDeltaDiff := float64(1e9) // Start with a large number

	t.priceMap.ForEach(func(key int32, value TickData) bool {
		if value.OptionType == "PE" { // Check if it's a Put option
			deltaDiff := math.Abs(float64(value.Delta) + delta) // Make delta positive for comparison
			if deltaDiff < nearestDeltaDiff {
				nearestDeltaDiff = deltaDiff
				nearestToken = key
			}
		}
		return true
	})

	if nearestToken == 0 {
		return 0, fmt.Errorf("no Put option found with a delta close to %f", delta)
	}
	return nearestToken, nil
}

// GetDeltaDifference calculates the absolute difference between the deltas of two tokens
func (t *TiqsGreeksClient) GetDeltaDifference(token1, token2 int32) (float64, error) {
	// Retrieve TickData for both tokens
	tickData1, ok1 := t.priceMap.Get(token1)
	tickData2, ok2 := t.priceMap.Get(token2)

	if !ok1 {
		return 0, fmt.Errorf("token %d not found in price map", token1)
	}
	if !ok2 {
		return 0, fmt.Errorf("token %d not found in price map", token2)
	}

	// Get absolute delta values
	delta1 := math.Abs(float64(tickData1.Delta))
	delta2 := math.Abs(float64(tickData2.Delta))

	// Calculate the difference
	deltaDifference := math.Abs(delta1 - delta2)

	return deltaDifference, nil
}

// GetDelta returns the absolute value of the delta for a given token
func (t *TiqsGreeksClient) GetDelta(token int32) (float64, error) {
	// Retrieve TickData for the token
	tickData, ok := t.priceMap.Get(token)
	if !ok {
		return 0, fmt.Errorf("token %d not found in price map", token)
	}

	// Return the absolute value of the delta
	return math.Abs(float64(tickData.Delta)), nil
}
