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

type UserRepo struct {
	common common.IRegistry
	master *sqlx.DB
}

type IUserRepo interface {
	FindUsernameLogin(ctx context.Context, req model.User) (*model.User, error)
	InsertData(ctx context.Context, tx *sqlx.Tx, payload model.User) (int, error)
	UpdateData(ctx context.Context, tx *sqlx.Tx, payload model.User) error
	UpdateStatus(ctx context.Context, tx *sqlx.Tx, payload model.User) error
	UpdatePassword(ctx context.Context, tx *sqlx.Tx, payload model.User) error

	FindAll(ctx context.Context, paylod payload.RequestGetUser) ([]*model.User, uint64, error)
	FindById(ctx context.Context, id int) (*model.User, error)
}

func NewUserRepo(common common.IRegistry, master *sqlx.DB) IUserRepo {
	return &UserRepo{
		common: common,
		master: master,
	}
}

func (r *UserRepo) FindUsernameLogin(ctx context.Context, req model.User) (*model.User, error) {

	const logCtx = "repository.masterTnc.GetId"
	var (
		data model.User
	)

	selectQuery := "SELECT "

	columns := commonDs.GetDbColumns(model.User{})
	for indexC, kVal := range columns {
		if indexC == len(columns)-1 {
			selectQuery += kVal

		} else {
			selectQuery += kVal + ","

		}
	}
	selectQuery += " FROM master_user where username = ? and status=? "

	// fmt.Println(selectQuery, countQuery)
	err := commonDs.Exec(ctx, r.master, commonDs.NewStatement(&data, selectQuery, req.Username, constant.StatusActive))
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}

	return &data, nil

}

func (r *UserRepo) FindById(ctx context.Context, id int) (*model.User, error) {
	const logCtx = "repository.masterTnc.GetId"
	var (
		data model.User
	)

	selectQuery := "SELECT "

	columns := commonDs.GetDbColumns(model.User{})
	for indexC, kVal := range columns {
		if indexC == len(columns)-1 {
			selectQuery += kVal

		} else {
			selectQuery += kVal + ","

		}
	}
	selectQuery += " FROM master_user where id = ? and status=? "

	// fmt.Println(selectQuery, countQuery)
	err := commonDs.Exec(ctx, r.master, commonDs.NewStatement(&data, selectQuery, id, constant.StatusActive))
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}

	return &data, nil

}

func (r *UserRepo) FindAll(ctx context.Context, payload payload.RequestGetUser) ([]*model.User, uint64, error) {

	const logCtx = "repository.masterTnc.Get"
	var (
		list              []*model.User
		totalTransactions uint64
		filters           []string
		args              []any
	)

	countQuery := "SELECT count(id)"
	selectQuery := "SELECT "

	columns := commonDs.GetDbColumns(model.Role{})
	for indexC, kVal := range columns {
		if indexC == len(columns)-1 {
			selectQuery += kVal

		} else {
			selectQuery += kVal + ","

		}
	}
	selectQuery += " FROM master_user"
	countQuery += " FROM master_user"

	if payload.Status != "" {
		filters = append(filters, "and status=? ")
		args = append(args, payload.Status)
	}

	if payload.Username != "" {
		filters = append(filters, "and username=? ")
		args = append(args, payload.Username)
	}

	if payload.Role != "" {
		filters = append(filters, "and role=? ")
		args = append(args, payload.Role)
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

func (r *UserRepo) InsertData(ctx context.Context, tx *sqlx.Tx, payload model.User) (int, error) {
	query := "insert into master_user (username,role, password,status, created_by,created_at) values (?,?,?,?,?,?)"
	stmx, err := tx.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}
	res, err := stmx.ExecContext(ctx,
		payload.Username,
		payload.Role,
		payload.Password,
		constant.StatusActive,
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

func (r *UserRepo) UpdateData(ctx context.Context, tx *sqlx.Tx, payload model.User) error {
	query := "update master_user set username=?,role=?, updated_by=?,updated_at=? where id=? and status=?"
	stmx, err := tx.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	res, err := stmx.ExecContext(ctx,
		payload.Username,
		payload.Role,
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

func (r *UserRepo) UpdateStatus(ctx context.Context, tx *sqlx.Tx, payload model.User) error {
	query := "update master_user set status=?, updated_by=?,updated_at=? where id=?"
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

func (r *UserRepo) UpdatePassword(ctx context.Context, tx *sqlx.Tx, payload model.User) error {
	query := "update master_user set password=?, updated_by=?,updated_at=? where id=?"
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
