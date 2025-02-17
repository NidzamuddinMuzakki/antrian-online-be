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

type RoleRepo struct {
	common common.IRegistry
	master *sqlx.DB
}

type IRoleRepo interface {
	FindAll(ctx context.Context, paylod payload.RequestGetRole) ([]*model.Role, uint64, error)
	FindByUuid(ctx context.Context, uuid string) (*model.Role, error)
	InsertData(ctx context.Context, tx *sqlx.Tx, payload model.Role) (string, error)
	UpdateData(ctx context.Context, tx *sqlx.Tx, payload model.Role) error
	UpdateStatusRole(ctx context.Context, tx *sqlx.Tx, payload model.Role) error
}

func NewRoleRepo(common common.IRegistry, master *sqlx.DB) IRoleRepo {
	return &RoleRepo{
		common: common,
		master: master,
	}
}

func (r *RoleRepo) FindByUuid(ctx context.Context, uuid string) (*model.Role, error) {

	const logCtx = "repository.masterTnc.GetId"
	var (
		data model.Role
	)

	selectQuery := "SELECT "

	columns := commonDs.GetDbColumns(model.Role{})
	for indexC, kVal := range columns {
		if indexC == len(columns)-1 {
			selectQuery += kVal

		} else {
			selectQuery += kVal + ","

		}
	}
	selectQuery += " FROM master_role where role_uuid = ? and status=? "

	// fmt.Println(selectQuery, countQuery)
	err := commonDs.Exec(ctx, r.master, commonDs.NewStatement(&data, selectQuery, uuid, constant.StatusActive))
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, err
	}

	return &data, nil

}

func (r *RoleRepo) FindAll(ctx context.Context, payload payload.RequestGetRole) ([]*model.Role, uint64, error) {

	const logCtx = "repository.masterTnc.Get"
	var (
		list              []*model.Role
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
	selectQuery += " FROM master_role"
	countQuery += " FROM master_role"

	if payload.Status != "" {
		filters = append(filters, "and status=? ")
		args = append(args, payload.Status)
	}

	if payload.Name != "" {
		filters = append(filters, "and name=? ")
		args = append(args, payload.Name)
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

func (r *RoleRepo) InsertData(ctx context.Context, tx *sqlx.Tx, payload model.Role) (string, error) {
	query := "insert into master_role (name,role_uuid, status, created_by,created_at) values (?,?,?,?,?)"
	stmx, err := tx.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return "", err
	}
	res, err := stmx.ExecContext(ctx,
		payload.Name,
		payload.RoleUuid,
		constant.StatusActive,
		payload.CreatedBy,
		payload.CreatedAt,
	)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return "", err
	}
	intss, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	if intss == 0 {
		err = errors.New("nothing insert")
		return "", err
	}

	return payload.RoleUuid, nil
}

func (r *RoleRepo) UpdateData(ctx context.Context, tx *sqlx.Tx, payload model.Role) error {
	query := "update master_role set name=?, updated_by=?,updated_at=? where role_uuid=? and status=?"
	stmx, err := tx.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	res, err := stmx.ExecContext(ctx,
		payload.Name,
		payload.UpdatedBy,
		payload.UpdatedAt,
		payload.RoleUuid,
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

func (r *RoleRepo) UpdateStatusRole(ctx context.Context, tx *sqlx.Tx, payload model.Role) error {
	query := "update master_role set status=?, updated_by=?,updated_at=? where role_uuid=?"
	stmx, err := tx.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	res, err := stmx.ExecContext(ctx,
		payload.Status,
		payload.UpdatedBy,
		payload.UpdatedAt,
		payload.RoleUuid,
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
