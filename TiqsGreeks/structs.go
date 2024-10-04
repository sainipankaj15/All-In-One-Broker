package tiqs_greeks_socket

import (
	"github.com/alphadose/haxmap"
	tiqs "github.com/sainipankaj15/All-In-One-Broker/Tiqs"
)

// TiqsGreeksClient represents the main client structure for the Tiqs Greeks calculation
type TiqsGreeksClient struct {
	appID                   string                           // Application ID for authentication
	accessToken             string                           // Access token for API calls
	enableLog               bool                             // Flag to enable/disable logging
	timeToExpireInDays      int                              // Time to expiry in days for options
	priceMap                *haxmap.Map[int32, TickData]     // Map to store latest tick data for each instrument
	optionChain             map[string]map[string]tiqs.Symbol // Nested map to store option chain data
	strikeToSyntheticFuture *haxmap.Map[int32, float64]      // Map to store synthetic future prices for each strike
	peTokenToCeToken        *haxmap.Map[int32, int32]        // Map to link PE (Put) tokens to corresponding CE (Call) tokens
}

// TickData represents the tick data and calculated Greeks for an option
type TickData struct {
	LTP         int32   // Last Traded Price in paisa
	Timestamp   int32   // Unix timestamp of the last update
	StrikePrice int32   // Strike price of the option
	OptionType  string  // Type of option (CE for Call, PE for Put)
	Delta       float64 // Delta of the option
	Theta       float64 // Theta of the option
	Vega        float64 // Vega of the option
	Gamma       float64 // Gamma of the option
	IV          float64 // Implied Volatility of the option
}
