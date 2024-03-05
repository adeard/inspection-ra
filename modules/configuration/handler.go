package configuration

import (
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/helpers"
	"inspection-ra/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type configurationHandler struct {
	configurationService Service
}

func NewConfigurationHandler(v1 *gin.RouterGroup, configurationService Service) {

	handler := &configurationHandler{configurationService}

	configuration := v1.Group("config")

	configuration.GET("check", middlewares.AuthService(), handler.CheckConfig)
}

// @Summary Check Current Config
// @Description Check Current Config
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {object} domain.ConfigurationResponse{data=domain.ConfigurationData}
// @Router /api/v1/config/check [get]
// @Tags Inspect Configuration
func (h *configurationHandler) CheckConfig(c *gin.Context) {
	start := time.Now()
	var configurationRequest domain.ConfigurationRequest

	c.ShouldBindJSON(&configurationRequest)

	configuration, err := h.configurationService.CheckConfig(configurationRequest)

	if err != nil {

		helpers.LogInit(err.Error())

		go helpers.SendLogLocal(c, configurationRequest, http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, domain.ConfigurationResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	go helpers.SendLogLocal(c, configurationRequest, http.StatusOK, "")

	c.JSON(http.StatusOK, domain.ConfigurationResponse{
		Data:        configuration,
		ElapsedTime: fmt.Sprint(time.Since(start))},
	)
}
