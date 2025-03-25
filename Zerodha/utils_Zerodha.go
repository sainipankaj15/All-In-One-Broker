package zerodha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadingAccessToken_Zerodha(userID_Zerodha string) (string, string, string, error) {
	fileName := userID_Zerodha + `.json`

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", "", "", err
	}

	var fileData readDataJsonZerodha

	err = json.Unmarshal(fileContent, &fileData)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON in ReadingAccessToken_Zerodha()")
		return "", "", "", err
	}

	return fileData.Date, fileData.ApiKey, fileData.AccessToken, nil
}
