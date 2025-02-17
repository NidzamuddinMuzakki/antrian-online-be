package http

import (
	"antrian-golang/common/logger"
	common "antrian-golang/common/registry"
	commonModel "antrian-golang/common/response/model"
	"antrian-golang/lib/security"
	"antrian-golang/model"
	"antrian-golang/payload"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"antrian-golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type LoketPanggil struct {
	Type         string      `json:"type"`
	TypePasienId int         `json:"tipe_pasien_id"`
	Body         interface{} `json:"body"`
}
type ILoketDelivery interface {
	FindAll(c *gin.Context)
	FindAllExternal(c *gin.Context)
	FindById(c *gin.Context)
	FindByIdExternal(c *gin.Context)
	InsertData(c *gin.Context)
	UpdateData(c *gin.Context)
	Activate(c *gin.Context)
	DeActivate(c *gin.Context)
	UpdateUserId(c *gin.Context)
}
type LoketDelivery struct {
	common          common.IRegistry
	serviceRegistry service.IRegistry
	JwtUtils        security.IJwtToken
}

func NewLoketDelivery(common common.IRegistry, serviceRegistry service.IRegistry, JwtUtils security.IJwtToken) ILoketDelivery {
	return &LoketDelivery{
		common:          common,
		serviceRegistry: serviceRegistry,
		JwtUtils:        JwtUtils,
	}
}

func (h *LoketDelivery) FindAll(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx = c.Request.Context()
		req = payload.RequestGetLoket{}
	)

	c.GetHeader("user-id")
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
	data, count, err := h.serviceRegistry.GetLoketService().FindAll(ctx, req)
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

func (h *LoketDelivery) FindAllExternal(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx = c.Request.Context()
	)

	data, err := h.serviceRegistry.GetLoketService().FindAllExternal(ctx)
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

