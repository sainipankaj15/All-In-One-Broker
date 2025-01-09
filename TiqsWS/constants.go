package tiqs_socket

// Constants
const (
	CODE_SUB           = "sub"
	CODE_UNSUB         = "unsub"
	MODE_FULL          = "full"
	MODE_LTP           = "ltp"
	maxRetries         = 20
	BUFFER_SIZE        = 100000
	FULLTICK_LENGTH    = 197
	ONLYLTPTICK_LENGTH = 13
)

// EndPoints
const SOCKET_URL = "wss://wss.tiqs.trading"

// Info Messages
const (
	INFO_CONNECTED_WEBSOCKET               = "🟢 Connected to socket"
	INFO_RECONNECT_REQUEST_IGNORED         = "🙈 Reconnect request ignored. Already requested."
	INFO_RECONNECT_LIMIT_REACHED           = "✋ Socket reconnection limit reached."
	INFO_SOCKET_RECONNECTING               = "🔄 Attempting socket reconnect"
	INFO_SOCKET_PING_DIFFERENCE            = "🆚 Socket ping difference exceeded: Reconnecting..."
	INFO_PROCCESSING_PENDING_REQUESTS      = "⏳ Processing pending requests..."
	INFO_PROCCESSING_PREVIOUS_SUBSCRIPTION = "⏳ Processing previous subscriptions if any..."
	INFO_CLOSED_WEBSOCKET                  = "🔴 WebSocket connection closed"
	INFO_INVALID_TICK_DATA                 = "Invalid tick data length "
	InfoSocketConnecting                   = "⏳ Connecting to socket..."
)

type OrderStatus string

const (
	COMPLETE OrderStatus = "COMPLETE"
	REJECTED OrderStatus = "REJECTED"
	PENDING  OrderStatus = "PENDING"
	OPEN     OrderStatus = "OPEN"
	CANCELED OrderStatus = "CANCELED"
)
