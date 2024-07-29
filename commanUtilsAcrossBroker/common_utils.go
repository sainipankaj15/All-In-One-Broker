package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
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

func UploadFileToTelegram(botToken, chatID, filePath string) error {
	// Telegram API endpoint for sending files
	url := "https://api.telegram.org/bot" + botToken + "/sendDocument"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new buffer to store the file contents
	fileContents := &bytes.Buffer{}

	// Create a multipart writer
	writer := multipart.NewWriter(fileContents)

	// Create a new form file field
	part, err := writer.CreateFormFile("document", filePath)
	if err != nil {
		return err
	}

	// Copy the file contents to the form file field
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	// Close the multipart writer
	writer.Close()

	// Create a new HTTP request
	request, err := http.NewRequest("POST", url, fileContents)
	if err != nil {
		return err
	}

	// Set the content type header
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Add chat_id parameter
	query := request.URL.Query()
	query.Add("chat_id", chatID)
	request.URL.RawQuery = query.Encode()

	// Send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check the response
	if response.StatusCode != http.StatusOK {
		return nil
	}

	return nil
}
