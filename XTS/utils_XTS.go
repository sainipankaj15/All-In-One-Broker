package xts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func ReadingAccessToken_XTS(userid_XTS string) (string, string, error) {
	fileName := userid_XTS + `.json`

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", "", err
	}

	var fileData readDataJsonXTS

	err = json.Unmarshal(fileContent, &fileData)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON in ReadingAccessToken for XTS)")
		return "", "", err
	}

	return fileData.Date, fileData.AccessToken, nil
}

func StringToInt(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("Error while converting string to int in StringToInt()")
		return 0
	}
	return num
}
