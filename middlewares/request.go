package middlewares

import (
	"bytes"
	"fmt"
	"inspection-ra/config"
	"inspection-ra/domain"
	"inspection-ra/helpers"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			c.AbortWithStatus(500)
			return
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		recorder := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = recorder

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()
		location, _ := time.LoadLocation("Local")
		currentDateTime := time.Now().In(location).Format("2006-01-02 15:04:05")

		requestLog := domain.RequestLogRequest{
			Method:     c.Request.Method,
			Source:     c.Request.Host + c.Request.RequestURI,
			ClientIp:   c.ClientIP(),
			UserAgent:  c.Request.UserAgent(),
			StatusCode: status,
			Duration:   fmt.Sprint(duration),
			Request:    string(bodyBytes),
			CreatedAt:  currentDateTime,
			UpdatedAt:  currentDateTime,
			Response:   recorder.body.String(),
		}

		db := config.Connect()

		logData := domain.MailboxRequest{
			StatusCode: status,
			Source:     c.Request.Host + c.Request.RequestURI,
			Request:    string(bodyBytes),
			Response:   recorder.body.String(),
		}

		if err := db.Create(&logData).Error; err != nil {
			log.Printf("Error saving request log: %v", err)

			requestLog.Response = err.Error()

			go helpers.SendErrorLog(requestLog)
		}

		if status != 200 {
			go helpers.SendErrorLog(requestLog)
		}
	}
}
