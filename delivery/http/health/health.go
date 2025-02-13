package health

import (
	"antrian-golang/common/response/model"
	"net/http"

	service "antrian-golang/service/health"

	common "antrian-golang/common/registry"

	"github.com/gin-gonic/gin"
)

type IHealth interface {
	Check(c *gin.Context)
}

type Health struct {
	common common.IRegistry
	health service.IHealth
}

func NewHealth(common common.IRegistry, health service.IHealth) *Health {
	return &Health{
		common: common,
		health: health,
	}
}

// Check Health
// @Summary Health check
// @Schemes
// @Description do health check for databases
// @Tags        check
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response
// @Router      /health [get]
func (h *Health) Check(c *gin.Context) {
	const logCtx = "delivery.http.health.Health.Check"

	var (
		ctx     = c.Request.Context()
		status  = http.StatusOK
		message = http.StatusText(status)
	)

	c.JSON(http.StatusOK, model.Response{
		Status:  model.StatusSuccess,
		Data:    h.health.Check(ctx),
		Message: message,
	})
	return
}
