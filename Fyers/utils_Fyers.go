package fyers

import (
	"encoding/json"
	"fmt"
	"math"
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

// GetATMOptionSymbols returns the ATM Call (CE) and Put (PE) option symbols
// for the given underlying symbol and user ID.
func GetATMOptionSymbols(symbol, userID string) (string, string, error) {
	optionChainResp, err := GetOptionChain(symbol, 1, userID)
	if err != nil {
		return "", "", err
	}

	optionsChain := optionChainResp.Data.OptionsChain

	var underlyingLTP float64

	// Get the underlying LTP from the option chain response.
	for _, option := range optionsChain {
		if option.OptionType == "" && option.StrikePrice == -1 {
			underlyingLTP = option.Ltp
			break
		}
	}

	if underlyingLTP == 0 {
		return "", "", fmt.Errorf("underlying LTP not found for %s", symbol)
	}

	// Find the strike closest to the underlying LTP.
	var atmStrike int
	minDifference := math.MaxFloat64

	for _, option := range optionsChain {
		if option.OptionType == "" || option.StrikePrice < 0 {
			continue
		}

		difference := math.Abs(underlyingLTP - float64(option.StrikePrice))

		if difference < minDifference {
			minDifference = difference
			atmStrike = int(option.StrikePrice)
		}
	}

	var callSymbol string
	var putSymbol string

	// Get CE and PE symbols for the ATM strike.
	for _, option := range optionsChain {
		if int(option.StrikePrice) != atmStrike {
			continue
		}

		switch option.OptionType {
		case OptionTypes.CALL:
			callSymbol = option.Symbol
		case OptionTypes.PUT:
			putSymbol = option.Symbol
		}
	}

	if callSymbol == "" || putSymbol == "" {
		return "", "", fmt.Errorf("ATM CE/PE symbols not found for %s", symbol)
	}

	return callSymbol, putSymbol, nil
}
