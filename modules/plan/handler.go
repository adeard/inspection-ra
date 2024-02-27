package plan

import (
	"encoding/json"
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/helpers"
	"inspection-ra/middlewares"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type planHandler struct {
	planService Service
}

func NewPlanHandler(v1 *gin.RouterGroup, planService Service) {

	handler := &planHandler{planService}

	plan := v1.Group("plan")

	plan.GET("", middlewares.AuthService(), handler.GetAll)
	plan.POST("", middlewares.AuthService(), handler.Store)
	plan.POST("batch", middlewares.AuthService(), handler.Insert)
}

// @Summary Get Plan
// @Description Get Plan
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param PlanRequest query domain.PlanRequest true " PlanRequest Schema "
// @Produce  json
// @Success 200 {object} domain.PlanResponse{data=domain.PlanData}
// @Router /api/v1/plan [get]
// @Tags Plan
func (h *planHandler) GetAll(c *gin.Context) {
	start := time.Now()
	var planRequest domain.PlanRequest

	c.ShouldBind(&planRequest)

	plan, err := h.planService.GetAll(planRequest)
	if err != nil {

		helpers.LogInit(err.Error())

		c.JSON(http.StatusBadRequest, domain.PlanResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	for i, val := range plan {
		planDateParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", val.PlanDate)
		planDate := planDateParse.Format("2006-01-02")

		inspectDateParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", val.InspectDate)
		inspectDate := inspectDateParse.Format("2006-01-02")
		if inspectDate == "1900-01-01" {
			inspectDate = ""
		}

		inspectTimeParse, _ := time.Parse("2006-01-02T15:04:05Z07:00", val.InspectTime)
		inspectTime := inspectTimeParse.Format("15:04:05")
		if inspectTime == "00:00:00" {
			inspectTime = ""
		}

		plan[i].PlanDate = planDate
		plan[i].InspectDate = inspectDate
		plan[i].InspectTime = inspectTime
	}

	go helpers.SendLogLocal(c, planRequest, http.StatusOK, "")

	c.JSON(http.StatusOK, domain.PlanResponse{
		Data:        plan,
		ElapsedTime: fmt.Sprint(time.Since(start))},
	)
}

// @Summary Store Plan
// @Description Store Plan
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param PlanRequest body domain.PlanRequest true " PlanRequest Schema "
// @Produce  json
// @Success 200 {object} domain.PlanResponse{data=domain.PlanData}
// @Router /api/v1/plan [post]
// @Tags Plan
func (h *planHandler) Store(c *gin.Context) {
	start := time.Now()
	var planRequest domain.PlanRequest

	c.ShouldBindJSON(&planRequest)

	validate := validator.New()
	err := validate.Struct(planRequest)
	if err != nil {
		errorMessages := []interface{}{}

		for _, v := range err.(validator.ValidationErrors) {
			errorArray := map[string]string{
				"field":   v.Field(),
				"message": v.ActualTag(),
			}

			errorMessages = append(errorMessages, errorArray)
		}

		go helpers.SendLogLocal(c, planRequest, http.StatusBadRequest, "Error Validation")

		c.JSON(http.StatusBadRequest, domain.PlanResponse{
			Data:        errorMessages,
			Message:     "Error Validation",
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	plan, err := h.planService.Store(planRequest)
	if err != nil {

		helpers.LogInit(err.Error())

		go helpers.SendLogLocal(c, planRequest, http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, domain.PlanResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	go helpers.SendLogLocal(c, planRequest, http.StatusOK, "")

	c.JSON(http.StatusOK, domain.PlanResponse{
		Data:        plan,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}

// @Summary Insert Plan
// @Description Insert Plan
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param PlanRequest body []domain.PlanRequest true " PlanRequest Schema "
// @Produce  json
// @Success 200 {object} domain.PlanResponse{data=[]domain.PlanData}
// @Router /api/v1/plan/batch [post]
// @Tags Plan
func (h *planHandler) Insert(c *gin.Context) {
	start := time.Now()
	planRequest := []domain.PlanRequest{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {

		go helpers.SendLog(c, planRequest, http.StatusBadRequest, err.Error())
		go helpers.SendLogLocal(c, planRequest, http.StatusBadRequest, err.Error())

		c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &planRequest)
	if err != nil {

		go helpers.SendLog(c, planRequest, http.StatusBadRequest, err.Error())
		go helpers.SendLogLocal(c, planRequest, http.StatusBadRequest, err.Error())

		c.AbortWithError(400, err)
		return
	}

	plan, err := h.planService.Insert(planRequest)
	if err != nil {

		helpers.LogInit(err.Error())

		go helpers.SendLog(c, planRequest, http.StatusBadRequest, err.Error())
		go helpers.SendLogLocal(c, planRequest, http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, domain.PlanResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	go helpers.SendLogLocal(c, planRequest, http.StatusOK, "")

	c.JSON(http.StatusOK, domain.PlanResponse{
		Data:        plan,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
