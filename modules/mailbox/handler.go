package mailbox

import (
	"fmt"
	"inspection-ra/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type mailboxHandler struct {
	loggingService Service
}

func NewMailboxHandler(v1 *gin.RouterGroup, loggingService Service) {

	handler := &mailboxHandler{loggingService}

	log := v1.Group("mailbox")
	log.POST("", handler.Create)
}

// @Summary Create Log
// @Description Create Log
// @Accept  json
// @Param MailboxRequest body domain.MailboxRequest true " MailboxRequest Schema "
// @Produce  json
// @Success 200 {object} domain.MailboxResponse{data=domain.MailboxRequest}
// @Router /api/v1/mailbox [post]
// @Tags Mailbox
func (h *mailboxHandler) Create(c *gin.Context) {
	start := time.Now()
	logInput := domain.MailboxRequest{}

	err := c.ShouldBindJSON(&logInput)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})

		return
	}

	log, err := h.loggingService.Store(logInput)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.MailboxResponse{
		Data:        log,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
