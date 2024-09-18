package notiftype

import (
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type notifTypeHandler struct {
	notifTypeService Service
}

func NewNotifTypeHandler(v1 *gin.RouterGroup, notifTypeService Service) {

	handler := &notifTypeHandler{notifTypeService}

	notifType := v1.Group("notiftype")

	notifType.GET("", middlewares.AuthService(), handler.GetAll)
}

// @Summary Get Notif Type
// @Description Get Notif Type
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param NotifTypeRequest query domain.NotifTypeRequest true " NotifTypeRequest Schema "
// @Produce  json
// @Success 200 {object} domain.RunAcctResponse{data=[]domain.NotifTypeData}
// @Router /api/v1/notiftype [get]
// @Tags Notif Type
func (h *notifTypeHandler) GetAll(c *gin.Context) {
	start := time.Now()

	input := domain.NotifTypeRequest{}

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	runAcct, err := h.notifTypeService.GetAll(input)
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
