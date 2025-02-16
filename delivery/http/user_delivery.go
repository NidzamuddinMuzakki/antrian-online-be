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
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	UpdateUserPassword(c *gin.Context)
	Activate(c *gin.Context)
	DeActivate(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
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

func (h *UserDelivery) CreateUser(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestCreateUser{}
		userid = c.GetHeader("user-id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	req.UserId = userid

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	data, err := h.serviceRegistry.GetUserService().CreateUser(ctx, req)
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

func (h *UserDelivery) UpdateUser(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestUpdateUser{}
		userid = c.GetHeader("user-id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	req.UserId = userid

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	err = h.serviceRegistry.GetUserService().UpdateData(ctx, req)
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
	})
}

func (h *UserDelivery) UpdateUserPassword(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestUpdateUserPassword{}
		userid = c.GetHeader("user-id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	req.UserId = userid

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	err = h.serviceRegistry.GetUserService().UpdateUserPassword(ctx, req)
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
	})
}

func (h *UserDelivery) Activate(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestUpdateUserStatus{}
		userid = c.GetHeader("user-id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	req.UserId = userid

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	err = h.serviceRegistry.GetUserService().Activate(ctx, req)
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
	})
}

func (h *UserDelivery) DeActivate(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestUpdateUserStatus{}
		userid = c.GetHeader("user-id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	req.UserId = userid

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	err = h.serviceRegistry.GetUserService().DeActivate(ctx, req)
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
	})
}

func (h *UserDelivery) FindById(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestUpdateUserFindById{}
		userid = c.GetHeader("user-id")
	)

	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	req.UserId = userid

	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	data, err := h.serviceRegistry.GetUserService().FindById(ctx, req)
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

func (h *UserDelivery) FindAll(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestGetUser{}
		userid = c.GetHeader("user-id")
	)

	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
	}

	req.UserId = userid
	if req.Page < 1 {
		req.Page = 10
	}
	if req.RowPerpage < 1 {
		req.RowPerpage = 10
	}
	err = h.common.GetValidator().Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}
	data, count, err := h.serviceRegistry.GetUserService().FindAll(ctx, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, commonModel.Response{
			Status:  commonModel.StatusError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commonModel.Response{
		Status:       commonModel.StatusSuccess,
		Message:      http.StatusText(http.StatusOK),
		Data:         data,
		CurrentPage:  req.Page,
		RowPerpage:   req.RowPerpage,
		TotalRecords: count,
	})
}
