package service

import (
	"antrian-golang/common/constant"
	"antrian-golang/common/logger"
	"antrian-golang/common/registry"
	"antrian-golang/lib/security"
	"antrian-golang/model"
	"antrian-golang/payload"
	"antrian-golang/repository"
	"context"
	"errors"
	"fmt"
)

type IUserService interface {
	Login(ctx context.Context, payload payload.RequestLogin) (string, error)
}

type UserService struct {
	common       registry.IRegistry
	repoRegistry repository.IRegistry
	JwtUtils     security.IJwtToken
}

func NewUserService(common registry.IRegistry, repoRegistry repository.IRegistry, JwtUtils security.IJwtToken) IUserService {
	return &UserService{
		common:       common,
		repoRegistry: repoRegistry,
		JwtUtils:     JwtUtils,
	}
}

func (s *UserService) Login(ctx context.Context, payload payload.RequestLogin) (string, error) {

	m := model.User{
		Username: payload.Username,
	}
	data, err := s.repoRegistry.GetUserRepository().FindUsernameLogin(ctx, m)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return "", err
	}
	if data.Status != constant.StatusActive || data.Role != "superadmin" {
		err = errors.New("status " + data.Status + ", role " + data.Role)
		logger.Error(ctx, err.Error(), err)
		return "", err
	}

	salt := data.Password[:20]
	err = security.ComparePassword(data.Password[21:], payload.Password, salt)
	if err != nil {
		err = errors.New("Username Or Password does not match")
		logger.Error(ctx, err.Error(), err)
		return "", err
	}
	payloadToken := security.JWT_Payload{
		ID:  data.Id,
		Key: fmt.Sprintf("%s:%d:", constant.JWTPrefix, data.Id),
	}
	token, err := s.JwtUtils.GenerateToken(ctx, payloadToken, "")
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return "", err
	}
	return token.AccessToken, nil
}
