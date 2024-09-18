package auth_v2

import (
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authHandler struct {
	authService Service
}

func NewAuthHandler(v2 *gin.RouterGroup, authService Service) {

	handler := &authHandler{authService}

	auth := v2.Group("auth")

	auth.POST("sign_in", handler.SignIn)
	auth.GET("logged", middlewares.AuthService(), handler.GetLogged)
}

// @Summary Sign In
// @Description Sign In
// @Accept  json
// @Param AuthRequest body domain.AuthRequest true " AuthRequest Schema "
// @Produce  json
// @Success 200 {object} domain.AuthResponse{data=domain.AuthData}
// @Router /api/v2/auth/sign_in [post]
// @Tags Auth V2
func (h *authHandler) SignIn(c *gin.Context) {
	start := time.Now()
	var authrequest domain.AuthRequest

	err := c.ShouldBindJSON(&authrequest)
	if err != nil {

		errorMessages := []string{}

		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s , condition : %s", v.Field(), v.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, domain.AuthResponse{
			Data:        errorMessages,
			Message:     "Error Validation",
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	auth, err := h.authService.SignIn(authrequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.AuthResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	if auth.Token == "" {
		c.JSON(http.StatusBadRequest, domain.AuthResponse{
			Message:     auth.Message,
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.AuthResponse{
		Data:        auth,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Get Logged User
// @Description Get Logged User
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {object} domain.AuthResponse{data=domain.AuthLoggedData}
// @Router /api/v2/auth/logged [get]
// @Tags Auth V2
func (h *authHandler) GetLogged(c *gin.Context) {
	start := time.Now()
	bearerToken := c.Request.Header.Get("Authorization")

	auth, err := h.authService.GetLogged(bearerToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.AuthResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	result := domain.AuthResponse{
		Data:        auth,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	}

	c.JSON(http.StatusOK, result)
}
