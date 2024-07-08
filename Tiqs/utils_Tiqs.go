package tiqs

import (
	"encoding/json"
	"io/ioutil"

	easyConversion "github.com/sainipankaj15/data-type-conversion"
)

func readingAccessToken_Tiqs(userID_Tiqs string) (string, string, error) {

	fileName := userID_Tiqs + `.json`

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", "", err
	}

	var fileData ReadDataJsonTiqs

	err = json.Unmarshal(fileContent, &fileData)
	if err != nil {
		return "", "", err
	}
	accessToken := fileData.AccessToken
	APPID := fileData.APPID

	return accessToken, APPID, nil
}

func CurrentQtyForAnySymbol_Tiqs(symbol string, productType string, UserId_Tiqs string) (string, error) {

	PositionAPIResp_Tiqs, err := PositionApi_Tiqs(UserId_Tiqs)

	if err != nil {
		return "", err
	}

	for i := 0; i < len(PositionAPIResp_Tiqs.NetPosition_Tiqss); i++ {
		position := PositionAPIResp_Tiqs.NetPosition_Tiqss[i]

		if position.Symbol == symbol {
			if position.Product == productType {
				return position.Qty, nil
			}
		}
	}

	return "0", nil
}

func ExitAllPosition_Tiqs(UserId_Tiqs string) (string, error) {

	PositionAPIResp_Tiqs, err := PositionApi_Tiqs(UserId_Tiqs)

	if err != nil {
		return "failed", err
	}

	for i := 0; i < len(PositionAPIResp_Tiqs.NetPosition_Tiqss); i++ {
		position := PositionAPIResp_Tiqs.NetPosition_Tiqss[i]

		go func(pos NetPosition_Tiqs) {
			buyQtyInString := position.DayBuyQty
			sellQtyInString := position.DaySellQty

			buyQty := easyConversion.StringToInt(buyQtyInString)
			sellQty := easyConversion.StringToInt(sellQtyInString)

			diff := buyQty - sellQty

			if diff > 0 {
				// it means long position : Have to cut it by opposite order
				qtyInString := easyConversion.IntToString(diff)
				go OrderPlaceMarket_Tiqs(position.Exchange, position.Token, qtyInString, "S", position.Product, UserId_Tiqs)
			} else if diff < 0 {
				// it means short position : Have to cut it by opposite order
				qtyInString := easyConversion.IntToString(diff)
				go OrderPlaceMarket_Tiqs(position.Exchange, position.Token, qtyInString, "B", position.Product, UserId_Tiqs)
			}

		}(position)
	}

	return "success", nil
}

func ExitByPositionID_Tiqs(symbol string, productType string, UserId_Tiqs string) error {

	PositionAPIResp_Tiqs, err := PositionApi_Tiqs(UserId_Tiqs)

	if err != nil {
		return err
	}

	for i := 0; i < len(PositionAPIResp_Tiqs.NetPosition_Tiqss); i++ {
		position := PositionAPIResp_Tiqs.NetPosition_Tiqss[i]

		go func(pos NetPosition_Tiqs) {

			if position.Symbol == symbol {
				if position.Product == productType {

					buyQtyInString := position.DayBuyQty
					sellQtyInString := position.DaySellQty

					buyQty := easyConversion.StringToInt(buyQtyInString)
					sellQty := easyConversion.StringToInt(sellQtyInString)

					diff := buyQty - sellQty

					if diff > 0 {
						// it means long position : Have to cut it by opposite order
						qtyInString := easyConversion.IntToString(diff)
						OrderPlaceMarket_Tiqs(position.Exchange, position.Token, qtyInString, "S", position.Product, UserId_Tiqs)
					} else if diff < 0 {
						// it means short position : Have to cut it by opposite order
						qtyInString := easyConversion.IntToString(diff)
						OrderPlaceMarket_Tiqs(position.Exchange, position.Token, qtyInString, "B", position.Product, UserId_Tiqs)
					}
				}
			}
		}(position)
	}
	return nil
}
