package http

import (
	"antrian-golang/common/logger"
	common "antrian-golang/common/registry"
	commonModel "antrian-golang/common/response/model"
	"antrian-golang/lib/security"
	"antrian-golang/payload"
	"net/http"
	"strconv"

	"antrian-golang/service"

	"github.com/gin-gonic/gin"
)

type IAntrianDelivery interface {
	InsertData(c *gin.Context)
	FindAll(c *gin.Context)
}
type AntrianDelivery struct {
	common          common.IRegistry
	serviceRegistry service.IRegistry
	JwtUtils        security.IJwtToken
}

func NewAntrianDelivery(common common.IRegistry, serviceRegistry service.IRegistry, JwtUtils security.IJwtToken) IAntrianDelivery {
	return &AntrianDelivery{
		common:          common,
		serviceRegistry: serviceRegistry,
		JwtUtils:        JwtUtils,
	}
}

func (h *AntrianDelivery) InsertData(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req = payload.AntrianPayloadInsert{}
	)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	const logCtx = "delivery.http.tnc.GetTncVersion"

	tipePasienIdS := c.Request.URL.Query()["tipe_pasien_id"][0]
	tipePasienId, err := strconv.Atoi(tipePasienIdS)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.TipePasienId = tipePasienId

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}

	for {
		datas := LoketPanggil{}

		err := conn.ReadJSON(&datas)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			continue
		}

	}

}

func (h *AntrianDelivery) FindAll(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx = c.Request.Context()
		req = payload.RequestGetAntrian{}
	)

	user_id := c.GetHeader("user-id")
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}
	a, _ := strconv.Atoi(user_id)
	req.UserId = a
	if req.Page < 1 {
		req.Page = 1
	}
	if req.RowPerpage < 1 {
		req.RowPerpage = 1
	}

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	data, count, err := h.serviceRegistry.GetAntrianService().FindAll(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commonModel.Response{
		Status:       commonModel.StatusSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data:         data,
		TotalRecords: count,
		CurrentPage:  req.Page,
		RowPerpage:   req.RowPerpage,
	})
}
