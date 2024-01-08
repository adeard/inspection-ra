package objpart

import (
	"fmt"
	"inspection-ra/domain"
	"inspection-ra/helpers"
	"inspection-ra/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type objPartHandler struct {
	objPartService Service
}

func NewObjPartHandler(v1 *gin.RouterGroup, objPartService Service) {

	handler := &objPartHandler{objPartService}

	objpart := v1.Group("objpart")

	objpart.GET("", middlewares.AuthService(), handler.GetAll)
}

// @Summary Get Object Part
// @Description Get Object Part
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {object} domain.ObjPartResponse{data=domain.ObjPartDataResponse}
// @Router /api/v1/objpart [get]
// @Tags Object Part
func (h *objPartHandler) GetAll(c *gin.Context) {
	start := time.Now()
	objPartDataResp := []domain.ObjPartDataResponse{}
	objPart, err := h.objPartService.GetAll()
	if err != nil {

		helpers.LogInit(err.Error())

		c.JSON(http.StatusBadRequest, domain.ObjPartResponse{
			Message:     err.Error(),
			ElapsedTime: fmt.Sprint(time.Since(start)),
		})

		return
	}

	for _, v := range objPart {
		objPartDataResp = append(objPartDataResp, domain.ObjPartDataResponse{
			Id:              v.Id,
			ObjPartRequest:  v.ObjPartRequest,
			VehicleTypeCode: v.VehicleTypeData.Code,
		})
	}

	c.JSON(http.StatusOK, domain.ObjPartResponse{
		Data:        objPartDataResp,
		ElapsedTime: fmt.Sprint(time.Since(start))},
	)
}
