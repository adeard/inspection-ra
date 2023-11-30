package runacct

import (
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type runAcctHandler struct {
	runAcctService Service
}

func NewRunAcctHandler(v1 *gin.RouterGroup, runAcctService Service) {

	handler := &runAcctHandler{runAcctService}

	runAcct := v1.Group("runacct")

	runAcct.GET("", middlewares.AuthService(), handler.GetAll)
}

// @Summary Get Running Account
// @Description Get Running Account
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {object} domain.RunAcctResponse{data=domain.RunAcctData}
// @Router /api/v1/runacct [get]
// @Tags Running Account
func (h *runAcctHandler) GetAll(c *gin.Context) {
	start := time.Now()
	runAcct, err := h.runAcctService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.RunAcctResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.RunAcctResponse{
		Data:        runAcct,
		ElapsedTime: fmt.Sprint(time.Since(start))},
	)
}
