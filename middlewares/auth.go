package middlewares

import (
	"fmt"
	"inspection-ra/helpers"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthService() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := helpers.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			context.Abort()
			return
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
		urlAddr := "https://user-api-simp.azurewebsites.net/api/APP782231326/verify_save"
		req, err := http.NewRequest("GET", urlAddr, nil)
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
		client := helpers.RequestClient(&http.Client{})
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
