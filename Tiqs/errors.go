package tiqs

import "errors"

var (
	// ErrOrderIDExists          = errors.New("order ID already exists")
	// ErrOnTick                 = errors.New("error while executing onTick()")
	ErrOrderPlacementFailed = errors.New("order placement failed")
	// ErrBasketMarginFailed     = errors.New("basket margin failed")
	// ErrMarginFailed           = errors.New("single instrument margin failed")
	// ErrOptionChainFailed      = errors.New("option chain fetching failed")
	// ErrGetOrderStatusFailed   = errors.New("get order status failed")
	// ErrOrderBookFailed        = errors.New("order book fetching failed")
	// ErrTradeBookFailed        = errors.New("trade book fetching failed")
	// ErrPositionBookFailed     = errors.New("position book fetching failed")
	// ErrGettingLTP             = errors.New("getting LTP failed")
	// ErrPositionNotFound       = errors.New("position not found")
	// ErrGettingExpiryDates     = errors.New("error getting expiry dates")
	// ErrSocketConnectionClosed = errors.New("🔴 Socket connection closed")
	// ErrSocketConnection       = errors.New("⛔ Error while connecting to socket, will try again for reconnect in 3 seconds")
	// ErrMarshlingMsg           = errors.New("⛔ Error while marshling message")
	// ErrUnsupportedMsgType     = errors.New("⛔ Unsupported message type")
	// ErrEmitingToSocket        = errors.New("⛔ Error emitting to socket")
	// ErrSocketNotConnected     = errors.New("⛔ Socket is not connected")
	// ErrClosingConnection      = errors.New("🔴 Error Closing WebSocket connection")
	// ErrInvalidByteSliceLength = errors.New("⛔ Invalid byte slice length")
	// ErrDecodingMessage        = errors.New("⛔ Error decoding message")
	// ErrReadingSocketMessage   = errors.New("😔 Error reading socket message")
	// ErrChannelBlocked         = errors.New("🚫 channel is blocked, dropping message")
)
