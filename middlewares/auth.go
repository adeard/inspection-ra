package middlewares

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthService() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := getTokenFromRequestSample(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		// url := "https://user-api-simp.azurewebsites.net/api/verify"
		urladdr := os.Getenv("USER_SERVICE_URL") + "/api/verify"
		req, err := http.NewRequest("GET", urladdr, nil)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		bearer := "Bearer " + token
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
		req.Header.Add("Accept", "application/json")

		// Send req using http Client
		proxyUrl, _ := url.Parse("http://10.126.111.123:4480")
		client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
		// client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		} else {
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusAccepted {
				data, _ := ioutil.ReadAll(resp.Body)
				var message struct {
					Message string
				}

				json.Unmarshal(data, &message)

				context.JSON(http.StatusForbidden, gin.H{"error": message.Message})
				context.Abort()
				return
			}
		}
		context.Next()
	}
}

func AuthService_Save() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := getTokenFromRequestSample(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		url := "https://user-api-simp.azurewebsites.net/api/APP940232184/verify_save"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		bearer := "Bearer " + token
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
		req.Header.Add("Accept", "application/json")

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		} else {
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusAccepted {
				data, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(string(data))
				context.JSON(http.StatusForbidden, gin.H{"error": "User not allowed"})
				context.Abort()
				return
			}
		}
		context.Next()
	}
}
