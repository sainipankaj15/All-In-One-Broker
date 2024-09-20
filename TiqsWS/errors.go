package tiqs_socket

import "errors"

var (
	ErrSocketConnectionClosed = errors.New("🔴 Socket connection closed")
	ErrSocketConnection       = errors.New("⛔ Error while connecting to socket, will try again for reconnect")
	ErrMarshlingMsg           = errors.New("⛔ Error while marshling message")
	ErrUnsupportedMsgType     = errors.New("⛔ Unsupported message type")
	ErrEmitingToSocket        = errors.New("⛔ Error emitting to socket")
	ErrSocketNotConnected     = errors.New("⛔ Socket is not connected")
	ErrClosingConnection      = errors.New("🔴 Error Closing WebSocket connection")
	ErrInvalidByteSliceLength = errors.New("⛔ Invalid byte slice length")
	ErrDecodingMessage        = errors.New("⛔ Error decoding message")
)
