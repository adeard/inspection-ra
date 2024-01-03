package mob

import (
	"encoding/json"
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/middlewares"
	"io/ioutil"
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
	mob.POST("batch", middlewares.AuthService(), handler.Insert)
}

// @Summary Get Mob
// @Description Get Mob
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param MobRequest query domain.MobRequest true " MobRequest Schema "
// @Produce  json
// @Success 200 {object} domain.MobResponse{data=domain.MobData}
// @Router /api/v1/mob [get]
// @Tags Mob
func (h *mobHandler) GetAll(c *gin.Context) {
	start := time.Now()
	mobRequest := domain.MobRequest{}

	c.ShouldBind(&mobRequest)

	mob, err := h.mobService.GetAll(mobRequest)
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

// @Summary Insert Inspection
// @Description Insert Inspection
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param MobRequest body []domain.MobRequest true " MobRequest Schema "
// @Produce  json
// @Success 200 {object} domain.MobResponse{data=[]domain.MobData}
// @Router /api/v1/mob/batch [post]
// @Tags Mob
func (h *mobHandler) Insert(c *gin.Context) {
	start := time.Now()
	mobRequest := []domain.MobRequest{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &mobRequest)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	mob, err := h.mobService.Insert(mobRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MobResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.MobResponse{
		Data:        mob,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
