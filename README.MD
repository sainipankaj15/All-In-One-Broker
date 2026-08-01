# All-In-One-Broker

All-In-One-Broker is a Go library for interacting with multiple Indian brokers through a consistent set of helpers. The current codebase includes integrations for Jainam, Fyers, Tiqs, Zerodha, and XTS, along with shared utilities for notifications, time handling, and market-data workflows.

## What this repository provides

The library focuses on practical trading operations such as:

- Authentication helpers that read broker credentials from local JSON files
- Order placement for supported brokers
- Position, holdings, funds, and order-book retrieval
- Market data access such as quotes, LTP, option chains, and market depth
- Option Greeks and live WebSocket-based market updates for Tiqs
- Shared utility helpers for Telegram/Slack notifications, holiday checks, and date/time handling

## Supported brokers

- Jainam
  - Order placement helpers
  - Access-token reading from local JSON files
- Fyers
  - Positions, holdings, funds, quotes, LTP, option chain, market depth, historical data, and order placement
  - Option-chain map helpers and symbol utilities
- Tiqs
  - Positions, quotes, LTP, option chain, expiry lists, Greeks, holidays, order status, order book, and order placement
  - WebSocket data streaming and Greeks calculation helpers
- Zerodha
  - Market orders, funds, holdings, and positions
- XTS
  - Market order placement and access-token helpers

## Repository layout

- /Jainam: Jainam broker helpers
- /Fyers: Fyers broker helpers and constants
- /Tiqs: Tiqs broker helpers, constants, and utility functions
- /TiqsWS: Tiqs WebSocket client
- /TiqsGreeks: Greeks calculation and live option analytics
- /Zerodha: Zerodha helpers
- /XTS: XTS helpers
- /commanUtilsAcrossBroker: Shared utilities such as Telegram/Slack notifications and time helpers


## Installation

```bash
go get github.com/sainipankaj15/All-In-One-Broker
```

## Authentication

Most helpers expect the broker credentials or access token to be available in a local JSON file named after the user ID. The exact file format depends on the broker helper, but the common pattern is that the helper reads a file such as:

```text
<userID>.json
```

Example usage:

```go
import (
    "fmt"

    fyers "github.com/sainipankaj15/All-In-One-Broker/Fyers"
    zerodha "github.com/sainipankaj15/All-In-One-Broker/Zerodha"
)

func main() {
    // Fyers helpers
    positions, err := fyers.GetPositions("XP03754")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(positions)

    // Zerodha helpers
    funds, err := zerodha.GetFunds("FC8173")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(funds)
}
```

## Order placement examples

### Fyers

```go
resp, err := fyers.PlaceLimitOrder(
    "NSE:SMALLCAP-EQ",
    45,
    1,
    fyers.TransactionSide.SELL,
    fyers.ProductType.INTRADAY,
    "XP03754",
)
```

### Zerodha

```go
resp, err := zerodha.PlaceMarketOrder(
    zerodha.Exchange.NSE,
    "SBIN",
    "1",
    zerodha.OrderType.MARKET,
    zerodha.TransactionSide.SELL,
    zerodha.ProductType.INTRADAY,
    "FC8173",
)
```

## Market data and analytics

### Fyers

```go
optionChain, err := fyers.GetOptionChain_Fyers("NIFTY", 10, "XP03754")
marketDepth, err := fyers.MarketDepthAPI_Fyers("NSE:NIFTY50-INDEX", "XP03754")
```

### Option-chain mapping helpers

```go
optionMap, err := fyers.GetOptionChainMap_Fyers(fyers.Index.NIFTY, 15, "XP03754")
```

## Position and portfolio helpers

```go
positions, err := fyers.GetPositions("XP03754")
holdings, err := zerodha.GetHoldings("FC8173")
```

## Utility helpers

```go
utils.TelegramSend(botToken, chatID, message)
utils.SlackSend(webhookURL, message)
fmt.Println(utils.CurrentDate())
utils.ApplicationStart(9, 15, 0)
```

## Dependencies

The main module depends on:

- github.com/alphadose/haxmap
- github.com/gorilla/websocket
- github.com/sainipankaj15/data-type-conversion

## Notes

- This repository is a library wrapper around broker APIs; it is not a standalone trading server.
- Token handling is based on local JSON files and the calling code is responsible for managing credentials.
- Some helpers are broker-specific and may require additional context such as app IDs, user IDs, or symbol/token details.

## Contributing

Contributions are welcome. Please open a pull request with a clear description of the change.

## License

See the LICENSE file for details.

## Disclaimer

This library is for educational and informational purposes only. Trading in financial markets involves risk. Always consult with a licensed financial advisor before making investment decisions.
