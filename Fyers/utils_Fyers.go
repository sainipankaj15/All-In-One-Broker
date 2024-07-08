package fyers

import (
	"encoding/json"
	"io/ioutil"
)

func readingAccessToken_Fyers(userFyersID string) (string, error) {

	fileName := userFyersID + `.json`

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	var fileData ReadDataJson_Fyers

	err = json.Unmarshal(fileContent, &fileData)
	if err != nil {
		return "", err
	}
	accessToken := fileData.AccessTokenWithAppID

	return accessToken, nil
}
