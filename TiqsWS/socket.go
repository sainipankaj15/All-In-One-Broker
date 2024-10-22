package tiqs_socket

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// NewTiqsWebSocket sets up the WebSocket connection and related processes
// This function should be called from your main function
func NewTiqsWebSocket(appID string, accessToken string, enableLog bool) (*TiqsWSClient, error) {
	tiqsWSClient := TiqsWSClient{
		appID:               appID,
		accessToken:         accessToken,
		subscriptions:       make(map[int]struct{}),
		dataChannel:         make(chan Tick, BUFFER_SIZE),
		orderChannel:        make(chan OrderUpdate, BUFFER_SIZE),
		stopReadMessagesSig: make(chan bool),
		stopPingListenerSig: make(chan bool),
		enableLog:           enableLog,
		wsURL:               fmt.Sprintf("%s?appId=%s&token=%s", SOCKET_URL, appID, accessToken),
	}

	// wsURL := fmt.Sprintf("%s?appId=%s&token=%s", SOCKET_URL, appID, accessToken)

	err := tiqsWSClient.connectSocket()
	if err != nil {
		tiqsWSClient.logger(err)
		return &tiqsWSClient, err
	}

	return &tiqsWSClient, nil
}

// connectSocket establishes a WebSocket connection to the given URL
// It also initializes various processes like ping checking and subscription handling
func (t *TiqsWSClient) connectSocket() error {
	// t.wsURL = wsURL

	dialer := websocket.DefaultDialer
	dialer.ReadBufferSize = 8192 // Increase buffer size (adjust as needed)

	var err error
	for i := 0; i < maxRetries; i++ {
		t.logger(InfoSocketConnecting, ". attempt:", i+1)
		t.socket, _, err = dialer.Dial(t.wsURL, nil)
		if err != nil { // failed to dail
			t.logger(ErrSocketConnection, ". reason:", err)

			// max limit reached. exit
			if i == maxRetries-1 {
				t.logger(INFO_RECONNECT_LIMIT_REACHED)
				close(t.orderChannel)
				close(t.dataChannel)
				return ErrSocketConnection
			}

			time.Sleep(3 * time.Second)
		} else { // dial was successful, break from loop
			break
		}
	}

	t.logger(INFO_CONNECTED_WEBSOCKET)
	// process previous connections
	t.subscribePreviousSubscriptions()
	t.processPendingRequests()
	t.socket.SetReadLimit(1024 * 1024) // Set max message size to 1MB (adjust as needed)
	// ping checker
	go t.startPingChecker()
	// read messages
	go t.readMessages()

	return nil
}

// readMessages continuously reads messages from the WebSocket
// It handles different types of messages, including PING messages
func (t *TiqsWSClient) readMessages() {
	t.logger("Starting read messages")
	defer t.logger("Stopped read messages")
	for {
		select {
		case <-t.stopReadMessagesSig:
			return
		default:
			// read message ---------------------------------------------------------
			_, message, err := t.socket.ReadMessage()
			if err != nil {
				if e, ok := err.(*websocket.CloseError); ok {
					t.logger(ErrReadingSocketMessage, ". reason:", e.Error())
				} else {
					t.logger(ErrReadingSocketMessage, ". reason:", err)
				}
				// reconnect
				t.closeAndReconnect()
				continue
			}

			// decode messages ------------------------------------------------------
			if len(message) == 1 { // ping from server
				t.lastPingTS = time.Now()

			} else if isOrderUpdate(string(message)) { // order update
				update, err := decodeOrderMessage(message)
				if err != nil {
					t.logger(ErrDecodingMessage)
					continue
				}
				t.orderChannel <- update

			} else if len(message) == FULLTICK_LENGTH { // tick update
				tick := t.parseTick(message)
				t.dataChannel <- tick

			} else { // unknown message
				t.logger(fmt.Sprintf("Received message with unexpected length: %d, message: %s", len(message), string(message[:min(50, len(message))])))
			}

		}

	}
}

func (t *TiqsWSClient) closeAndReconnect() {
	go func() {
		// first stop reading messages
		t.stopReadMessagesSig <- true
		t.stopPingListenerSig <- true
		t.CloseConnection()
		t.connectSocket()
	}()
}

