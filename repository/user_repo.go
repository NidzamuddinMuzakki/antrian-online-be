package repository

import (
	"antrian-golang/common/constant"
	commonDs "antrian-golang/common/data_source"
	"antrian-golang/common/logger"
	common "antrian-golang/common/registry"
	"antrian-golang/model"
	"context"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	common common.IRegistry
	master *sqlx.DB
}

type IUserRepo interface {
	FindUsernameLogin(ctx context.Context, req model.User) (*model.User, error)
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
