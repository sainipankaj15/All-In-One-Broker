package tiqs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	var fileData ReadDataJsonTiqs

	err = json.Unmarshal(fileContent, &fileData)
	if err != nil {
		return "", "", err
	}

	accessToken := fileData.AccessToken
	APPID := fileData.APPID

	return accessToken, APPID, nil
}

func CurrentQtyForAnySymbol_Tiqs(symbolExchToken string, productType string, UserId_Tiqs string) (string, error) {

	PositionAPIResp_Tiqs, err := PositionApi_Tiqs(UserId_Tiqs)

	if err != nil {
		return "", err
	}

	for i := 0; i < len(PositionAPIResp_Tiqs.NetPosition_Tiqss); i++ {
		position := PositionAPIResp_Tiqs.NetPosition_Tiqss[i]

		if position.Token == symbolExchToken {
			if position.Product == productType {
				return position.Qty, nil
			}
		}
	}

	return "", nil
}

func ExitAllPosition_Tiqs(UserId_Tiqs string) (string, error) {

	PositionAPIResp_Tiqs, err := PositionApi_Tiqs(UserId_Tiqs)

	if err != nil {
		return "failed", err
	}

	for i := 0; i < len(PositionAPIResp_Tiqs.NetPosition_Tiqss); i++ {
		position := PositionAPIResp_Tiqs.NetPosition_Tiqss[i]

		go func(pos NetPosition_Tiqs) {
			buyQtyInString := position.DayBuyQty
			sellQtyInString := position.DaySellQty

			buyQty := typeConversion.StringToInt(buyQtyInString)
			sellQty := typeConversion.StringToInt(sellQtyInString)

			diff := buyQty - sellQty

			if diff > 0 {
				// it means long position : Have to cut it by opposite order
				qtyInString := typeConversion.IntToString(diff)
				go OrderPlaceMarket_Tiqs(position.Exchange, position.Token, qtyInString, "S", position.Product, UserId_Tiqs)
			} else if diff < 0 {
				// it means short position : Have to cut it by opposite order
				qtyInString := typeConversion.IntToString(diff)
				go OrderPlaceMarket_Tiqs(position.Exchange, position.Token, qtyInString, "B", position.Product, UserId_Tiqs)
			}

		}(position)
	}

	return "success", nil
}

func ExitByPositionID_Tiqs(symbolExchToken string, productType string, UserId_Tiqs string) error {

	PositionAPIResp_Tiqs, err := PositionApi_Tiqs(UserId_Tiqs)

	if err != nil {
		return err
	}

	for i := 0; i < len(PositionAPIResp_Tiqs.NetPosition_Tiqss); i++ {
		position := PositionAPIResp_Tiqs.NetPosition_Tiqss[i]

		go func(pos NetPosition_Tiqs) {

			if position.Token == symbolExchToken {
				if position.Product == productType {

					buyQtyInString := position.DayBuyQty
					sellQtyInString := position.DaySellQty

					buyQty := typeConversion.StringToInt(buyQtyInString)
					sellQty := typeConversion.StringToInt(sellQtyInString)

					diff := buyQty - sellQty

					if diff > 0 {
						// it means long position : Have to cut it by opposite order
						qtyInString := typeConversion.IntToString(diff)
						OrderPlaceMarket_Tiqs(position.Exchange, position.Token, qtyInString, "S", position.Product, UserId_Tiqs)
					} else if diff < 0 {
						// it means short position : Have to cut it by opposite order
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

	resp, err := GetExpiryList_Tiqs(UserId_Tiqs)

	if err != nil {
		return "", err
	}

	allExpiryList := resp.Data[indexName]
	lastExpiryDate := allExpiryList[0]
	return lastExpiryDate, nil
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
