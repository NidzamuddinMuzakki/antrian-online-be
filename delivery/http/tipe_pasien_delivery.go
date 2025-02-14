package http

import (
	common "antrian-golang/common/registry"
	commonModel "antrian-golang/common/response/model"
	"antrian-golang/payload"

	"antrian-golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ITipePasienDelivery interface {
	FindAll(c *gin.Context)
	FindAllExternal(c *gin.Context)
	FindById(c *gin.Context)
	FindByIdExternal(c *gin.Context)
	InsertData(c *gin.Context)
	UpdateData(c *gin.Context)
	Activate(c *gin.Context)
	DeActivate(c *gin.Context)
}
type TipePasienDelivery struct {
	common          common.IRegistry
	serviceRegistry service.IRegistry
}

func NewTipePasienDelivery(common common.IRegistry, serviceRegistry service.IRegistry) ITipePasienDelivery {
	return &TipePasienDelivery{
		common:          common,
		serviceRegistry: serviceRegistry,
	}
}

func (h *TipePasienDelivery) FindAll(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx = c.Request.Context()
		req = payload.RequestGetTipePasien{}
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}
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
	data, count, err := h.serviceRegistry.GetTipePasienService().FindAll(ctx, req)
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

func (h *TipePasienDelivery) FindAllExternal(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx = c.Request.Context()
	)

	data, err := h.serviceRegistry.GetTipePasienService().FindAllExternal(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commonModel.Response{
		Status:  commonModel.StatusSuccess,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	})
}

func (h *TipePasienDelivery) FindById(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx = c.Request.Context()
		req = payload.RequestGetTipePasienById{}
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	data, err := h.serviceRegistry.GetTipePasienService().FindById(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commonModel.Response{
		Status:  commonModel.StatusSuccess,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	})
}

func (h *TipePasienDelivery) FindByIdExternal(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx = c.Request.Context()
		req = payload.RequestGetTipePasienById{}
	)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	data, err := h.serviceRegistry.GetTipePasienService().FindByIdExternal(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commonModel.Response{
		Status:  commonModel.StatusSuccess,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	})
}

func (h *TipePasienDelivery) InsertData(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx     = c.Request.Context()
		req     = payload.RequestInsertTipePasien{}
		user_id = c.GetHeader("user-id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	req.UserId = user_id

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	data, err := h.serviceRegistry.GetTipePasienService().InsertData(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commonModel.Response{
		Status:  commonModel.StatusSuccess,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	})
}

func (h *TipePasienDelivery) UpdateData(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx     = c.Request.Context()
		req     = payload.RequestUpdateTipePasien{}
		user_id = c.GetHeader("user-id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	req.UserId = user_id

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	err = h.serviceRegistry.GetTipePasienService().UpdateData(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commonModel.Response{
		Status:  commonModel.StatusSuccess,
		Message: http.StatusText(http.StatusOK),
	})
}

func (h *TipePasienDelivery) Activate(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx     = c.Request.Context()
		req     = payload.RequestUpdateStatus{}
		user_id = c.GetHeader("user-id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	req.UserId = user_id

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	err = h.serviceRegistry.GetTipePasienService().Activate(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commonModel.Response{
		Status:  commonModel.StatusSuccess,
		Message: http.StatusText(http.StatusOK),
	})
}

func (h *TipePasienDelivery) DeActivate(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx     = c.Request.Context()
		req     = payload.RequestUpdateStatus{}
		user_id = c.GetHeader("user-id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	req.UserId = user_id

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	err = h.serviceRegistry.GetTipePasienService().DeActivate(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commonModel.Response{
		Status:  commonModel.StatusSuccess,
		Message: http.StatusText(http.StatusOK),
	})
}
