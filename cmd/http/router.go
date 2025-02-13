package http

import (
	"net/http"

	commonMiddleware "antrian-golang/common/middleware/gin"
	common "antrian-golang/common/registry"
	commonResponse "antrian-golang/common/response"

	delivery "antrian-golang/delivery/http"

	// commonSignature "bitbucket.org/moladinTech/go-lib-common/signature"
	// sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

type Router interface {
	Register() *gin.Engine
	swagger()
}

type router struct {
	engine   *gin.Engine
	common   common.IRegistry
	delivery delivery.IRegistry
}

func NewRouter(
	common common.IRegistry,
	delivery delivery.IRegistry,

) Router {
	return &router{
		engine:   gin.Default(),
		common:   common,
		delivery: delivery,
	}
}

// @title          insurance-cofi Swagger API
// @version        1.0
// @description    insurance-cofi Swagger API
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func (r *router) Register() *gin.Engine {

	// Middleware
	r.engine.Use(
		commonMiddleware.CORS(),
		commonMiddleware.RequestID(),
		r.common.GetPanicRecoveryMiddleware().PanicRecoveryMiddleware(),
	)

	// handle no-route error (404 not found)
	commonResponse.RouteNotFound(r.engine)

	// Landing
	r.engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
	})

	// Health Check
	r.engine.GET("/health", r.delivery.GetHealth().Check)

	// v1
	// Configuration
	// r.swagger()
	r.v1()

	return r.engine
}

func (r *router) swagger() {
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// Route: /docs/index.html
	// r.engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (r *router) v1() {
	// signature, err := commonSignature.NewSignature(commonSignature.WithAlgorithm(commonSignature.Sha256))
	// if err != nil {
	// 	panic(err)
	// }

	// common := common.NewRegistry()
	// middlewareImpl := middleware.NewMiddleware(common)

	// v1 := r.engine.Group("/v1")
	// v1.POST("/helper/register-signature", r.delivery.GetSignature().RegisterSignature)
	// v1.GET("/helper/verify-signature", r.delivery.GetSignature().VerifySignature)
}
