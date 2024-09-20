package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"time"
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

func CurrentDate() string {

	// Set the time zone to Indian Standard Time (IST)
	ist, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Println("Error loading IST location:", err)
		return ""
	}

	// Get the current date adn time time in IST
	currentTimeIST := time.Now().In(ist)

	// Format the date as a string
	istFormat := "2006-01-02"
	formattedDate := currentTimeIST.Format(istFormat)

	return formattedDate
}

func CurrentTime() string {

	// Set the time zone to Indian Standard Time (IST)
	ist, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Println("Error loading IST location:", err)
		return ""
	}

	// Get the current time in IST
	currentTimeIST := time.Now().In(ist)

	// Format the time as a string
	istFormat := "15:04:05"
	formattedTime := currentTimeIST.Format(istFormat)

	return formattedTime
}

func ApplicationStart(StartingHour, StartingMinutes, StartingSeconds int) {

	// Load IST location (India)
	ist, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Println("Error loading IST location:", err)
		return
	}

	// Specify the target time in IST (For Example 09:30:02 PM) : Application start time
	targetTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), StartingHour, StartingMinutes, StartingSeconds, 0, ist)

	// Calculate the duration until the target time
	durationUntilTarget := time.Until(targetTime)

	// Put sleep for that duration
	time.Sleep(durationUntilTarget)

	// Sleeping Done, Now we can Resume our application
}

func ApplicationClosing(ClosingHour, ClosingMinutes, ClosingSeconds int, isWorkDone chan<- time.Time) {

	// Load IST location (India)
	ist, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Println("Error loading IST location:", err)
		return
	}

	// Specify the target time in IST (3:30:45 PM)
	targetTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), ClosingHour, ClosingMinutes, ClosingSeconds, 0, ist)

	for {
		// Calculate the duration until the target time
		durationUntilTarget := time.Until(targetTime)

		// Put sleep for that duration
		time.Sleep(durationUntilTarget)

		// When Sleeping Done, Push current time to channel
		isWorkDone <- time.Now()
		break
	}
}

type Number interface {
	int | int32 | int64 | float32 | float64
}

// RoundOff works with any Number type that satisfies the Number constraint
func RoundOff[T Number](a, b T) T {
	if b == 0 {
		return a // Avoid division by zero
	}

	// Perform the rounding
	rounded := T(math.Round(float64(a)/float64(b)) * float64(b))

	return rounded
}
