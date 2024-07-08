package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func TelegramSend(botToken, chatID, text string) {

	requestURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	client := &http.Client{}

	values := map[string]string{"text": text, "chat_id": chatID}
	jsonParameters, err := json.Marshal(values)
	if err != nil {
		log.Println("Error marshaling JSON : ", err)
		return
	}

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonParameters))
	if err != nil {
		log.Println("Error creating request in telegram", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}
	log.Println("Telegram Call Response Body:", string(body))
}
