package fyers

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// ReadAccessToken_Fyers reads the Fyers access token for the given user.
func ReadAccessToken_Fyers(userID string) (string, error) {
	fileName := userID + ".json"

	fileContent, err := os.ReadFile(fileName)
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

// GetOptionChainMap retrieves the option chain map for a specified symbol and user ID from Fyers.
// It returns a nested map where the outer key is the strike price (as an integer),
// and the inner map contains option types (CE or PE) as keys and their corresponding symbol details.
func GetOptionChainMap(symbol string, strikeCount int, userID string) (map[int]map[string]Symbol, error) {
	optionChainResp, err := GetOptionChain(symbol, strikeCount, userID)
	if err != nil {
		return nil, err
	}

	optionMap := make(map[int]map[string]Symbol)

	for _, option := range optionChainResp.Data.OptionsChain {
		if _, exists := optionMap[int(option.StrikePrice)]; !exists {
			optionMap[int(option.StrikePrice)] = make(map[string]Symbol)
		}

		optionMap[int(option.StrikePrice)][option.OptionType] = Symbol{
			Name:          option.Symbol,
			FyToken:       option.FyToken,
			TradingSymbol: getTradingSymbolFromName(option.Symbol),
		}
	}

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

// getTradingSymbolFromName returns the trading symbol by removing the exchange prefix
// (for example, "NSE:" or "BSE:") from the full Fyers symbol name.
func getTradingSymbolFromName(symbolName string) string {
	if len(symbolName) >= 4 {
		return symbolName[4:]
	}

	return ""
}
