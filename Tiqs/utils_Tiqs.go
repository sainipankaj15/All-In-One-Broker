package tiqs

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	typeConversion "github.com/sainipankaj15/data-type-conversion"
)

// ReadingAccessToken_Tiqs reads the access token, session and APPID from a file.
//
// The file name is the userID_Tiqs + `.json`.
// The file must contain a JSON object with the following fields:
// - AccessToken: the access token.
// - APPID: the APPID.
//
// If the file does not exist, or if there is an error reading the file,
// or if the JSON object does not contain the required fields,
// the function returns an error.
func ReadingAccessToken_Tiqs(userID_Tiqs string) (string, string, error) {
	fileName := userID_Tiqs + `.json`

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", "", err
	}

	var fileData readDataJsonTiqs

	err = json.Unmarshal(fileContent, &fileData)
	if err != nil {
		return "", "", err
	}

	accessToken := fileData.AccessToken
	APPID := fileData.APPID

	return accessToken, APPID, nil
}

// CurrentQtyForAnySymbol_Tiqs returns the current quantity for the given symbol.
//
// It takes a symbol, product type and a UserID_Tiqs as parameters.
// It fetches the net position from the position API and iterates over all the positions
// to find the current quantity for the given symbol and product type.
// If the position is found, it returns the quantity in string format.
// If the position is not found, it returns an empty string.
// If there is an error, it returns an error.
func CurrentQtyForAnySymbol_Tiqs(symbolExchToken string, productType string, UserId_Tiqs string) (string, error) {

	positionAPIResp_Tiqs, err := PositionApi_Tiqs(UserId_Tiqs)

	if err != nil {
		return "", err
	}

	for i := 0; i < len(positionAPIResp_Tiqs.NetPosition_Tiqss); i++ {
		position := positionAPIResp_Tiqs.NetPosition_Tiqss[i]

		if position.Token == symbolExchToken {
			if position.Product == productType {
				return position.Qty, nil
			}
		}
	}

	return "", nil
}

// ExitAllPosition_Tiqs exits all open positions for a given user by placing market orders in the opposite direction.
// It takes the UserID_Tiqs as a parameter and returns a success message and an error if something goes wrong.
func ExitAllPosition_Tiqs(UserId_Tiqs string) (string, error) {

	// Fetch current positions using the Position API
	positionAPIResp_Tiqs, err := PositionApi_Tiqs(UserId_Tiqs)
	if err != nil {
		return "failed", err
	}

	// Iterate over all net positions
	for i := 0; i < len(positionAPIResp_Tiqs.NetPosition_Tiqss); i++ {
		position := positionAPIResp_Tiqs.NetPosition_Tiqss[i]

		// Use a goroutine to exit positions concurrently
		go func(pos NetPosition_Tiqs) {
			// Extract buy and sell quantities as strings
			buyQtyInString := position.DayBuyQty
			sellQtyInString := position.DaySellQty

			// Convert quantities to integers
			buyQty := typeConversion.StringToInt(buyQtyInString)
			sellQty := typeConversion.StringToInt(sellQtyInString)

			// Calculate the difference to determine net position
			diff := buyQty - sellQty

			if diff > 0 {
				// Long position: place a sell order to exit
				qtyInString := typeConversion.IntToString(diff)
				go OrderPlaceMarket_Tiqs(position.Exchange, position.Token, qtyInString, "S", position.Product, UserId_Tiqs)
			} else if diff < 0 {
				// Short position: place a buy order to exit
				qtyInString := typeConversion.IntToString(diff)
				go OrderPlaceMarket_Tiqs(position.Exchange, position.Token, qtyInString, "B", position.Product, UserId_Tiqs)
			}

		}(position)
	}

	return "success", nil
}

