package http

import (
	common "antrian-golang/common/registry"
	commonModel "antrian-golang/common/response/model"
	"antrian-golang/payload"
	"net/http"

	"antrian-golang/service"

	"github.com/gin-gonic/gin"
)

type IRoleDelivery interface {
	CreateRole(c *gin.Context)
	UpdateRole(c *gin.Context)

	Activate(c *gin.Context)
	DeActivate(c *gin.Context)
	FindByUuid(c *gin.Context)
	FindAll(c *gin.Context)
}
type RoleDelivery struct {
	common          common.IRegistry
	serviceRegistry service.IRegistry
}

func NewRoleDelivery(common common.IRegistry, serviceRegistry service.IRegistry) IRoleDelivery {
	return &RoleDelivery{
		common:          common,
		serviceRegistry: serviceRegistry,
	}
}

func (h *RoleDelivery) CreateRole(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestCreateRole{}
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
	data, err := h.serviceRegistry.GetRoleService().InsertData(ctx, req)
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

func (h *RoleDelivery) UpdateRole(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestUpdateRole{}
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
	err = h.serviceRegistry.GetRoleService().UpdateData(ctx, req)
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

func (h *RoleDelivery) Activate(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestUpdateRoleStatus{}
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
	err = h.serviceRegistry.GetRoleService().Activate(ctx, req)
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

func (h *RoleDelivery) DeActivate(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestUpdateRoleStatus{}
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
	err = h.serviceRegistry.GetRoleService().DeActivate(ctx, req)
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

func (h *RoleDelivery) FindByUuid(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestRoleFindByUUID{}
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
	data, err := h.serviceRegistry.GetRoleService().FindById(ctx, req)
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

func (h *RoleDelivery) FindAll(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx    = c.Request.Context()
		req    = payload.RequestGetRole{}
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
	data, count, err := h.serviceRegistry.GetRoleService().FindAll(ctx, req)
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
