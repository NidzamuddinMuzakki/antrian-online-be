package security

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"strings"
	"time"

	"antrian-golang/config"

	"antrian-golang/common/constant"

	"antrian-golang/common/logger"

	"github.com/golang-jwt/jwt"
)

var (
// APPLICATION_NAME = config.Cold.AppName

// JWT_SIGNING_METHOD = jwt.SigningMethodES256

)

type Claims struct {
	jwt.StandardClaims
	JWTPayload
}

type JWTPayload struct {
	Key string `json:"key,omitempty"`
}

type JWT_Payload struct {
	Key string `json:"key,omitempty"`
	ID  int    `json:"id,omitempty"`
}

type ResponseToken struct {
	AccessToken string `json:"access_token"`
}

type JwtToken struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

type IJwtToken interface {
	ExtractToken(ctx context.Context, header string) (token string, err error)
	ParseToken(ctx context.Context, token string) (res *jwt.Token, err error)
	GenerateToken(ctx context.Context, data JWT_Payload, refreshToken string) (res ResponseToken, err error)
}

func NewJwtUtils(PrivateKey *ecdsa.PrivateKey, PublicKey *ecdsa.PublicKey) IJwtToken {

	return &JwtToken{
		PrivateKey: PrivateKey,
		PublicKey:  PublicKey,
	}
}

func (lib *JwtToken) GenerateToken(ctx context.Context, data JWT_Payload, refreshToken string) (res ResponseToken, err error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.Cold.DurationJwt) * time.Hour).Unix(),
			Issuer:    constant.JWTPrefix,
			Subject:   fmt.Sprintf("%d", data.ID),
		},
		JWTPayload: JWTPayload{
			Key: data.Key + constant.JWTKeyCookieAccessToken,
		},
	}

	// Get the encoded bytes
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["typ"] = constant.JWTKeyCookieAccessToken
	accessToken, err := token.SignedString(lib.PrivateKey)
	if err != nil {
		logger.Error(ctx, "Error: %v", err)
		return res, err
	}

	return ResponseToken{
		AccessToken: accessToken,
	}, nil
}

func (lib *JwtToken) ExtractToken(c context.Context, header string) (string, error) {
	// cookie, err := c.Cookie(common.JWTKeyCookieAccessToken)
	// if err != nil {
	// 	lib.Logger.Error(traceID, "Checking Auth", err)
	// 	return "", err
	// }

	// return cookie, nil

	if header == "" {
		errrr := errors.New("Error: Token is empty")
		logger.Error(c, "Checking Auth", errrr)
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, "Bearer ")
	if len(jwtToken) != 2 {
		err := errors.New("Error: Token is invalid")
		logger.Error(c, "incorrect formatted header", err)
		return "", errors.New("incorrect formatted header")
	}

	return jwtToken[1], nil
}

func (lib *JwtToken) ParseToken(ctx context.Context, tokens string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		if method, isOk := token.Method.(*jwt.SigningMethodECDSA); !isOk || method != jwt.SigningMethodES256 {
			err := errors.New("Error: Token is invalid")
			logger.Error(ctx, "bad signed method received", err)
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return lib.PublicKey, nil
	})

	if err != nil {
		err := errors.New("bad token received")
		logger.Error(ctx, "bad token received", err)
		return nil, err
	}
	return parsedToken, nil
}
