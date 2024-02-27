package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func LogInit(message string) {

	now := time.Now().Format("2006-01-02")

	f, err := os.OpenFile("logs/"+now+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(message)
}

func SendLog(context *gin.Context, request any, statusCode int, response string) (string, error) {

	var requestString string

	result, err := json.Marshal(request)
	if err != nil {
		requestString = "request not struct : " + err.Error()
	} else {
		requestString = string(result)
	}

	if _, ok := request.(string); ok {
		requestString = fmt.Sprintf("%v", request)
	}

	errorLog := struct {
		Status_code int
		Source      string
		Request     string
		Response    string
	}{
		statusCode,
		context.Request.Host + context.Request.URL.String(),
		requestString,
		response,
	}

	toJson, _ := json.Marshal(errorLog)

	urladdr := "http://dev.indoagri.co.id/LogSystemApi/api/v1/log"
	req, err := http.NewRequest("POST", urladdr, bytes.NewBuffer([]byte(toJson)))
	if err != nil {
		return " ", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return string(body), err
}
