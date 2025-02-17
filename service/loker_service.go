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
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type ILoketService interface {
	FindAll(ctx context.Context, payload payload.RequestGetLoket) ([]*model.Loket, uint64, error)
	FindAllExternal(ctx context.Context) ([]*model.LoketAntrian2, error)
	FindById(ctx context.Context, payload payload.RequestGetLoketById) (*model.Loket, error)
	FindByIdExternal(ctx context.Context, payload payload.RequestGetLoketById) (*model.Loket, error)
	InsertData(ctx context.Context, payload payload.RequestInsertLoket) (int, error)
	UpdateData(ctx context.Context, payload payload.RequestUpdateLoket) error
	Activate(ctx context.Context, payload payload.RequestUpdateLoketStatus) error
	DeActivate(ctx context.Context, payload payload.RequestUpdateLoketStatus) error
	UserIdLoket(ctx context.Context, payload payload.RequestUpdateLoketUserId) error
}

type LoketService struct {
	common       registry.IRegistry
	repoRegistry repository.IRegistry
}

func NewLoketService(common registry.IRegistry, repoRegistry repository.IRegistry) ILoketService {
	return &LoketService{
		common:       common,
		repoRegistry: repoRegistry,
	}
}

func (s *LoketService) UpdateData(ctx context.Context, payload payload.RequestUpdateLoket) error {

	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		getData, err := s.repoRegistry.GetLoketRepository().FindById(ctx, payload.Id)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		if getData.Status != constant.StatusActive {
			err = errors.New("status not active, can't update data")
			logger.Error(ctx, err.Error(), err)
			return err
		}
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)
		m := model.Loket{
			Id:        payload.Id,
			Name:      payload.Name,
			UpdatedBy: payload.UserId,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetLoketRepository().Update(ctx, tx, m)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}

		return nil
	})

	err := s.repoRegistry.GetUtilTx().WithTx(ctx, doFunc, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	return nil
}
func (s *LoketService) UserIdLoket(ctx context.Context, payload payload.RequestUpdateLoketUserId) error {

	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {

		_, err := s.repoRegistry.GetLoketRepository().FindById(ctx, payload.Id)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)
		userId, _ := strconv.Atoi(payload.UserId)
		userId2, _ := strconv.Atoi(payload.UserId2)

		m := model.Loket{
			Id:        payload.Id,
			UserId:    userId,
			UpdatedBy: payload.UserId,
			UpdatedAt: &now,
		}
		fmt.Println("nidzam-ganteng", m, userId2)
		err = s.repoRegistry.GetLoketRepository().UpdateUserId(ctx, tx, m, userId2)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}

		return nil
	})

	err := s.repoRegistry.GetUtilTx().WithTx(ctx, doFunc, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	return nil
}

func (s *LoketService) Activate(ctx context.Context, payload payload.RequestUpdateLoketStatus) error {

	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {

		_, err := s.repoRegistry.GetLoketRepository().FindById(ctx, payload.Id)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)
		m := model.Loket{
			Id:        payload.Id,
			Status:    constant.StatusActive,
			UpdatedBy: payload.UserId,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetLoketRepository().UpdateStatus(ctx, tx, m)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}

		return nil
	})

	err := s.repoRegistry.GetUtilTx().WithTx(ctx, doFunc, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	return nil
}

func (s *LoketService) DeActivate(ctx context.Context, payload payload.RequestUpdateLoketStatus) error {

	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		_, err := s.repoRegistry.GetLoketRepository().FindById(ctx, payload.Id)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)
		m := model.Loket{
			Id:        payload.Id,
			Status:    constant.StatusNotActive,
			UpdatedBy: payload.UserId,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetLoketRepository().UpdateStatus(ctx, tx, m)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		return nil
	})

	err := s.repoRegistry.GetUtilTx().WithTx(ctx, doFunc, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	return nil
}
func (s *LoketService) InsertData(ctx context.Context, payload payload.RequestInsertLoket) (int, error) {
	ids := 0
	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)
		m := model.Loket{
			Name:      payload.Name,
			Status:    constant.StatusActive,
			CreatedBy: payload.UserId,
			CreatedAt: now,
		}
		idR, err := s.repoRegistry.GetLoketRepository().Insert(ctx, tx, m)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		ids = idR
		return nil
	})

	err := s.repoRegistry.GetUtilTx().WithTx(ctx, doFunc, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}
	return ids, nil
}
func (s *LoketService) FindById(ctx context.Context, payload payload.RequestGetLoketById) (*model.Loket, error) {
	data, err := s.repoRegistry.GetLoketRepository().FindById(ctx, payload.Id)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}
	return data, nil
}

func (s *LoketService) FindByIdExternal(ctx context.Context, payload payload.RequestGetLoketById) (*model.Loket, error) {
	data, err := s.repoRegistry.GetLoketRepository().FindById(ctx, payload.Id)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}
	if data.Status != constant.StatusActive {
		err = errors.New("status data not active")
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}
	return data, nil
}

func (s *LoketService) FindAllExternal(ctx context.Context) ([]*model.LoketAntrian2, error) {
	data, err := s.repoRegistry.GetLoketRepository().FindAllExternal(ctx)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}

	return data, nil
}

func (s *LoketService) FindAll(ctx context.Context, payload payload.RequestGetLoket) ([]*model.Loket, uint64, error) {
	// const logCtx = "service.tipe_pasien.FindAll"

	data, total, err := s.repoRegistry.GetLoketRepository().FindAll(ctx, payload)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, 0, err
	}

	return data, total, nil

}
