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

type LoketRepo struct {
	common common.IRegistry
	master *sqlx.DB
}

type ILoketRepo interface {
	FindAll(ctx context.Context, payload payload.RequestGetLoket) ([]*model.Loket, uint64, error)
	FindAllExternal(ctx context.Context) ([]*model.LoketAntrian2, error)
	Insert(ctx context.Context, tx *sqlx.Tx, payload model.Loket) (int, error)
	Update(ctx context.Context, tx *sqlx.Tx, payload model.Loket) error
	UpdateStatus(ctx context.Context, tx *sqlx.Tx, payload model.Loket) error
	FindById(ctx context.Context, id int) (*model.Loket, error)
	UpdateUserId(ctx context.Context, tx *sqlx.Tx, payload model.Loket, userID int) error
}

func NewLoketRepo(common common.IRegistry, master *sqlx.DB) ILoketRepo {
	return &LoketRepo{
		common: common,
		master: master,
	}
}
func (r *LoketRepo) Insert(ctx context.Context, tx *sqlx.Tx, payload model.Loket) (int, error) {
	query := "insert into master_loket (name, status, created_by,created_at) values (?,?,?,?)"
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

func (r *LoketRepo) Update(ctx context.Context, tx *sqlx.Tx, payload model.Loket) error {
	query := "update master_loket set name=?, updated_by=?,updated_at=? where id=? and status=?"
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

func (r *LoketRepo) UpdateStatus(ctx context.Context, tx *sqlx.Tx, payload model.Loket) error {
	query := "update master_loket set status=?, updated_by=?,updated_at=? where id=?"
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

func (r *LoketRepo) UpdateUserId(ctx context.Context, tx *sqlx.Tx, payload model.Loket, userID int) error {
	query := "update master_loket set user_id=?, updated_by=?,updated_at=? where id=? and (user_id=? or user_id=? or (0=? and user_id=?)) and status=?"
	stmx, err := tx.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	fmt.Println("nidzamganteng", query, payload.UserId, payload.Id, userID)
	res, err := stmx.ExecContext(ctx,
		payload.UserId,
		payload.UpdatedBy,
		payload.UpdatedAt,
		payload.Id,
		0,
		payload.UserId,
		payload.UserId,
		userID,
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

func (r *LoketRepo) FindById(ctx context.Context, id int) (*model.Loket, error) {

	const logCtx = "repository.masterTnc.GetId"
	var (
		data model.Loket
	)

	selectQuery := "SELECT "

	columns := commonDs.GetDbColumns(model.Loket{})
	for indexC, kVal := range columns {
		if indexC == len(columns)-1 {
			selectQuery += kVal

		} else {
			selectQuery += kVal + ","

		}
	}
	selectQuery += " FROM master_loket where id = ? "

	// fmt.Println(selectQuery, countQuery)
	err := commonDs.Exec(ctx, r.master, commonDs.NewStatement(&data, selectQuery, id))
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}

	return &data, nil

}

func (r *LoketRepo) FindAll(ctx context.Context, payload payload.RequestGetLoket) ([]*model.Loket, uint64, error) {

	const logCtx = "repository.masterTnc.Get"
	var (
		list              []*model.Loket
		totalTransactions uint64
		filters           []string
		args              []any
	)

	countQuery := "SELECT count(id)"
	selectQuery := "SELECT "

	columns := commonDs.GetDbColumns(model.Loket{})
	for indexC, kVal := range columns {
		if indexC == len(columns)-1 {
			selectQuery += kVal

		} else {
			selectQuery += kVal + ","

		}
	}
	selectQuery += " FROM master_loket"
	countQuery += " FROM master_loket"

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

func (r *LoketRepo) FindAllExternal(ctx context.Context) ([]*model.LoketAntrian2, error) {

	const logCtx = "repository.masterTnc.Get"
	var (
		list []*model.LoketAntrian2
	)

	selectQuery := "SELECT "

	columns := commonDs.GetDbColumns(model.Loket{})
	for indexC, kVal := range columns {
		kVal = "ml." + kVal
		if indexC == len(columns)-1 {
			selectQuery += kVal

		} else {
			selectQuery += kVal + ","

		}
	}
	selectQuery += " , ma.number as number, mtp.name as tipe_pasien_name FROM master_loket ml left join master_antrian ma on ml.id=ma.loket_id left join master_tipe_pasien mtp on ma.tipe_pasien_id=mtp.id where ml.status=? and ma.status='call' ORDER BY id DESC"

	// fmt.Println(selectQuery, countQuery)
	err := commonDs.Exec(ctx, r.master, commonDs.NewStatement(&list, selectQuery, constant.StatusActive))
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}

	return list, nil

}
