package plan

import (
	"encoding/json"
	"fmt"
	"inspection-ra/domain"
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

	c.ShouldBindJSON(&planRequest)

	plan, err := h.planService.GetAll(planRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.PlanResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

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

		c.JSON(http.StatusBadRequest, domain.PlanResponse{
			Data:        errorMessages,
			Message:     "Error Validation",
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	plan, err := h.planService.Store(planRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.PlanResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

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
		c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &planRequest)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	plan, err := h.planService.Insert(planRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.PlanResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.PlanResponse{
		Data:        plan,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
