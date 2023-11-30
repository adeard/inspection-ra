package plan

import (
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type planHandler struct {
	planService Service
}

func NewPlanHandler(v1 *gin.RouterGroup, planService Service) {

	handler := &planHandler{planService}

	plan := v1.Group("plan")

	plan.GET("", middlewares.AuthService(), handler.GetAll)
}

// @Summary Get Plan
// @Description Get Plan
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {object} domain.PlanResponse{data=domain.PlanData}
// @Router /api/v1/plan [get]
// @Tags Plan
func (h *planHandler) GetAll(c *gin.Context) {
	start := time.Now()
	plan, err := h.planService.GetAll()
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
