package attachment

import (
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/middlewares"
	"net/http"
	"os"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

type attachmentHandler struct {
	attachmentService Service
}

func NewAttachmentHandler(v1 *gin.RouterGroup, attachmentService Service) {

	handler := &attachmentHandler{attachmentService}

	attachment := v1.Group("attachment")

	attachment.POST("", middlewares.AuthService(), handler.Store)
}

// @Summary Upload File
// @Description Upload File
// @Accept  mpfd
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param file formData file true " AttachmentRequest Schema "
// @Param no_inspec formData string true " AttachmentRequest Schema "
// @Param image_category formData string true " AttachmentRequest Schema "
// @Produce  json
// @Success 200 {object} domain.AttachmentResponse
// @Router /api/v1/attachment [post]
// @Tags Attachment
func (h *attachmentHandler) Store(c *gin.Context) {
	start := time.Now()
	var input domain.AttachmentRequest

	c.ShouldBind(&input)

	path, _ := os.Getwd()

	file, _ := c.FormFile("file")

	openedFile, _ := file.Open()

	ext, _ := mimetype.DetectReader(openedFile)

	filename := input.ImageCategory + "_" + time.Now().Format("20060102150405") + ext.Extension()

	dst := path + `\images\` + filename

	input.Filename = filename

	result, err := h.attachmentService.Store(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.AttachmentResponse{
			Data:        err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.AttachmentResponse{
			Data:        err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	c.JSON(http.StatusOK, domain.AttachmentResponse{
		Data:        result,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
