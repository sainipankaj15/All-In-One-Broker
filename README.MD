# All-In-One-Broker

A comprehensive Go library for interacting with multiple Indian stock brokers' APIs, including Jainam, Fyers, and Tiqs.

## Overview

This library provides a unified interface to interact with various Indian stock brokers' trading APIs. It handles authentication, order placement, position management, market data retrieval, and more across different brokers.

## Supported Brokers

- **Jainam**: Full support for Jainam Broking APIs
- **Fyers**: Complete integration with Fyers trading platform
- **Tiqs**: Comprehensive support for Tiqs trading APIs

## Features

- **Authentication**: Secure token management for each broker
- **Order Management**: 
  - Place market and limit orders
  - Modify and cancel orders
  - Track order status
- **Position Management**: 
  - View current positions
  - Exit positions individually or in bulk
- **Market Data**: 
  - Real-time quotes and LTP (Last Traded Price)
  - Option chains with Greeks calculations
  - Market depth information
- **Historical Data**: Access to historical price data
- **WebSocket Integration**: Real-time data streaming
- **Option Analytics**: 
  - Greeks calculation (Delta, Gamma, Theta, Vega)
  - Synthetic future pricing
- **Utility Functions**: 
  - Holiday calendar integration
  - Notification services (Telegram, Slack)
  - Time management utilities

## Package Structure

- `/Jainam`: Jainam broker API integration
- `/Fyers`: Fyers broker API integration
- `/Tiqs`: Tiqs broker API integration
- `/TiqsWS`: WebSocket implementation for Tiqs
- `/TiqsGreeks`: Option Greeks calculations for Tiqs
- `/commanUtilsAcrossBroker`: Shared utilities across brokers

## Installation

```bash
go get github.com/sainipankaj15/All-In-One-Broker
```

## Usage Examples

### Authentication

```go
// Jainam authentication
date, accessToken, err := jainam.ReadingAccessToken_Jainam("userID")

// Fyers authentication
accessToken, err := fyers.ReadingAccessToken_Fyers("userFyersID")

// Tiqs authentication
accessToken, appID, err := tiqs.ReadingAccessToken_Tiqs("userID_Tiqs")
```

### Market Data

```go
// Get option chain from Fyers
optionChainMap, err := fyers.GetOptionChain_Fyers("NIFTY", 10, "userID")

// Get LTP from Tiqs
ltp, err := tiqs.LTPInPaisa_Tiqs(tokenNumber, "userID_Tiqs")

// Get Greeks from Tiqs
greeksData, err := tiqs.GetGreeks_Tiqs(tokenNumber, "userID_Tiqs")

// Get market depth from Fyers
marketDepth, err := fyers.MarketDepthAPI_Fyers(symbolName, "userID_Fyers")
```

### Order Placement

```go
// Place market order with Jainam
response, err := jainam.OrderPlaceMarket_Jainam(exchange, token, quantity, priceType, orderType, transSide, productType, "userID_Jainam")

// Place limit order with Fyers
success, err := fyers.PlaceOrder_Fyers(symbolName, limitPrice, qty, fyers.OrderSide.BUY, fyers.ProductType.INTRADAY, "userID")

// Place market order with Tiqs
response, err := tiqs.OrderPlaceMarket_Tiqs(exchange, token, quantity, tiqs.OrderSide.BUY, string(tiqs.Product.INTRADAY), "userID_Tiqs")
```

### Position Management

```go
// Get positions from Fyers
positions, err := fyers.PositionApi_Fyers("userID_Fyers")

// Get positions from Tiqs
positions, err := tiqs.PositionApi_Tiqs("userID_Tiqs")

// Exit all positions in Fyers
err := fyers.ExitingAllPosition(sides, segments, productTypes, "userID_Fyers")

// Exit specific position in Fyers
err := fyers.ExitPositionByID_Fyers("userID_Fyers", symbolName)
```

### WebSocket Integration

```go
// Initialize Tiqs WebSocket client
client, err := tiqs_socket.NewTiqsWebSocket(appID, accessToken, true)

// Get data channel for real-time ticks
dataChannel := client.GetDataChannel()

// Get order updates channel
orderChannel := client.GetOrderChannel()

// Subscribe to a token
client.AddSubscription(tokenNumber)
```

### Option Greeks Calculation

```go
// Initialize Tiqs Greeks calculation
client, err := tiqs_greeks_socket.NewTiqsGreeksSocket(appID, accessToken, true)

// Start WebSocket for a specific symbol
err := client.StartWebSocket("NIFTY", niftyToken)

// Get delta for an option
delta, err := client.GetDelta(optionToken)

// Find nearest option with specific delta
token, err := client.GetNearestCallToken(0.5)
```

### Utility Functions

```go
// Send Telegram notification
utils.TelegramSend(botToken, chatID, message)

// Send Slack notification
utils.SlackSend(webhookURL, message)

// Get current date in IST
date := utils.CurrentDate()

// Schedule application start/stop
utils.ApplicationStart(9, 15, 0)  // Start at 9:15 AM
```

## Dependencies

- github.com/alphadose/haxmap: High-performance thread-safe concurrent map
- github.com/gorilla/websocket: WebSocket implementation for Go
- github.com/sainipankaj15/data-type-conversion: Data type conversion utilities

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[For License information, Please Check Licencse file]

## Disclaimer

This library is for educational and informational purposes only. Trading in financial markets involves risk. Always consult with a licensed financial advisor before making investment decisions.
