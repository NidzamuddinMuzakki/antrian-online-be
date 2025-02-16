package middleware

import (
	common "antrian-golang/common/registry"
	"antrian-golang/lib/security"
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type IMiddleware interface {
	Auth() gin.HandlerFunc
}

type middleware struct {
	common   common.IRegistry
	JwtUtils security.IJwtToken
}

func NewMiddleware(common common.IRegistry, JwtUtils security.IJwtToken) IMiddleware {
	return &middleware{
		common:   common,
		JwtUtils: JwtUtils,
	}
}

// func (m *middleware) CheckPermission(permissionName []string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx := c.Request.Context()

// 		requestID := ctx.Value(constant.XRequestIdHeader).(string)

// 		token := c.Request.Header["Authorization"]
// 		tokenLogin := ""

// 		if len(token) == 0 || token == nil {
// 			c.JSON(http.StatusUnauthorized, response.Unauthorised(ctx))
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		client := action.Init(
// 			config.Cold.AuthHttpHost,
// 			config.Cold.AuthGrpcHost,
// 		)
// 		spl := strings.Split(token[0], " ")
// 		if spl[0] == "Bearer" {
// 			tokenLogin = spl[1]
// 		} else {
// 			tokenLogin = spl[0]
// 		}

// 		auth, err := client.Authorize(tokenLogin, requestID, permissionName)

// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, response.Unauthorised(ctx))
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		//*Forbidden access*//
// 		if !auth.IsAuthorize {
// 			c.JSON(http.StatusForbidden, response.Forbidden(ctx))
// 			c.AbortWithStatus(http.StatusForbidden)
// 			return
// 		}

// 		me, err := client.Me(tokenLogin, "")

// 		if me == nil {
// 			c.JSON(http.StatusUnauthorized, response.Unauthorised(ctx))
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, response.Unauthorised(ctx))
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		role := model.UserRole{
// 			Name: me.Roles[0],
// 		}
// 		user := model.UserDetail{
// 			UserId: me.UserId,
// 			Name:   me.Name,
// 			Email:  me.Email,
// 			Role:   role,
// 		}

// 		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), constant.XUserDetail, user))
// 	}
// }

// func (m *middleware) CheckRequestId() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		requestID := c.GetHeader("X-Request-Id")
// 		if requestID == "" {
// 			requestID = uuid.New().String()
// 		}
// 		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), constant.XRequestIdHeader, requestID))
// 	}
// }

// func (m *middleware) CheckSignatureKey() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx := c.Request.Context()

// 		serviceName := c.GetHeader(constant.XServiceNameHeader)
// 		requestId := c.GetHeader(constant.XRequestIdHeader)
// 		signature := c.GetHeader(constant.XRequestSignatureHeader)
// 		requestAt := c.GetHeader(constant.XRequestAtHeader)

// 		secretKey := config.Cold.SecretKeyApiCallback

// 		if serviceName == "" || requestId == "" || requestAt == "" || signature == "" {
// 			c.JSON(http.StatusUnauthorized, response.Unauthorised(ctx))
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		key := serviceName + ":" + requestId + ":" + requestAt + ":" + secretKey

// 		valid := m.common.GetSignature().Verify(ctx, key, signature)
// 		if !valid {
// 			c.JSON(http.StatusUnauthorized, response.InvalidSignature(ctx))
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 	}
// // }

func (m *middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		authT := c.GetHeader("Authorization")
		token, err := m.JwtUtils.ExtractToken(ctx, authT)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString, err := m.JwtUtils.ParseToken(ctx, token)
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

		typeToken := tokenString.Claims.(jwt.MapClaims)["sub"].(string)
		if typeToken == "" {

			c.JSON(http.StatusUnauthorized, errors.New("user-id is empty"))
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		c.Request.Header.Set("user-id", typeToken)

	}
}