func (h *LoketDelivery) FindById(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx = c.Request.Context()
		req = payload.RequestGetLoketById{}
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
	data, err := h.serviceRegistry.GetLoketService().FindById(ctx, req)
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

func (h *LoketDelivery) FindByIdExternal(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx = c.Request.Context()
		req = payload.RequestGetLoketById{}
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
	data, err := h.serviceRegistry.GetLoketService().FindByIdExternal(ctx, req)
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

func (h *LoketDelivery) InsertData(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx     = c.Request.Context()
		req     = payload.RequestInsertLoket{}
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
	data, err := h.serviceRegistry.GetLoketService().InsertData(ctx, req)
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

func (h *LoketDelivery) UpdateData(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx     = c.Request.Context()
		req     = payload.RequestUpdateLoket{}
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
	err = h.serviceRegistry.GetLoketService().UpdateData(ctx, req)
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

var clients = make(map[*websocket.Conn]bool) // Connected clients

func (h *LoketDelivery) UpdateUserId(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	loketId := 0
	kkk := ""
	typeToken := ""

	if err != nil {
		fmt.Println(err, "nidzam3")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	clients[conn] = true
	ctx := c.Request.Context()
	defer func(idLoket *int, typeToken *string, ctx context.Context) {

		// fmt.Println("Recovered. Error:\n", *errs)
		clients[conn] = false
		defer conn.Close()
		if typeToken != nil && *typeToken != "" {

			mmm := payload.RequestUpdateLoketUserId{
				Id:      *idLoket,
				UserId:  "0",
				UserId2: *typeToken,
			}
			err = h.serviceRegistry.GetLoketService().UserIdLoket(ctx, mmm)
			fmt.Println("nidzam4", *idLoket, *typeToken, err)
			if err != nil {
				if err.Error() == "nothing update" {
					err = errors.New("loket sudah di pakai")
				}
				c.JSON(http.StatusBadRequest, err)
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
		}

	}(&loketId, &typeToken, ctx)
	if len(c.Request.URL.Query()["token"]) > 0 {

		kkk = c.Request.URL.Query()["token"][0]
	}
	IdLoket := ""
	if len(c.Request.URL.Query()["loket_id"]) > 0 {
		IdLoket = c.Request.URL.Query()["loket_id"][0]

	}

	if IdLoket != "" {

		tokenString, err := h.JwtUtils.ParseToken(ctx, kkk)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		_, isOK := tokenString.Claims.(jwt.MapClaims)
		if !isOK {
			c.JSON(http.StatusUnauthorized, errors.New("Unable to parse claims"))
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		typeToken = tokenString.Claims.(jwt.MapClaims)["sub"].(string)
		fmt.Println(typeToken, "nidzam-ganteng")
		if typeToken == "" {

			c.JSON(http.StatusUnauthorized, errors.New("user-id is empty"))
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}
		loketId, err = strconv.Atoi(IdLoket)
		if err != nil {

			c.JSON(http.StatusBadRequest, err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	mmm := payload.RequestUpdateLoketUserId{
		Id:      loketId,
		UserId:  typeToken,
		UserId2: typeToken,
	}
	for {
		datas := LoketPanggil{}

		err := conn.ReadJSON(&datas)
		if err != nil {
			fmt.Println(err, "nidzam")
			continue
		}
		if datas.Type == "register" {
			err = h.serviceRegistry.GetLoketService().UserIdLoket(ctx, mmm)
			fmt.Println("nidzam2", err)

			if err != nil {
				if err.Error() == "nothing update" {
					err = errors.New("loket sudah di pakai")
				}
				c.JSON(http.StatusBadRequest, err)
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			dat := model.User{
				Id: 2,
			}
			m, _ := json.Marshal(dat)
			conn.WriteMessage(websocket.TextMessage, m)

		} else if datas.Type == "call" {
			aa, _ := json.Marshal(datas.Body)
			req := model.Antrian{}
			json.Unmarshal(aa, &req)
			req.UpdatedBy = typeToken
			req.Status = "call"
			err := h.serviceRegistry.GetAntrianService().UpdateData(ctx, req)
			if err != nil {
				res := commonModel.Response{
					Status:  commonModel.StatusError,
					Message: err.Error(),
				}
				rs2, _ := json.Marshal(res)
				conn.WriteMessage(websocket.TextMessage, rs2)
				logger.Error(ctx, err.Error(), err)
				c.JSON(http.StatusBadRequest, res)
				continue
			}
			res := commonModel.Response{
				Type:    "panggil",
				Status:  commonModel.StatusSuccess,
				Message: http.StatusText(http.StatusOK),
			}
			mm, _ := json.Marshal(res)
			for client := range clients {
				client.WriteMessage(websocket.TextMessage, mm)

			}

			c.JSON(http.StatusOK, res)

		} else if datas.Type == "selesai" {
			aa, _ := json.Marshal(datas.Body)
			req := model.Antrian{}
			json.Unmarshal(aa, &req)
			req.UpdatedBy = typeToken
			req.Status = "selesai"
			err := h.serviceRegistry.GetAntrianService().UpdateData(ctx, req)
			if err != nil {
				res := commonModel.Response{
					Status:  commonModel.StatusError,
					Message: err.Error(),
				}
				rs2, _ := json.Marshal(res)
				conn.WriteMessage(websocket.TextMessage, rs2)
				logger.Error(ctx, err.Error(), err)
				c.JSON(http.StatusBadRequest, res)
				continue
			}
			res := commonModel.Response{
				Type:    "selesai",
				Status:  commonModel.StatusSuccess,
				Message: http.StatusText(http.StatusOK),
			}
			mm, _ := json.Marshal(res)
			for client := range clients {
				client.WriteMessage(websocket.TextMessage, mm)

			}

			c.JSON(http.StatusOK, res)

		} else if datas.Type == "recall" {
			fmt.Println(datas.TypePasienId)

		} else if datas.Type == "insert-antrian" {
			if datas.Body != nil {
				aa, _ := json.Marshal(datas.Body)
				req := payload.AntrianPayloadInsert{}
				json.Unmarshal(aa, &req)
				data, err := h.serviceRegistry.GetAntrianService().InsertData(ctx, req)
				if err != nil {
					res := commonModel.Response{
						Status:  commonModel.StatusError,
						Message: err.Error(),
					}
					rs2, _ := json.Marshal(res)
					conn.WriteMessage(websocket.TextMessage, rs2)
					logger.Error(ctx, err.Error(), err)
					c.JSON(http.StatusBadRequest, res)
					continue
				}
				kk := commonModel.Response{
					Type:    "insert-antrian",
					Status:  commonModel.StatusSuccess,
					Message: http.StatusText(http.StatusOK),
					Data:    data,
				}
				k, _ := json.Marshal(kk)
				for client := range clients {
					client.WriteMessage(websocket.TextMessage, k)

				}
				c.JSON(http.StatusOK, kk)

			}
			continue
		} else if datas.Type == "list-antrian" {

		}

		// conn.WriteMessage(websocket.TextMessage, m)
		// time.Sleep(time.Second)
	}
}

func (h *LoketDelivery) Activate(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx     = c.Request.Context()
		req     = payload.RequestUpdateLoketStatus{}
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
	err = h.serviceRegistry.GetLoketService().Activate(ctx, req)
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

func (h *LoketDelivery) DeActivate(c *gin.Context) {
	const logCtx = "delivery.http.tnc.GetTncVersion"
	var (
		ctx     = c.Request.Context()
		req     = payload.RequestUpdateLoketStatus{}
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
	err = h.serviceRegistry.GetLoketService().DeActivate(ctx, req)
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
