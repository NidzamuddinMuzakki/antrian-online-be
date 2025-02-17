package service

import (
	"antrian-golang/common/constant"
	commonDs "antrian-golang/common/data_source"
	"antrian-golang/common/logger"
	"antrian-golang/common/registry"
	commonTime "antrian-golang/common/time"
	"antrian-golang/lib/security"
	"antrian-golang/model"
	"antrian-golang/payload"
	"antrian-golang/repository"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type IUserService interface {
	Login(ctx context.Context, payload payload.RequestLogin) (string, error)
	CreateUser(ctx context.Context, payload payload.RequestCreateUser) (int, error)
	UpdateUserPassword(ctx context.Context, payload payload.RequestUpdateUserPassword) error
	UpdateData(ctx context.Context, payload payload.RequestUpdateUser) error
	FindById(ctx context.Context, payload payload.RequestUpdateUserFindById) (*model.User, error)
	FindAll(ctx context.Context, payload payload.RequestGetUser) ([]*model.User, uint64, error)
	Activate(ctx context.Context, payload payload.RequestUpdateUserStatus) error
	DeActivate(ctx context.Context, payload payload.RequestUpdateUserStatus) error
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
func (s *UserService) FindAll(ctx context.Context, payload payload.RequestGetUser) ([]*model.User, uint64, error) {
	idR, err := strconv.Atoi(payload.UserId)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, 0, err
	}
	getUserAdmin, err := s.repoRegistry.GetUserRepository().FindById(ctx, idR)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, 0, err
	}
	if getUserAdmin.Role != "849c9eee-e30f-4dc5-9816-9b395b0121b7" {
		err = errors.New("role anda bukan superadmin")
		logger.Error(ctx, err.Error(), err)
		return nil, 0, err
	}
	getData, count, err := s.repoRegistry.GetUserRepository().FindAll(ctx, payload)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, 0, err
	}
	return getData, count, nil
}

func (s *UserService) FindById(ctx context.Context, payload payload.RequestUpdateUserFindById) (*model.User, error) {
	getData, err := s.repoRegistry.GetUserRepository().FindById(ctx, payload.Id)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}
	idR, err := strconv.Atoi(payload.UserId)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}
	getUserAdmin, err := s.repoRegistry.GetUserRepository().FindById(ctx, idR)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}
	if getUserAdmin.Role != "849c9eee-e30f-4dc5-9816-9b395b0121b7" {
		if payload.Id != getUserAdmin.Id {
			err = errors.New("role anda bukan superadmin, dan anda tidak berhak mengubah akun lain")
			logger.Error(ctx, err.Error(), err)
			return nil, err
		}
	}

	return getData, nil
}
func (s *UserService) UpdateUserPassword(ctx context.Context, payload payload.RequestUpdateUserPassword) error {
	getUser, err := s.repoRegistry.GetUserRepository().FindById(ctx, payload.Id)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	if getUser.Status == constant.StatusActive {
		err = errors.New("status already " + constant.StatusActive)
		logger.Error(ctx, err.Error(), err)
		return err
	}
	idR, err := strconv.Atoi(payload.UserId)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	getUserAdmin, err := s.repoRegistry.GetUserRepository().FindById(ctx, idR)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	if getUserAdmin.Role != "849c9eee-e30f-4dc5-9816-9b395b0121b7" {
		if getUser.Id != getUserAdmin.Id {
			err = errors.New("role anda bukan superadmin, dan anda tidak berhak mengubah akun lain")
			logger.Error(ctx, err.Error(), err)
			return err
		}
	}

	salt, err := security.GenerateSalt(10)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	password, err := security.HashPassword(payload.Password, salt)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}

	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)

		ds := model.User{
			Id:        payload.Id,
			Password:  password,
			UpdatedBy: payload.UserId,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetUserRepository().UpdatePassword(ctx, tx, ds)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		return nil
	})

	err = s.repoRegistry.GetUtilTx().WithTx(ctx, doFunc, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}

	return nil
}

