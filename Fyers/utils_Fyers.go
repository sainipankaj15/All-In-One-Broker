package fyers

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadingAccessToken_Fyers(userFyersID string) (string, error) {

	fileName := userFyersID + `.json`

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	var fileData ReadDataJson_Fyers

	err = json.Unmarshal(fileContent, &fileData)
	if err != nil {
		return "", err
	}
	accessToken := fileData.AccessTokenWithAppID

	return accessToken, nil
}

// GetOptionChainMap_Fyers retrieves the option chain map for a specified symbol and user ID from Fyers.
// It returns a nested map where the outer key is the strike price (as an integer),
// and the inner map contains option types (CE or PE) as keys and their corresponding symbol details.
func GetOptionChainMap_Fyers(SymbolName string, StrikeCount int, UserID_Fyers string) (map[int]map[string]Symbol, error) {

	// Fetch option chain from Fyers
	optionChainResp, err := GetOptionChain_Fyers(SymbolName, StrikeCount, UserID_Fyers)
	if err != nil {
		return nil, err
	}

	// Initialize a map to store the option chain data
	optionMap := make(map[int]map[string]Symbol)

	// Iterate through the OptionsChain from the response to populate the map
	for _, option := range optionChainResp.Data.OptionsChain {
		// Ensure the outer map (keyed by StrikePrice) exists
		if _, exists := optionMap[int(option.StrikePrice)]; !exists {
			optionMap[int(option.StrikePrice)] = make(map[string]Symbol)
		}

		// Create the Symbol struct with relevant details
		symbol := Symbol{
			Name:          option.Symbol,
			FyToken:       option.FyToken,
			TradingSymbol: getTradingSymbolFromName(option.Symbol),
		}

		// Populate the inner map with OptionType as key and Symbol struct as value
		optionMap[int(option.StrikePrice)][option.OptionType] = symbol
	}

	// Return the constructed map
	return optionMap, nil
}

// PrintOptionChainMap returns the nested option chain map as a formatted string.
func PrintOptionChainMap(optionMap map[int]map[string]Symbol) string {
	var builder strings.Builder
	builder.WriteString("Option Chain Data:\n")
	for strikePrice, innerMap := range optionMap {
		builder.WriteString("Strike Price: ")
		builder.WriteString(strconv.Itoa(strikePrice))
		builder.WriteString("\n")
		for optionType, symbol := range innerMap {
			builder.WriteString("  Option Type: ")
			builder.WriteString(optionType)
			builder.WriteString("\n")
			builder.WriteString("    Symbol: ")
			builder.WriteString(symbol.Name)
			builder.WriteString("\n")
			builder.WriteString("    FyToken: ")
			builder.WriteString(symbol.FyToken)
			builder.WriteString("\n")
			builder.WriteString("    Trading Symbol: ")
			builder.WriteString(symbol.TradingSymbol)
			builder.WriteString("\n")
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

// getTradingSymbolFromName returns a substring of `s` starting from the 4th character.
// This is used to extract the trading symbol
// If the length of `s` is less than 4, an empty string is returned.
func getTradingSymbolFromName(symbolName string) string {
	// Check if the length of `s` is greater than or equal to 4
	if len(symbolName) >= 4 {
		// Return the substring starting from the 4th character
		// Skip First 4 characters (NSE: and BSE:)
		return symbolName[4:]
	}
	// Return empty string if the length of `s` is less than 4
	return ""
}