// ExitByPositionID_Tiqs exits a position by ID.
// It takes a symbol exch token, product type and a UserID_Tiqs as parameters.
// It fetches the current positions using the Position API and iterates over all the positions.
// If the position is found, it places a market order in the opposite direction to exit the position.
// If there is an error, it returns an error.
func ExitByPositionID_Tiqs(symbolExchToken string, productType string, UserId_Tiqs string) error {

	// Fetch current positions using the Position API
	positionAPIResp_Tiqs, err := PositionApi_Tiqs(UserId_Tiqs)

	if err != nil {
		return err
	}

	// Iterate over all net positions
	for i := 0; i < len(positionAPIResp_Tiqs.NetPosition_Tiqss); i++ {
		position := positionAPIResp_Tiqs.NetPosition_Tiqss[i]

		// Use a goroutine to exit positions concurrently
		go func(pos NetPosition_Tiqs) {

			// Check if the position is the one we want to exit
			if position.Token == symbolExchToken {
				if position.Product == productType {

					// Extract buy and sell quantities as strings
					buyQtyInString := position.DayBuyQty
					sellQtyInString := position.DaySellQty

					// Convert quantities to integers
					buyQty := typeConversion.StringToInt(buyQtyInString)
					sellQty := typeConversion.StringToInt(sellQtyInString)

					// Calculate the difference to determine net position
					diff := buyQty - sellQty

					if diff > 0 {
						// Long position: place a sell order to exit
						qtyInString := typeConversion.IntToString(diff)
						OrderPlaceMarket_Tiqs(position.Exchange, position.Token, qtyInString, "S", position.Product, UserId_Tiqs)
					} else if diff < 0 {
						// Short position: place a buy order to exit
						qtyInString := typeConversion.IntToString(diff)
						OrderPlaceMarket_Tiqs(position.Exchange, position.Token, qtyInString, "B", position.Product, UserId_Tiqs)
					}
				}
			}
		}(position)
	}
	return nil
}

// ClosestExpiryDate_Tiqs retrieves the closest expiry date for a given index.
// It fetches the list of expiry dates for the given user ID and returns the first
// expiry date in the list.
func ClosestExpiryDate_Tiqs(indexName string, UserId_Tiqs string) (string, error) {

	tiqsExpireDate := ""

	resp, err := GetExpiryList_Tiqs(UserId_Tiqs)

	// if err != nil {
	// 	return "", err
	// }
	if err == nil && len(resp.Data[indexName]) != 0 {
		tiqsExpireDate = resp.Data[indexName][0]
	}

	// allExpiryList := resp.Data[indexName]
	// lastExpiryDate := allExpiryList[0]

	dates, err := nseOptionChainFromNSE(indexName)
	if err != nil {
		fmt.Println("Error from NSE otpion chain")
		return tiqsExpireDate, nil
	}

	// if dates[0] == allExpiryList[0] {
	// 	return lastExpiryDate, nil
	// }
	return dates[0], nil
}

// NextExpiryDateOnExpiry_Tiqs retrieves the next expiry date for a given index,
// provided the current date is the expiry date of the first expiry in the list.
// If the current date is not the expiry date of the first expiry, it will return
// the first expiry date in the list.
func NextExpiryDateOnExpiry_Tiqs(indexName string, UserId_Tiqs string) (string, error) {
	resp, err := GetExpiryList_Tiqs(UserId_Tiqs)

	if err != nil {
		return "", err
	}

	allExpiryList := resp.Data[indexName]

	today := strings.ToUpper(time.Now().Format("2-Jan-2006"))

	// Check if today is the expiry date of the first expiry in the list
	if allExpiryList[0] == today {
		// Return the next expiry date
		return allExpiryList[1], nil
	}

	// Return the first expiry date in the list
	return allExpiryList[0], nil
}

// GetMonthlyExpiry_Tiqs retrieves the monthly expiry date for a given index.
// It checks the expiry dates and returns the last expiry date within the current month.
// If there's no month change, it returns the last expiry date in the list.
func GetMonthlyExpiry_Tiqs(indexName string, UserId_Tiqs string) (string, error) {
	// Fetch the list of expiry dates for the given user ID
	resp, err := GetExpiryList_Tiqs(UserId_Tiqs)
	if err != nil {
		return "", err
	}

	// Retrieve all expiry dates for the specified index
	allExpiryList := resp.Data[indexName]

	// Check if there are at least two expiry dates
	if len(allExpiryList) < 2 {
		return "", fmt.Errorf("not enough expiry dates for %s", indexName)
	}

	// Parse the first expiry date to determine the current month
	firstExpiryDate, err := time.Parse("2-Jan-2006", allExpiryList[0])
	if err != nil {
		return "", fmt.Errorf("error parsing first expiry date: %v", err)
	}
	currentMonth := firstExpiryDate.Month()

	// Iterate through the expiry dates
	for i, expiryDate := range allExpiryList {
		// Parse each expiry date
		t, err := time.Parse("2-Jan-2006", expiryDate)
		if err != nil {
			return "", fmt.Errorf("error parsing expiry date: %v", err)
		}

		// Check if the month has changed
		if t.Month() != currentMonth {
			// Return the previous expiry date (last one in the current month)
			return allExpiryList[i-1], nil
		}
	}

	// If no month change is found, return the last date in the list
	return allExpiryList[len(allExpiryList)-1], nil
}

