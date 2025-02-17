package http

import (
	"crypto/ecdsa"
	"net/http"

	"antrian-golang/cmd/middleware"
	commonMiddleware "antrian-golang/common/middleware/gin"
	common "antrian-golang/common/registry"
	commonResponse "antrian-golang/common/response"
	"antrian-golang/config"
	"antrian-golang/lib/security"

	delivery "antrian-golang/delivery/http"

	// commonSignature "bitbucket.org/moladinTech/go-lib-common/signature"
	// sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
	var err error
	common := common.NewRegistry()

	var privatecdsaKey *ecdsa.PrivateKey
	if privatecdsaKey, err = jwt.ParseECPrivateKeyFromPEM([]byte(config.Cold.SignaturePrivateKey)); err != nil {
		panic(err)
	}

	var publicecdsaKey *ecdsa.PublicKey
	if publicecdsaKey, err = jwt.ParseECPublicKeyFromPEM([]byte(config.Cold.SignaturePublicKey)); err != nil {
		panic(err)
	}

	sec := security.NewJwtUtils(privatecdsaKey, publicecdsaKey)
	middlewareImpl := middleware.NewMiddleware(common, sec)

	v1 := r.engine.Group("/antrian/v1")

	tipePasienexternal := v1.Group("/tipe_pasien")
	tipePasienexternal.GET("/list", r.delivery.GetTipePasienDelivery().FindAllExternal)
	tipePasienexternal.GET("/detail/:id", r.delivery.GetTipePasienDelivery().FindByIdExternal)

	//antrian
	antrianexternal := v1.Group("/antrian")
	antrianexternal.GET("/list", middlewareImpl.Auth(), r.delivery.GetAntrianDelivery().FindAll)

	antrianexternal.GET("/insert-antrian", r.delivery.GetAntrianDelivery().InsertData)

	//loket
	loketexternal := v1.Group("/loket")
	loketexternal.GET("/list", r.delivery.GetLoketDelivery().FindAllExternal)
	loketexternal.GET("/detail/:id", r.delivery.GetLoketDelivery().FindByIdExternal)
	loketexternal.GET("/user-id", r.delivery.GetLoketDelivery().UpdateUserId)

	//admin
	admin := v1.Group("/admin")

	user := admin.Group("/user")

	user.POST("/login", r.delivery.GetUserDelivery().Login)
	user.POST("/", middlewareImpl.Auth(), r.delivery.GetUserDelivery().CreateUser)
	user.PUT("/", middlewareImpl.Auth(), r.delivery.GetUserDelivery().UpdateUser)
	user.POST("/change-password", middlewareImpl.Auth(), r.delivery.GetUserDelivery().UpdateUserPassword)
	user.POST("/activate", middlewareImpl.Auth(), r.delivery.GetUserDelivery().Activate)
	user.POST("/deactivate", middlewareImpl.Auth(), r.delivery.GetUserDelivery().DeActivate)
	user.GET("/list", middlewareImpl.Auth(), r.delivery.GetUserDelivery().FindAll)
	user.GET("/detail/:id", middlewareImpl.Auth(), r.delivery.GetUserDelivery().FindById)

	role := admin.Group("/role")
	role.POST("/", middlewareImpl.Auth(), r.delivery.GetRoleDelivery().CreateRole)
	role.PUT("/", middlewareImpl.Auth(), r.delivery.GetRoleDelivery().UpdateRole)

	role.POST("/activate", middlewareImpl.Auth(), r.delivery.GetRoleDelivery().Activate)
	role.POST("/deactivate", middlewareImpl.Auth(), r.delivery.GetRoleDelivery().DeActivate)
	role.GET("/list", middlewareImpl.Auth(), r.delivery.GetRoleDelivery().FindAll)
	role.GET("/detail/:uuid", middlewareImpl.Auth(), r.delivery.GetRoleDelivery().FindByUuid)

	// tipe pasien
	tipePasien := admin.Group("/tipe_pasien", middlewareImpl.Auth())
	tipePasien.GET("/list", r.delivery.GetTipePasienDelivery().FindAll)
	tipePasien.GET("/detail/:id", r.delivery.GetTipePasienDelivery().FindById)
	tipePasien.POST("/activate", r.delivery.GetTipePasienDelivery().Activate)
	tipePasien.POST("/deactivate", r.delivery.GetTipePasienDelivery().DeActivate)
	tipePasien.POST("", r.delivery.GetTipePasienDelivery().InsertData)
	tipePasien.PUT("", r.delivery.GetTipePasienDelivery().UpdateData)

	//
	loket := admin.Group("/loket", middlewareImpl.Auth())
	loket.GET("/list", r.delivery.GetLoketDelivery().FindAll)
	loket.GET("/detail/:id", r.delivery.GetLoketDelivery().FindById)
	loket.POST("/activate", r.delivery.GetLoketDelivery().Activate)
	loket.POST("/deactivate", r.delivery.GetLoketDelivery().DeActivate)
	loket.POST("", r.delivery.GetLoketDelivery().InsertData)
	loket.PUT("", r.delivery.GetLoketDelivery().UpdateData)

}
