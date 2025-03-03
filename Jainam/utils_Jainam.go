package jainam

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadingAccessToken_Jainam(userID_Jainam string) (string, string, error) {
	fileName := userID_Jainam + `.json`

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", "", err
	}

	var fileData readDataJsonJainam

	err = json.Unmarshal(fileContent, &fileData)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON in ReadingAccessToken_Jainam()")
		return "", "", err
	}

	return fileData.Date, fileData.AccessToken, nil
}
