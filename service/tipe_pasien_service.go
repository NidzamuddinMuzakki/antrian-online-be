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
	"time"

	"github.com/jmoiron/sqlx"
)

type ITipePasienService interface {
	FindAll(ctx context.Context, payload payload.RequestGetTipePasien) ([]*model.TipePasien, uint64, error)
	FindAllExternal(ctx context.Context) ([]*model.TipePasien, error)
	FindById(ctx context.Context, payload payload.RequestGetTipePasienById) (*model.TipePasien, error)
	FindByIdExternal(ctx context.Context, payload payload.RequestGetTipePasienById) (*model.TipePasien, error)
	InsertData(ctx context.Context, payload payload.RequestInsertTipePasien) (int, error)
	UpdateData(ctx context.Context, payload payload.RequestUpdateTipePasien) error
	Activate(ctx context.Context, payload payload.RequestUpdateStatus) error
	DeActivate(ctx context.Context, payload payload.RequestUpdateStatus) error
}

type TipePasienService struct {
	common       registry.IRegistry
	repoRegistry repository.IRegistry
}

func NewTipePasienService(common registry.IRegistry, repoRegistry repository.IRegistry) ITipePasienService {
	return &TipePasienService{
		common:       common,
		repoRegistry: repoRegistry,
	}
}

func (s *TipePasienService) UpdateData(ctx context.Context, payload payload.RequestUpdateTipePasien) error {

	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		getData, err := s.repoRegistry.GetTipePasienRepository().FindById(ctx, payload.Id)
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
		m := model.TipePasien{
			Id:        payload.Id,
			Name:      payload.Name,
			UpdatedBy: payload.UserId,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetTipePasienRepository().Update(ctx, tx, m)
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

func (s *TipePasienService) Activate(ctx context.Context, payload payload.RequestUpdateStatus) error {

	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {

		_, err := s.repoRegistry.GetTipePasienRepository().FindById(ctx, payload.Id)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)
		m := model.TipePasien{
			Id:        payload.Id,
			Status:    constant.StatusActive,
			UpdatedBy: payload.UserId,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetTipePasienRepository().UpdateStatus(ctx, tx, m)
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

func (s *TipePasienService) DeActivate(ctx context.Context, payload payload.RequestUpdateStatus) error {

	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		_, err := s.repoRegistry.GetTipePasienRepository().FindById(ctx, payload.Id)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)
		m := model.TipePasien{
			Id:        payload.Id,
			Status:    constant.StatusNotActive,
			UpdatedBy: payload.UserId,
			UpdatedAt: &now,
		}
		err = s.repoRegistry.GetTipePasienRepository().UpdateStatus(ctx, tx, m)
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
func (s *TipePasienService) InsertData(ctx context.Context, payload payload.RequestInsertTipePasien) (int, error) {
	ids := 0
	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)
		m := model.TipePasien{
			Name:      payload.Name,
			Status:    constant.StatusActive,
			CreatedBy: payload.UserId,
			CreatedAt: now,
		}
		idR, err := s.repoRegistry.GetTipePasienRepository().Insert(ctx, tx, m)
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
func (s *TipePasienService) FindById(ctx context.Context, payload payload.RequestGetTipePasienById) (*model.TipePasien, error) {
	data, err := s.repoRegistry.GetTipePasienRepository().FindById(ctx, payload.Id)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}
	return data, nil
}

func (s *TipePasienService) FindByIdExternal(ctx context.Context, payload payload.RequestGetTipePasienById) (*model.TipePasien, error) {
	data, err := s.repoRegistry.GetTipePasienRepository().FindById(ctx, payload.Id)
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

func (s *TipePasienService) FindAllExternal(ctx context.Context) ([]*model.TipePasien, error) {
	data, err := s.repoRegistry.GetTipePasienRepository().FindAllExternal(ctx)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}

	return data, nil
}

func (s *TipePasienService) FindAll(ctx context.Context, payload payload.RequestGetTipePasien) ([]*model.TipePasien, uint64, error) {
	// const logCtx = "service.tipe_pasien.FindAll"

	data, total, err := s.repoRegistry.GetTipePasienRepository().FindAll(ctx, payload)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, 0, err
	}

	return data, total, nil

}
