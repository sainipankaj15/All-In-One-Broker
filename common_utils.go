package allInOneBroker

import (
	"encoding/json"
	"io/ioutil"
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