// emit sends a message through the WebSocket
// If the socket is not connected, it queues the message (unless volatile is true)
func (t *TiqsWSClient) emit(message interface{}, volatile bool) {

	var msg []byte
	var err error

	switch v := message.(type) {
	case string:
		msg = []byte(v)
	case SocketMessage:
		msg, err = json.Marshal(v)
		if err != nil {
			t.logger(ErrMarshlingMsg)
			return
		}
	default:
		t.logger(ErrUnsupportedMsgType)
		return
	}

	if t.socket != nil { // server connected
		err := t.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			t.logger(ErrEmitingToSocket)
			if !volatile {
				t.pendingQueue = append(t.pendingQueue, message)
			}
		} //else {
		// 	t.logger("⬆ Emitted to socket:", string(msg))
		// }
	} else { // server is not connected
		t.logger(ErrSocketNotConnected)
		if !volatile {
			t.pendingQueue = append(t.pendingQueue, message)
		}
	}
}

// startPingChecker initiates a periodic check to ensure the connection is alive
// If no PING is received within 35 seconds, it triggers a reconnection
func (t *TiqsWSClient) startPingChecker() {
	t.logger("Starting ping checker")
	defer t.logger("Stopped ping checker")

	t.lastPingTS = time.Now()
	window := 10 * time.Second
	ticker := time.NewTicker(window)
	for {
		select {
		case <-t.stopPingListenerSig:
			ticker.Stop()
			return
		case <-ticker.C:
			diff := time.Since(t.lastPingTS)
			if diff > window {
				t.logger(INFO_SOCKET_PING_DIFFERENCE)
				// close and reconnect connection
				t.closeAndReconnect()
			}
		default:
			time.Sleep(1 * time.Second)
		}

	}
}

// processPendingRequests sends any queued messages that couldn't be sent earlier
// due to connection issues
func (t *TiqsWSClient) processPendingRequests() {

	if len(t.pendingQueue) > 0 {
		t.logger(INFO_PROCCESSING_PENDING_REQUESTS)
		for _, request := range t.pendingQueue {
			t.emit(request, false)
		}
		t.pendingQueue = nil
	}
}

// subscribePreviousSubscriptions resubscribes to all previously subscribed topics
// This is useful when reconnecting to ensure all subscriptions are maintained
func (t *TiqsWSClient) subscribePreviousSubscriptions() {
	if len(t.subscriptions) != 0 {
		t.logger(INFO_PROCCESSING_PREVIOUS_SUBSCRIPTION)
		for token := range t.subscriptions {
			t.emit(SocketMessage{
				Code: CODE_SUB,
				Mode: MODE_FULL,
				Full: []int{token},
			}, false)
		}
	}
}

// AddSubscription adds a new subscription to the store
func (t *TiqsWSClient) AddSubscription(token int) {
	t.subscriptions[token] = struct{}{}
	// subscribePreviousSubscriptions()
	t.emit(SocketMessage{
		Code: CODE_SUB,
		Mode: MODE_FULL,
		Full: []int{token},
	}, false)
}

// RemoveSubscription removes a subscription from the store
func (t *TiqsWSClient) RemoveSubscription(token int) {
	delete(t.subscriptions, token)
	t.emit(SocketMessage{
		Code: CODE_UNSUB,
		Mode: MODE_FULL,
		Full: []int{token},
	}, false)
}

// GetSubscriptions returns the current subscriptions
func (t *TiqsWSClient) GetSubscriptions() map[int]struct{} {
	return t.subscriptions
}

// GetDataChannel returns the data channel
func (t *TiqsWSClient) GetDataChannel() <-chan Tick {
	return t.dataChannel
}

// GetOrderChannel returns the order update channel
func (t *TiqsWSClient) GetOrderChannel() <-chan OrderUpdate {
	return t.orderChannel
}

// CloseConnection closes the WebSocket connection
func (t *TiqsWSClient) CloseConnection() {
	t.logger(INFO_CLOSED_WEBSOCKET)
	if t.socket == nil {
		return
	}

	t.socket.Close()
	t.socket = nil

	// Clear pending queue
	t.pendingQueue = nil
}

// bytesToInt32 takes a byte slice as input and parses it into an int32
// The byte slice must be of length 4
// It returns the parsed int32 value
func bytesToInt32(data []byte) int32 {
	// Check if the length of the byte slice is as expected
	if len(data) != 4 {
		// t.logger(ErrInvalidByteSliceLength)
		return 0
	}

	// Parse the byte slice into an int32
	// We use bitwise operations to shift the bytes into their correct positions
	// The result is the parsed int32 value
	value := int32(data[0])<<24 | int32(data[1])<<16 | int32(data[2])<<8 | int32(data[3])

	return value
}

