package attachment

import (
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/helpers"
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
	attachment.POST("db", middlewares.AuthService(), handler.StoreFile)
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

	// path, _ := os.Getwd()

	file, _ := c.FormFile("file")

	openedFile, err := file.Open()
	if err != nil {
		helpers.LogInit(err.Error())

		go helpers.SendLogLocal(c, input, http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, domain.AttachmentResponse{
			Data:        err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	ext, err := mimetype.DetectReader(openedFile)
	if err != nil {
		helpers.LogInit(err.Error())

		go helpers.SendLogLocal(c, input, http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, domain.AttachmentResponse{
			Data:        err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	filename := input.ImageCategory + "_" + input.NoInspec + "_" + time.Now().Format("20060102150405") + ext.Extension()

	// dst := path + `\images\`
	dst := `D:\File_attachment\Inspection_RA\images\`

	input.Filename = filename

	result, err := h.attachmentService.Store(input, dst)
	if err != nil {

		helpers.LogInit(err.Error())

		go helpers.SendLogLocal(c, input, http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, domain.AttachmentResponse{
			Data:        err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	err = c.SaveUploadedFile(file, dst+filename)
	if err != nil {

		helpers.LogInit(err.Error())

		go helpers.SendLogLocal(c, input, http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, domain.AttachmentResponse{
			Data:        err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	go helpers.SendLogLocal(c, input, http.StatusOK, "")

	c.JSON(http.StatusOK, domain.AttachmentResponse{
		Data:        result,
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
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
// @Router /api/v1/attachment/db [post]
// @Tags Attachment
func (h *attachmentHandler) StoreFile(c *gin.Context) {
	start := time.Now()
	var input domain.AttachmentRequest

	c.ShouldBind(&input)

	path, _ := os.Getwd()

	file, _ := c.FormFile("file")

	openedFile, err := file.Open()
	if err != nil {
		helpers.LogInit(err.Error())

		go helpers.SendLogLocal(c, input, http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, domain.AttachmentResponse{
			Data:        err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	ext, err := mimetype.DetectReader(openedFile)
	if err != nil {
		helpers.LogInit(err.Error())

		go helpers.SendLogLocal(c, input, http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, domain.AttachmentResponse{
			Data:        err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	filename := file.Filename + "_" + input.NoInspec + ext.Extension()

	dst := path + `\db\`

	err = c.SaveUploadedFile(file, dst+filename)
	if err != nil {

		helpers.LogInit(err.Error())

		go helpers.SendLogLocal(c, input, http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, domain.AttachmentResponse{
			Data:        err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	go helpers.SendLogLocal(c, input, http.StatusOK, "")

	c.JSON(http.StatusOK, domain.AttachmentResponse{
		Data:        "success upload",
		ElapsedTime: fmt.Sprint(time.Since(start)),
	})
}
