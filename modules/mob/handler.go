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
	mob.POST("damaged", middlewares.AuthService(), handler.StoreDamaged)
	mob.POST("damaged/batch", middlewares.AuthService(), handler.InsertDamaged)
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

	for i, val := range mob {

		damDateParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", val.DamDate)
		damDate := damDateParse.Format("2006-01-02")
		if damDate == "1900-01-01" {
			damDate = ""
		}

		planDateParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", val.PlanDate)
		planDate := planDateParse.Format("2006-01-02")
		if planDate == "1900-01-01" {
			planDate = ""
		}

		createdDateParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", val.CreatedDate)
		createdDate := createdDateParse.Format("2006-01-02")
		if createdDate == "1900-01-01" {
			createdDate = ""
		}

		createdTimeParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", val.CreatedTime)
		createdTime := createdTimeParse.Format("15:04:05")
		if createdTime == "00:00:00" {
			createdTime = ""
		}

		mob[i].DamDate = damDate
		mob[i].PlanDate = planDate
		mob[i].CreatedDate = createdDate
		mob[i].CreatedTime = createdTime
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

// @Summary Store Mob Item Damage
// @Description Store Mob Item Damage
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param MobItemDamagedRequest body domain.MobItemDamagedRequest true " MobItemDamagedRequest Schema "
// @Produce  json
// @Success 200 {object} domain.MobResponse{data=domain.MobItemDamagedRequest}
// @Router /api/v1/mob/damaged [post]
// @Tags Mob
func (h *mobHandler) StoreDamaged(c *gin.Context) {
	start := time.Now()
	var input domain.MobItemDamagedRequest

	c.ShouldBindJSON(&input)

	res, err := h.mobService.StoreDamaged(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MobResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.MobResponse{
		Data:        res,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Insert Batch Mob Item Damage
// @Description Insert Batch Mob Item Damage
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param MobItemDamagedRequest body []domain.MobItemDamagedRequest true " MobItemDamagedRequest Schema "
// @Produce  json
// @Success 200 {object} domain.MobResponse
// @Router /api/v1/mob/damaged/batch [post]
// @Tags Mob
func (h *mobHandler) InsertDamaged(c *gin.Context) {
	start := time.Now()
	var input []domain.MobItemDamagedRequest

	c.ShouldBindJSON(&input)

	res, err := h.mobService.InsertDamaged(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MobResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.MobResponse{
		Data:        res,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