// parseTick takes a byte slice as input and parses it into a Tick struct
// It returns a Tick and a boolean indicating whether the parsing was successful
func (t *TiqsWSClient) parseTick(data []byte) Tick {

	// Create a new Tick struct and fill it with data from the byte slice
	var tick = Tick{
		Token:              bytesToInt32(data[0:4]),               // Token
		LTP:                bytesToInt32(data[4:8]),               // Last traded price
		NetChangeIndicator: int32(data[8]),                        // Net change indicator
		NetChange:          bytesToInt32(data[9:13]),              // Net change
		LTQ:                bytesToInt32(data[13:17]),             // Last traded quantity
		AvgPrice:           bytesToInt32(data[17:21]),             // Average traded price
		TotalBuyQuantity:   bytesToInt32(data[21:25]),             // Total buy quantity
		TotalSellQuantity:  bytesToInt32(data[25:29]),             // Total sell quantity
		Open:               bytesToInt32(data[29:33]),             // Open price
		High:               bytesToInt32(data[33:37]),             // High price
		Close:              bytesToInt32(data[37:41]),             // Close price
		Low:                bytesToInt32(data[41:45]),             // Low price
		Volume:             bytesToInt32(data[45:49]),             // Volume
		LTT:                bytesToInt32(data[49:53]),             // Last traded time
		Time:               bytesToInt32(data[53:57]) + 315513000, // Time
		OI:                 bytesToInt32(data[57:61]),             // Open interest
		OIDayHigh:          bytesToInt32(data[61:65]),             // Open interest day high
		OIDayLow:           bytesToInt32(data[65:69]),             // Open interest day low
		LowerLimit:         bytesToInt32(data[69:73]),             // Lower limit
		UpperLimit:         bytesToInt32(data[73:77]),             // Upper limit
	}

	// Return a Tick
	return tick
}

func (t *TiqsWSClient) logger(msg ...any) {
	if t.enableLog {
		log.Println(msg...)
	}
}

// Check if the keyword 'orderUpdate' exists in the given string
func isOrderUpdate(input string) bool {
	// Use strings.Contains to check for the keyword
	return strings.Contains(input, "orderUpdate")
}

func decodeOrderMessage(message []byte) (OrderUpdate, error) {
	var rawOrder map[string]string
	err := json.Unmarshal(message, &rawOrder)
	if err != nil {
		return OrderUpdate{}, err
	}

	var orderUpdate OrderUpdate

	// Assign string fields directly without checks
	orderUpdate.ID = rawOrder["id"]
	orderUpdate.Type = rawOrder["type"]
	orderUpdate.UserID = rawOrder["userId"]
	orderUpdate.Exchange = rawOrder["exchange"]
	orderUpdate.Symbol = rawOrder["symbol"]
	orderUpdate.Product = rawOrder["product"]
	orderUpdate.Status = rawOrder["status"]
	orderUpdate.ReportType = rawOrder["reportType"]
	orderUpdate.TransactionType = rawOrder["transactionType"]
	orderUpdate.Order = rawOrder["order"]
	orderUpdate.Retention = rawOrder["retention"]
	orderUpdate.Reason = rawOrder["reason"]
	orderUpdate.ExchangeOrderId = rawOrder["exchangeOrderId"]
	orderUpdate.CancelQty = rawOrder["cancelQty"]
	orderUpdate.Tags = rawOrder["tags"]
	orderUpdate.DisclosedQty = rawOrder["disclosedQty"]
	orderUpdate.TriggerPrice = rawOrder["triggerPrice"]

	// Convert Token to int with existence check
	if val, ok := rawOrder["token"]; ok {
		if intVal, err := strconv.Atoi(val); err == nil {
			orderUpdate.Token = intVal
		}
	}

	// Convert Qty to int with existence check
	if val, ok := rawOrder["qty"]; ok {
		if intVal, err := strconv.Atoi(val); err == nil {
			orderUpdate.Qty = intVal
		}
	}

	// Convert Price to float64 with existence check
	if val, ok := rawOrder["price"]; ok {
		if floatVal, err := strconv.ParseFloat(val, 64); err == nil {
			orderUpdate.Price = floatVal
		}
	}

	// Convert AvgPrice to float64 with existence check
	if val, ok := rawOrder["avgPrice"]; ok {
		if floatVal, err := strconv.ParseFloat(val, 64); err == nil {
			orderUpdate.AvgPrice = floatVal
		}
	}

	// Convert Timestamp to time.Time with existence check
	if val, ok := rawOrder["timestamp"]; ok {
		if timeVal, err := strconv.Atoi(val); err == nil {
			orderUpdate.Timestamp = time.Unix(int64(timeVal), 0)
		}
	}

	// Convert exchangeTime to time.Time with existence check
	if val, ok := rawOrder["exchangeTime"]; ok {
		if timeVal, err := time.Parse("02-01-2006 15:04:05", val); err == nil {
			orderUpdate.ExchangeTime = timeVal
		}
	}

	return orderUpdate, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
