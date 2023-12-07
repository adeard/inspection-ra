package vehicletype

import (
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type vehicleTypeHandler struct {
	vehicleTypeService Service
}

func NewVehicleTypeHandler(v1 *gin.RouterGroup, vehicleTypeService Service) {

	handler := &vehicleTypeHandler{vehicleTypeService}

	vehicletype := v1.Group("vehicletype")

	vehicletype.GET("", middlewares.AuthService(), handler.GetAll)
}

// @Summary Get Vehicle Type
// @Description Get Vehicle Type
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {object} domain.VehicleTypeResponse{data=domain.VehicleTypeData}
// @Router /api/v1/vehicletype [get]
// @Tags Vehicle Type
func (h *vehicleTypeHandler) GetAll(c *gin.Context) {
	start := time.Now()
	vehicletype, err := h.vehicleTypeService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.VehicleTypeResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.VehicleTypeResponse{
		Data:        vehicletype,
		ElapsedTime: fmt.Sprint(time.Since(start))},
	)
}
