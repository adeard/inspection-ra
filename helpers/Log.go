package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"inspection-ra/config"
	"inspection-ra/domain"
	"io"
	"log"
	"net/http"
	"time"
)

func SendErrorLog(errorLog domain.RequestLogRequest) (string, error) {
	start := time.Now()

	toJson, _ := json.Marshal(errorLog)

	urladdr := "http://dev.indoagri.co.id/LogSystemApi/api/v1/log"
	req, err := http.NewRequest("POST", urladdr, bytes.NewBuffer([]byte(toJson)))
	if err != nil {
		return " ", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := RequestClient(&http.Client{})
	resp, err := client.Do(req)
	if err != nil {
		duration := time.Since(start)
		location, _ := time.LoadLocation("Local")
		currentDateTime := time.Now().In(location).Format("2006-01-02 15:04:05")

		requestLog := domain.RequestLogRequest{
			Method:     "POST",
			Source:     urladdr,
			ClientIp:   "127.0.0.1",
			UserAgent:  errorLog.UserAgent,
			StatusCode: resp.StatusCode,
			Duration:   fmt.Sprint(duration),
			Request:    string(toJson),
			CreatedAt:  currentDateTime,
			UpdatedAt:  currentDateTime,
			Response:   err.Error(),
		}

		db := config.Connect()
		if err := db.Create(&requestLog).Error; err != nil {
			log.Printf("Error saving request log: %v", err)
		}
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return string(body), err
}
