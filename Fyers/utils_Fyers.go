package fyers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// First string will be Strike after that for that strike value will again a map , in internal map key is string that is option type CE/PE and for both respective there is fytoken and symbolName and TradingSymbol
func GetOptionChainMap_Fyers(SymbolName string, StrikeCount int, UserID_Fyers string) (map[int]map[string]Symbol, error) {

	// First we will fetch option chain from Fyers
	optionChainResp, err := GetOptionChain_Fyers(SymbolName, StrikeCount, UserID_Fyers)
	if err != nil {
		return nil, err
	}

	// Then we will convert it into map
	// Create the nested map structure
	optionMap := make(map[int]map[string]Symbol)

	// Iterate through the OptionsChain from the response
	for _, option := range optionChainResp.Data.OptionsChain {
		// Ensure the outer map (keyed by StrikePrice) exists
		if _, exists := optionMap[int(option.StrikePrice)]; !exists {
			optionMap[int(option.StrikePrice)] = make(map[string]Symbol)
		}

		// Create the Symbol struct
		symbol := Symbol{
			Name:          option.Symbol,
			FyToken:       option.FyToken,
			TradingSymbol: getTradingSymbolFromName(option.Symbol),
		}

		// Populate the inner map with OptionType as key and Symbol struct as value
		optionMap[int(option.StrikePrice)][option.OptionType] = symbol
	}

	return optionMap, nil

}

// Function to print the nested map in a systematic and readable manner
func PrintOptionChainMap(optionMap map[int]map[string]Symbol) {
	fmt.Println("Option Chain Data:")
	for strikePrice, innerMap := range optionMap {
		fmt.Printf("Strike Price: %d\n", strikePrice)
		for optionType, symbol := range innerMap {
			fmt.Printf("  Option Type: %s\n", optionType)
			fmt.Printf("    Symbol: %s\n", symbol.Name)
			fmt.Printf("    FyToken: %s\n", symbol.FyToken)
		}
		fmt.Println() // Add a line break for readability between strike prices
	}
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
