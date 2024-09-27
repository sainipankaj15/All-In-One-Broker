package tiqs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	typeConversion "github.com/sainipankaj15/data-type-conversion"
)

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

func ClosestExpiryDate_Tiqs(indexName string, UserId_Tiqs string) (string, error) {

	resp, _, err := GetExpiryList_Tiqs(UserId_Tiqs)

	if err != nil {
		return "", err
	}

	allExpiryList := resp.Data[indexName]
	lastExpiryDate := allExpiryList[0]
	return lastExpiryDate, nil
}

func NextExpiryDateOnExpiry_Tiqs(indexName string, UserId_Tiqs string) (string, error) {

	resp, _, err := GetExpiryList_Tiqs(UserId_Tiqs)

	if err != nil {
		return "", err
	}

	allExpiryList := resp.Data[indexName]

	today := strings.ToUpper(time.Now().Format("2-Jan-2006"))
	if allExpiryList[0] == today {
		return allExpiryList[1], nil
	}

	return allExpiryList[0], nil
}

func GetMonthlyExpiry_Tiqs(indexName string, UserId_Tiqs string) (string, error) {
	resp, _, err := GetExpiryList_Tiqs(UserId_Tiqs)

	if err != nil {
		return "", err
	}

	allExpiryList := resp.Data[indexName]

	if len(allExpiryList) < 2 {
		return "", fmt.Errorf("not enough expiry dates for %s", indexName)
	}

	firstExpiryDate, err := time.Parse("2-Jan-2006", allExpiryList[0])
	if err != nil {
		return "", fmt.Errorf("error parsing first expiry date: %v", err)
	}
	currentMonth := firstExpiryDate.Month()

	for i, expiryDate := range allExpiryList {
		t, err := time.Parse("2-Jan-2006", expiryDate)
		if err != nil {
			return "", fmt.Errorf("error parsing expiry date: %v", err)
		}

		if t.Month() != currentMonth {
			// Return the previous expiry date (which is the last one in the current month)
			return allExpiryList[i-1], nil
		}
	}

	// If we've gone through all dates and haven't found a month change,
	// return the last date in the list
	return allExpiryList[len(allExpiryList)-1], nil
}

func GetOptionChainMap_Tiqs(TargetSymbol, TargetSymbolToken, OptionChainLength string) (map[string]map[string]Symbol, error) {
	// First we will fetch closest expiry for that Index
	closestExpiry, err := ClosestExpiryDate_Tiqs(TargetSymbol, ADMIN_TIQS)
	if err != nil {
		return nil, err
	}

	// Now we will fetch option chain from Tiqs using that closestExpiry
	optionChainFromTiqs, _, err := GetOptionChain_Tiqs(TargetSymbolToken, OptionChainLength, closestExpiry, ADMIN_TIQS)
	if err != nil {
		return nil, err
	}

	// Create a local variable to store the option chain
	optionChain := make(map[string]map[string]Symbol)

	for _, data := range optionChainFromTiqs.Data {
		optionChain[data.StrikePrice] = make(map[string]Symbol)
	}

	for _, data := range optionChainFromTiqs.Data {
		optionChain[data.StrikePrice][data.OptionType] = Symbol{
			Name:  data.Symbol,
			Token: data.Token,
		}
	}

	newOptionChain := make(map[string]map[string]Symbol)

	for strike, innerMap := range optionChain {
		strikeInFloat := typeConversion.StringToFloat64(strike)
		strikeInInt := typeConversion.Float64ToInt(strikeInFloat)
		strikeRounded := typeConversion.IntToString(strikeInInt)

		// If the rounded strike doesn't exist in the new map, initialize it
		if _, exists := newOptionChain[strikeRounded]; !exists {
			newOptionChain[strikeRounded] = make(map[string]Symbol)
		}

		// Copy CE and PE data to the rounded strike
		for optionType, symbol := range innerMap {
			newOptionChain[strikeRounded][optionType] = symbol
		}
	}

	// Return the processed option chain
	return newOptionChain, nil
}
