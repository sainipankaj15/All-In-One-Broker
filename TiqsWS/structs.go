package tiqs_socket

import (
	"time"

	"github.com/gorilla/websocket"
)

// TiqsWSClient represents the tiqs Websocket client
type TiqsWSClient struct {
	appID       string
	accessToken string
	socket      *websocket.Conn
	// pingCheckerTimer    *time.Timer
	lastPingTS   time.Time
	pendingQueue []interface{}
	// retryCount          int
	// reconnectTimer      *time.Timer
	// mu                  sync.Mutex
	wsURL                string
	enableLog            bool
	stopReadMessagesSig  chan bool
	stopPingListenerSig  chan bool
	isReconnectRequested bool

	subscriptions map[int]struct{} // All active subscriptions
	dataChannel   chan Tick        // data channel where data will come
	orderChannel  chan OrderUpdate // data channel where order update will come

}

// Tick represents the structure of a tick
type Tick struct {
	// Token
	Token int32
	// Last traded price
	LTP int32

	/*
		// Net change indicator
		NetChangeIndicator int32
		// Net change
		NetChange int32
		// Last traded quantity
		LTQ int32
		// Average traded price
		AvgPrice int32
		// Total buy quantity
		TotalBuyQuantity int32
		// Total sell quantity
		TotalSellQuantity int32
		// Open price
		Open int32
		// High price
		High int32
		// Close price
		Close int32
		// Low price
		Low int32
		// Volume
		Volume int32
		// Last traded time
		LTT int32

		// Open interest
		OI int32
		// Open interest day high
		OIDayHigh int32
		// Open interest day low
		OIDayLow int32
		// Lower limit
		LowerLimit int32
		// Upper limit
		UpperLimit int32
	*/
	// Time
	Time int32
}

// SocketMessage represents the structure of a socket message : which we are going to send to the websocket
type SocketMessage struct {
	Code string `json:"code"`
	Mode string `json:"mode"`
	Ltp  []int  `json:"ltp"`
}

// Define the structure to match the incoming JSON message
type OrderUpdate struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	UserID          string    `json:"userId"`
	Exchange        string    `json:"exchange"`
	Symbol          string    `json:"symbol"`
	Token           int       `json:"token"`
	Qty             int       `json:"qty"`
	Price           float64   `json:"price"`
	Product         string    `json:"product"`
	Status          string    `json:"status"`
	ReportType      string    `json:"reportType"`
	TransactionType string    `json:"transactionType"`
	Order           string    `json:"order"`
	Retention       string    `json:"retention"`
	AvgPrice        float64   `json:"avgPrice"`
	Reason          string    `json:"reason"`
	ExchangeOrderId string    `json:"exchangeOrderId"`
	CancelQty       string    `json:"cancelQty"`
	Tags            string    `json:"tags"`
	DisclosedQty    string    `json:"disclosedQty"`
	TriggerPrice    string    `json:"triggerPrice"`
	ExchangeTime    time.Time `json:"exchangeTime"`
	Timestamp       time.Time `json:"timestamp"`
}
