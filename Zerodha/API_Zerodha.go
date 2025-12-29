package zerodha

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// PlaceMarketOrder places a market order on the Zerodha platform.
// It takes the exchange, trading symbol, quantity, order type, transaction side, product type and user ID as parameters.
// Returns a PlaceOrderResponse with the order details and an error if any occurs.
func PlaceMarketOrder(
	exchange,
	tradingSymbol,
	quantity,
	orderType,
	transactionSide,
	productType,
	userID_Zerodha string,
) (PlaceOrderResponse, error) {

	// Read access token
	_, apiKey, accessToken, err := ReadingAccessToken_Zerodha(userID_Zerodha)
	if err != nil {
		return PlaceOrderResponse{}, err
	}

	// Form data
	form := url.Values{}
	form.Set("tradingsymbol", tradingSymbol)
	form.Set("exchange", exchange)
	form.Set("transaction_type", transactionSide)
	form.Set("order_type", orderType)
	form.Set("quantity", quantity)
	form.Set("product", productType)
	form.Set("validity", "DAY")

	req, err := http.NewRequest(
		http.MethodPost,
		placeOrderUrl,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return PlaceOrderResponse{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "token "+apiKey+":"+accessToken)
	req.Header.Set("X-Kite-Version", "3")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return PlaceOrderResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PlaceOrderResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return PlaceOrderResponse{}, fmt.Errorf(
			"order placement failed: status=%d body=%s",
			resp.StatusCode,
			body,
		)
	}

	var orderResp PlaceOrderResponse
	if err := json.Unmarshal(body, &orderResp); err != nil {
		return PlaceOrderResponse{}, err
	}

	if orderResp.Status != apiResponseStatus.SUCCESS {
		return orderResp, fmt.Errorf("order rejected: %s", orderResp.Message)
	}

	return orderResp, nil
}

// GetFunds fetches the funds from the Zerodha API.
// It takes a user ID as a parameter and returns the funds response and an error if something goes wrong.
func GetFunds(userID string) (FundsResponse, error) {
	// Read access token
	_, apiKey, accessToken, err := ReadingAccessToken_Zerodha(userID)
	if err != nil {
		return FundsResponse{}, err
	}

	// Create a new HTTP GET request
	req, err := http.NewRequest(
		http.MethodGet,
		marginsURL,
		nil,
	)
	if err != nil {
		return FundsResponse{}, err
	}

	// Set the headers of the request
	req.Header.Set("X-Kite-Version", "3")
	req.Header.Set("Authorization", "token "+apiKey+":"+accessToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return FundsResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return FundsResponse{}, err
	}

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return FundsResponse{}, fmt.Errorf(
			"funds api failed: status=%d body=%s",
			resp.StatusCode,
			body,
		)
	}

	// Unmarshal the response body into the FundsResponse struct
	var fundsResp FundsResponse
	if err := json.Unmarshal(body, &fundsResp); err != nil {
		return FundsResponse{}, err
	}

	return fundsResp, nil
}

// GetHoldings fetches the holdings from the Zerodha API.
// It takes a user ID as a parameter and returns the holdings response and an error if something goes wrong.
func GetHoldings(userID string) (HoldingsResponse, error) {
	// Read access token
	_, apiKey, accessToken, err := ReadingAccessToken_Zerodha(userID)
	if err != nil {
		return HoldingsResponse{}, err
	}

	// Create a new HTTP GET request
	req, err := http.NewRequest(
		http.MethodGet,
		holdingsURL,
		nil,
	)
	if err != nil {
		return HoldingsResponse{}, err
	}

	// Set the headers of the request
	req.Header.Set("X-Kite-Version", "3")
	req.Header.Set("Authorization", "token "+apiKey+":"+accessToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return HoldingsResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return HoldingsResponse{}, err
	}

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return HoldingsResponse{}, fmt.Errorf(
			"holdings api failed: status=%d body=%s",
			resp.StatusCode,
			body,
		)
	}

	// Unmarshal the response body into the HoldingsResponse struct
	var holdingResp HoldingsResponse
	if err := json.Unmarshal(body, &holdingResp); err != nil {
		return HoldingsResponse{}, err
	}

	return holdingResp, nil
}

// GetPositions retrieves the positions for a given user from the Zerodha API.
// It takes the user ID as a parameter and returns the positions response and an error if any occurs.
func GetPositions(userID string) (PositionsResponse, error) {
	// Read access token
	_, apiKey, accessToken, err := ReadingAccessToken_Zerodha(userID)
	if err != nil {
		return PositionsResponse{}, err
	}

	// Create a new HTTP GET request
	req, err := http.NewRequest(
		http.MethodGet,
		positionsURL,
		nil,
	)
	if err != nil {
		return PositionsResponse{}, err
	}

	// Set the headers of the request
	req.Header.Set("X-Kite-Version", "3")
	req.Header.Set("Authorization", "token "+apiKey+":"+accessToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return PositionsResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PositionsResponse{}, err
	}

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return PositionsResponse{}, fmt.Errorf(
			"positions api failed: status=%d body=%s",
			resp.StatusCode,
			body,
		)
	}

	// Unmarshal the response body into the PositionsResponse struct
	var positionsResp PositionsResponse
	if err := json.Unmarshal(body, &positionsResp); err != nil {
		return PositionsResponse{}, err
	}

	return positionsResp, nil
}
