package service

import (
	commonDs "antrian-golang/common/data_source"
	"antrian-golang/common/logger"
	"antrian-golang/common/registry"
	commonTime "antrian-golang/common/time"
	"antrian-golang/model"
	"antrian-golang/payload"
	"antrian-golang/repository"
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type DataModelNumber struct {
	Number int    `json:"number"`
	Loket  string `json:"loket"`
}
type IAntrianService interface {
	InsertData(ctx context.Context, payload payload.AntrianPayloadInsert) (DataModelNumber, error)
	UpdateData(ctx context.Context, payload model.Antrian) error

	FindAll(ctx context.Context, payload payload.RequestGetAntrian) ([]*model.Antrian2, uint64, error)
}

type AntrianService struct {
	common       registry.IRegistry
	repoRegistry repository.IRegistry
}

func NewAntrianService(common registry.IRegistry, repoRegistry repository.IRegistry) IAntrianService {
	return &AntrianService{
		common:       common,
		repoRegistry: repoRegistry,
	}
}

func (s *AntrianService) UpdateData(ctx context.Context, payload model.Antrian) error {
	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)

		m := model.Antrian{
			Id:           payload.Id,
			TipePasienId: payload.TipePasienId,
			LoketId:      payload.LoketId,
			Status:       payload.Status,
			UpdatedBy:    payload.UpdatedBy,
			UpdatedAt:    &now,
		}

		err := s.repoRegistry.GetAntrianRepository().UpdateData(ctx, tx, m)
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
func (s *AntrianService) InsertData(ctx context.Context, payload payload.AntrianPayloadInsert) (DataModelNumber, error) {
	aaa := DataModelNumber{}
	doFunc := commonDs.TxFunc(func(tx *sqlx.Tx) error {
		loc := commonTime.LoadTimeZoneAsiaJakarta()
		now := time.Now().In(loc)
		getDataTipe, err := s.repoRegistry.GetTipePasienRepository().FindById(ctx, payload.TipePasienId)
		if err != nil {

			logger.Error(ctx, err.Error(), err)
			return err

		}
		aaa.Loket = getDataTipe.Name
		m := model.Antrian{
			TipePasienId: payload.TipePasienId,
			Number:       0,
			CreatedBy:    "SYSTEM",
			CreatedAt:    now,
		}
		dataLast, err := s.repoRegistry.GetAntrianRepository().FindLastNumber(ctx, tx, m)
		if err != nil {
			if err.Error() != "sql: no rows in result set" {

				logger.Error(ctx, err.Error(), err)
				return err
			}
		}
		m.Number = dataLast.Number + 1
		_, err = s.repoRegistry.GetAntrianRepository().Insert(ctx, tx, m)
		if err != nil {
			logger.Error(ctx, err.Error(), err)
			return err
		}
		aaa.Number = dataLast.Number + 1
		return nil
	})

	err := s.repoRegistry.GetUtilTx().WithTx(ctx, doFunc, nil)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return aaa, err
	}
	return aaa, nil
}

func (s *AntrianService) FindAll(ctx context.Context, payload payload.RequestGetAntrian) ([]*model.Antrian2, uint64, error) {
	if payload.Page < 1 {
		payload.Page = 1
	}

	if payload.RowPerpage < 1 {
		payload.RowPerpage = 1
	}

	data, count, err := s.repoRegistry.GetAntrianRepository().FindAll(ctx, payload)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, 0, err
	}
	fmt.Println("result-data", data, count)

	return data, count, nil
}