func (s *UserService) CreateUser(ctx context.Context, payload payload.RequestCreateUser) (int, error) {
	getRole, err := s.repoRegistry.GetRoleRepository().FindByUuid(ctx, payload.Role)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}
	idR, err := strconv.Atoi(payload.UserId)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}
	getUserAdmin, err := s.repoRegistry.GetUserRepository().FindById(ctx, idR)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}
	if getUserAdmin.Role != "849c9eee-e30f-4dc5-9816-9b395b0121b7" {
		err = errors.New("role anda bukan superadmin")
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}

	ids := 0
	salt, err := security.GenerateSalt(10)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}
	password, err := security.HashPassword(payload.Password, salt)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}
	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)

		ds := model.User{
			Username:  payload.Username,
			Password:  password,
			Role:      getRole.RoleUuid,
			CreatedBy: payload.UserId,
			CreatedAt: now,
		}
		ids, err = s.repoRegistry.GetUserRepository().InsertData(ctx, tx, ds)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		return nil
	})

	err = s.repoRegistry.GetUtilTx().WithTx(ctx, doFunc, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}

	return ids, nil

}
func (s *UserService) UpdateData(ctx context.Context, payload payload.RequestUpdateUser) error {
	getUser, err := s.repoRegistry.GetUserRepository().FindById(ctx, payload.Id)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	if getUser.Status == constant.StatusActive {
		err = errors.New("status already " + constant.StatusActive)
		logger.Error(ctx, err.Error(), err)
		return err
	}
	idR, err := strconv.Atoi(payload.UserId)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	getRole, err := s.repoRegistry.GetRoleRepository().FindByUuid(ctx, payload.Role)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	getUserAdmin, err := s.repoRegistry.GetUserRepository().FindById(ctx, idR)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	if getUserAdmin.Role != "849c9eee-e30f-4dc5-9816-9b395b0121b7" {
		err = errors.New("role anda bukan superadmin")
		logger.Error(ctx, err.Error(), err)
		return err
	}

	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)

		ds := model.User{
			Id:        payload.Id,
			UpdatedBy: payload.UserId,
			Username:  payload.Username,
			Role:      getRole.RoleUuid,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetUserRepository().UpdateData(ctx, tx, ds)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		return nil
	})

	err = s.repoRegistry.GetUtilTx().WithTx(ctx, doFunc, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}

	return nil

}
func (s *UserService) Activate(ctx context.Context, payload payload.RequestUpdateUserStatus) error {
	getUser, err := s.repoRegistry.GetUserRepository().FindById(ctx, payload.Id)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	if getUser.Status == constant.StatusActive {
		err = errors.New("status already " + constant.StatusActive)
		logger.Error(ctx, err.Error(), err)
		return err
	}
	idR, err := strconv.Atoi(payload.UserId)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	getUserAdmin, err := s.repoRegistry.GetUserRepository().FindById(ctx, idR)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	if getUserAdmin.Role != "849c9eee-e30f-4dc5-9816-9b395b0121b7" {
		err = errors.New("role anda bukan superadmin")
		logger.Error(ctx, err.Error(), err)
		return err
	}

	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)

		ds := model.User{
			Id:        payload.Id,
			UpdatedBy: payload.UserId,
			Status:    constant.StatusActive,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetUserRepository().UpdateStatus(ctx, tx, ds)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		return nil
	})

	err = s.repoRegistry.GetUtilTx().WithTx(ctx, doFunc, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}

	return nil

}

func (s *UserService) DeActivate(ctx context.Context, payload payload.RequestUpdateUserStatus) error {
	getUser, err := s.repoRegistry.GetUserRepository().FindById(ctx, payload.Id)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	if getUser.Status == constant.StatusNotActive {
		err = errors.New("status already " + constant.StatusNotActive)
		logger.Error(ctx, err.Error(), err)
		return err
	}
	idR, err := strconv.Atoi(payload.UserId)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	getUserAdmin, err := s.repoRegistry.GetUserRepository().FindById(ctx, idR)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	if getUserAdmin.Role != "849c9eee-e30f-4dc5-9816-9b395b0121b7" {
		err = errors.New("role anda bukan superadmin")
		logger.Error(ctx, err.Error(), err)
		return err
	}

	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)

		ds := model.User{
			Id:        payload.Id,
			UpdatedBy: payload.UserId,
			Status:    constant.StatusNotActive,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetUserRepository().UpdateStatus(ctx, tx, ds)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		return nil
	})

	err = s.repoRegistry.GetUtilTx().WithTx(ctx, doFunc, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}

	return nil

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
	// _, err := s.repoRegistry.GetRoleRepository().FindByUuid(ctx, data.Role)
	// if err != nil {
	// 	logger.Error(ctx, err.Error(), err)
	// 	return "", err
	// }
	if data.Status != constant.StatusActive {
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
