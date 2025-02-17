package service

import (
	"antrian-golang/common/constant"
	commonDs "antrian-golang/common/data_source"
	"antrian-golang/common/logger"
	"antrian-golang/common/registry"
	commonTime "antrian-golang/common/time"
	"antrian-golang/model"
	"antrian-golang/payload"
	"antrian-golang/repository"
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IRoleService interface {
	FindAll(ctx context.Context, payload payload.RequestGetRole) ([]*model.Role, uint64, error)
	FindById(ctx context.Context, payload payload.RequestRoleFindByUUID) (*model.Role, error)
	InsertData(ctx context.Context, payload payload.RequestCreateRole) (string, error)
	UpdateData(ctx context.Context, payload payload.RequestUpdateRole) error
	Activate(ctx context.Context, payload payload.RequestUpdateRoleStatus) error
	DeActivate(ctx context.Context, payload payload.RequestUpdateRoleStatus) error
}

type RoleService struct {
	common       registry.IRegistry
	repoRegistry repository.IRegistry
}

func NewRoleService(common registry.IRegistry, repoRegistry repository.IRegistry) IRoleService {
	return &RoleService{
		common:       common,
		repoRegistry: repoRegistry,
	}
}

func (s *RoleService) FindAll(ctx context.Context, payload payload.RequestGetRole) ([]*model.Role, uint64, error) {
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
	getData, count, err := s.repoRegistry.GetRoleRepository().FindAll(ctx, payload)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, 0, err
	}
	return getData, count, nil
}

func (s *RoleService) FindById(ctx context.Context, payload payload.RequestRoleFindByUUID) (*model.Role, error) {
	getData, err := s.repoRegistry.GetRoleRepository().FindByUuid(ctx, payload.Uuid)
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

		err = errors.New("role anda bukan superadmin")
		logger.Error(ctx, err.Error(), err)
		return nil, err

	}

	return getData, nil
}

func (s *RoleService) InsertData(ctx context.Context, payload payload.RequestCreateRole) (string, error) {

	idR, err := strconv.Atoi(payload.UserId)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return "", err
	}
	getUserAdmin, err := s.repoRegistry.GetUserRepository().FindById(ctx, idR)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return "", err
	}
	if getUserAdmin.Role != "849c9eee-e30f-4dc5-9816-9b395b0121b7" {
		err = errors.New("role anda bukan superadmin")
		logger.Error(ctx, err.Error(), err)
		return "", err
	}
	idUUId := uuid.New()
	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)

		ds := model.Role{
			Name:      payload.Name,
			Status:    constant.StatusActive,
			RoleUuid:  idUUId.String(),
			CreatedBy: payload.UserId,
			CreatedAt: now,
		}
		_, err = s.repoRegistry.GetRoleRepository().InsertData(ctx, tx, ds)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		return nil
	})

	err = s.repoRegistry.GetUtilTx().WithTx(ctx, doFunc, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return "", err
	}

	return idUUId.String(), nil

}
func (s *RoleService) UpdateData(ctx context.Context, payload payload.RequestUpdateRole) error {

	idR, err := strconv.Atoi(payload.UserId)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	getRole, err := s.repoRegistry.GetRoleRepository().FindByUuid(ctx, payload.Uuid)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	if getRole.Status != constant.StatusActive {
		err = errors.New("status not active")
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

		ds := model.Role{
			RoleUuid:  payload.Uuid,
			UpdatedBy: payload.UserId,
			Name:      payload.Name,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetRoleRepository().UpdateData(ctx, tx, ds)
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
func (s *RoleService) Activate(ctx context.Context, payload payload.RequestUpdateRoleStatus) error {
	getUser, err := s.repoRegistry.GetRoleRepository().FindByUuid(ctx, payload.Uuid)
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

		ds := model.Role{
			RoleUuid:  payload.Uuid,
			UpdatedBy: payload.UserId,
			Status:    constant.StatusActive,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetRoleRepository().UpdateStatusRole(ctx, tx, ds)
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

func (s *RoleService) DeActivate(ctx context.Context, payload payload.RequestUpdateRoleStatus) error {
	getUser, err := s.repoRegistry.GetRoleRepository().FindByUuid(ctx, payload.Uuid)
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

		ds := model.Role{
			RoleUuid:  payload.Uuid,
			UpdatedBy: payload.UserId,
			Status:    constant.StatusNotActive,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetRoleRepository().UpdateStatusRole(ctx, tx, ds)
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
