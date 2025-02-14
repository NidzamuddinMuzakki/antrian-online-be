package http

import (
	common "antrian-golang/common/registry"
	commonModel "antrian-golang/common/response/model"
	"antrian-golang/payload"
	"net/http"

	"antrian-golang/service"

	"github.com/gin-gonic/gin"
)

type IUserDelivery interface {
	Login(c *gin.Context)
}
type UserDelivery struct {
	common          common.IRegistry
	serviceRegistry service.IRegistry
}

func NewUserDelivery(common common.IRegistry, serviceRegistry service.IRegistry) IUserDelivery {
	return &UserDelivery{
		common:          common,
		serviceRegistry: serviceRegistry,
	}
}

func (h *UserDelivery) Login(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx = c.Request.Context()
		req = payload.RequestLogin{}
	)

	err := c.ShouldBindJSON(&req)
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
	data, err := h.serviceRegistry.GetUserService().Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, commonModel.Response{
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
