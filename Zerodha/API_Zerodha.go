package zerodha

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

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
