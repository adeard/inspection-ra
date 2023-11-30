package mob

import (
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type mobHandler struct {
	mobService Service
}

func NewMobHandler(v1 *gin.RouterGroup, mobService Service) {

	handler := &mobHandler{mobService}

	mob := v1.Group("mob")

	mob.GET("", middlewares.AuthService(), handler.GetAll)
}

// @Summary Get Mob
// @Description Get Mob
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {object} domain.MobResponse{data=domain.MobData}
// @Router /api/v1/mob [get]
// @Tags Mob
func (h *mobHandler) GetAll(c *gin.Context) {
	start := time.Now()
	mob, err := h.mobService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MobResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.MobResponse{
		Data:        mob,
		ElapsedTime: fmt.Sprint(time.Since(start))},
	)
}