// GetOptionChainMap_Tiqs retrieves the option chain map for a given target symbol and its token.
// It returns a nested map where the outer key is the strike price (rounded to the nearest integer)
// and the inner map contains option types (CE or PE) as keys and their corresponding symbols.
func GetOptionChainMap_Tiqs(TargetSymbol, TargetSymbolToken, OptionChainLength string) (map[string]map[string]Symbol, error) {
	// Fetch the closest expiry date for the given symbol
	closestExpiry, err := ClosestExpiryDate_Tiqs(TargetSymbol, ADMIN_TIQS)
	if err != nil {
		return nil, err
	}

	// Fetch the option chain data for the target symbol using the closest expiry date
	optionChainFromTiqs, _, err := GetOptionChain_Tiqs(TargetSymbolToken, OptionChainLength, closestExpiry, ADMIN_TIQS)
	if err != nil {
		return nil, err
	}

	// Initialize a map to store the option chain data
	optionChain := make(map[string]map[string]Symbol)

	// Populate the initial option chain map with strike prices
	for _, data := range optionChainFromTiqs.Data {
		optionChain[data.StrikePrice] = make(map[string]Symbol)
	}

	// Map each option type to its corresponding symbol for each strike price
	for _, data := range optionChainFromTiqs.Data {
		optionChain[data.StrikePrice][data.OptionType] = Symbol{
			Name:  data.Symbol,
			Token: data.Token,
		}
	}

	// Initialize a new map to store the processed option chain with rounded strike prices
	newOptionChain := make(map[string]map[string]Symbol)

	// Iterate over the option chain to round strike prices and populate the new map
	for strike, innerMap := range optionChain {
		strikeInFloat := typeConversion.StringToFloat64(strike)
		strikeInInt := typeConversion.Float64ToInt(strikeInFloat)
		strikeRounded := typeConversion.IntToString(strikeInInt)

		// Ensure the new map has an entry for the rounded strike price
		if _, exists := newOptionChain[strikeRounded]; !exists {
			newOptionChain[strikeRounded] = make(map[string]Symbol)
		}

		// Copy option type and symbol data to the new map with the rounded strike price
		for optionType, symbol := range innerMap {
			newOptionChain[strikeRounded][optionType] = symbol
		}
	}

	// Return the processed option chain map
	return newOptionChain, nil
}

// IsHoliday_Tiqs checks if the current day is a holiday or not.
// It uses the GetHolidays_Tiqs function to fetch the holidays and
// checks if the current date is present in the holidays map.
// It returns true if it's a holiday and false otherwise.
func IsHoliday_Tiqs(UserID_Tiqs string) (bool, error) {
	// Get today's date in the required format (DD-MM-YYYY)
	today := time.Now().Format("02-01-2006")

	// Call GetHolidays_Tiqs to fetch the holidays
	holidaysResp, err := GetHolidays_Tiqs(UserID_Tiqs)
	if err != nil {
		return false, err
	}

	// Check if today's date is in the holidays map
	if _, exists := holidaysResp.Data.Holidays[today]; exists {
		return true, nil // Today is a holiday
	}

	return false, nil // Today is not a holiday
}

func nseOptionChainFromNSE(symbol string) ([]string, error) {
	// Create a new HTTP client
	client := &http.Client{}

	endpoint := fmt.Sprintf("https://www.nseindia.com/api/option-chain-indices?symbol=%v", symbol)
	// Create a new request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("authority", "www.nseindia.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Create a variable of the response struct
	var optionChainResp nseOptionChainResp

	// Unmarshal the JSON response into the struct
	err = json.Unmarshal(body, &optionChainResp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	// Format all dates to have uppercase months
	formattedDates := make([]string, len(optionChainResp.Records.ExpiryDates))
	for i, date := range optionChainResp.Records.ExpiryDates {
		formattedDates[i] = formatDateString(date)
	}

	return formattedDates, nil
}

func formatDateString(date string) string {
	// Split the date string by "-"
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return date
	}

	// Capitalize the month part (parts[1])
	month := strings.ToUpper(parts[1])

	// Reconstruct the date string with uppercase month
	return fmt.Sprintf("%s-%s-%s", parts[0], month, parts[2])
}
