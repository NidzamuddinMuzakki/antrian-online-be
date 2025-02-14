package repository

import (
	"antrian-golang/common/constant"
	commonDs "antrian-golang/common/data_source"
	"antrian-golang/common/logger"
	common "antrian-golang/common/registry"
	"antrian-golang/model"
	"antrian-golang/payload"
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TipePasienRepo struct {
	common common.IRegistry
	master *sqlx.DB
}

type ITipePasienRepo interface {
	FindAll(ctx context.Context, payload payload.RequestGetTipePasien) ([]*model.TipePasien, uint64, error)
	FindAllExternal(ctx context.Context) ([]*model.TipePasien, error)
	Insert(ctx context.Context, tx *sqlx.Tx, payload model.TipePasien) (int, error)
	Update(ctx context.Context, tx *sqlx.Tx, payload model.TipePasien) error
	UpdateStatus(ctx context.Context, tx *sqlx.Tx, payload model.TipePasien) error
	FindById(ctx context.Context, id int) (*model.TipePasien, error)
}

func NewTipePasienRepo(common common.IRegistry, master *sqlx.DB) ITipePasienRepo {
	return &TipePasienRepo{
		common: common,
		master: master,
	}
}
func (r *TipePasienRepo) Insert(ctx context.Context, tx *sqlx.Tx, payload model.TipePasien) (int, error) {
	query := "insert into master_tipe_pasien (name, status, created_by,created_at) values (?,?,?,?)"
	stmx, err := tx.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}
	res, err := stmx.ExecContext(ctx,
		payload.Name,
		payload.Status,
		payload.CreatedBy,
		payload.CreatedAt,
	)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}
	intss, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	if intss == 0 {
		err = errors.New("nothing insert")
		return 0, err
	}

	return int(intss), nil
}

func (r *TipePasienRepo) Update(ctx context.Context, tx *sqlx.Tx, payload model.TipePasien) error {
	query := "update master_tipe_pasien set name=?, updated_by=?,updated_at=? where id=? and status=?"
	stmx, err := tx.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	res, err := stmx.ExecContext(ctx,
		payload.Name,
		payload.UpdatedBy,
		payload.UpdatedAt,
		payload.Id,
		constant.StatusActive,
	)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	intss, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if intss == 0 {
		err = errors.New("nothing update")
		return err
	}

	return nil
}

func (r *TipePasienRepo) UpdateStatus(ctx context.Context, tx *sqlx.Tx, payload model.TipePasien) error {
	query := "update master_tipe_pasien set status=?, updated_by=?,updated_at=? where id=?"
	stmx, err := tx.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	res, err := stmx.ExecContext(ctx,
		payload.Status,
		payload.UpdatedBy,
		payload.UpdatedAt,
		payload.Id,
	)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	intss, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if intss == 0 {
		err = errors.New("nothing update")
		return err
	}

	return nil
}

func (r *TipePasienRepo) FindById(ctx context.Context, id int) (*model.TipePasien, error) {

	const logCtx = "repository.masterTnc.GetId"
	var (
		data model.TipePasien
	)

	selectQuery := "SELECT "

	columns := commonDs.GetDbColumns(model.TipePasien{})
	for indexC, kVal := range columns {
		if indexC == len(columns)-1 {
			selectQuery += kVal

		} else {
			selectQuery += kVal + ","

		}
	}
	selectQuery += " FROM master_tipe_pasien where id = ? "

	// fmt.Println(selectQuery, countQuery)
	err := commonDs.Exec(ctx, r.master, commonDs.NewStatement(&data, selectQuery, id))
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}

	return &data, nil

}

func (r *TipePasienRepo) FindAll(ctx context.Context, payload payload.RequestGetTipePasien) ([]*model.TipePasien, uint64, error) {

	const logCtx = "repository.masterTnc.Get"
	var (
		list              []*model.TipePasien
		totalTransactions uint64
		filters           []string
		args              []any
	)

	countQuery := "SELECT count(id)"
	selectQuery := "SELECT "

	columns := commonDs.GetDbColumns(model.TipePasien{})
	for indexC, kVal := range columns {
		if indexC == len(columns)-1 {
			selectQuery += kVal

		} else {
			selectQuery += kVal + ","

		}
	}
	selectQuery += " FROM master_tipe_pasien"
	countQuery += " FROM master_tipe_pasien"

	if payload.Status != "" {
		filters = append(filters, "and status=? ")
		args = append(args, payload.Status)
	}

	for _, f := range filters {
		countQuery = fmt.Sprintf("%s %s", countQuery, f)
		selectQuery = fmt.Sprintf("%s %s", selectQuery, f)
	}

	offset := (payload.RowPerpage * payload.Page) - payload.RowPerpage
	selectQuery = fmt.Sprintf("%s ORDER BY id LIMIT %d OFFSET %d", selectQuery, payload.RowPerpage, offset)

	// fmt.Println(selectQuery, countQuery)
	err := commonDs.Exec(ctx, r.master, commonDs.NewStatement(&list, selectQuery, args...))
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, 0, err
	}

	err = commonDs.Exec(ctx, r.master, commonDs.NewStatement(&totalTransactions, countQuery, args...))
	if err != nil {

		logger.Error(ctx, err.Error(), err)
		return nil, 0, err
	}

	return list, totalTransactions, nil

}

func (r *TipePasienRepo) FindAllExternal(ctx context.Context) ([]*model.TipePasien, error) {

	const logCtx = "repository.masterTnc.Get"
	var (
		list []*model.TipePasien
	)

	selectQuery := "SELECT "

	columns := commonDs.GetDbColumns(model.TipePasien{})
	for indexC, kVal := range columns {
		if indexC == len(columns)-1 {
			selectQuery += kVal

		} else {
			selectQuery += kVal + ","

		}
	}
	selectQuery += " FROM master_tipe_pasien where status=? ORDER BY id DESC"

	// fmt.Println(selectQuery, countQuery)
	err := commonDs.Exec(ctx, r.master, commonDs.NewStatement(&list, selectQuery, constant.StatusActive))
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}

	return list, nil

}
