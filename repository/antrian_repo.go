package repository

import (
	commonDs "antrian-golang/common/data_source"
	"antrian-golang/common/logger"
	common "antrian-golang/common/registry"
	commonTime "antrian-golang/common/time"
	"antrian-golang/model"
	"antrian-golang/payload"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type AntrianRepo struct {
	common common.IRegistry
	master *sqlx.DB
}

type IAntrianRepo interface {
	Insert(ctx context.Context, tx *sqlx.Tx, payload model.Antrian) (int, error)
	FindLastNumber(ctx context.Context, tx *sqlx.Tx, payload model.Antrian) (model.Antrian, error)
	FindAll(ctx context.Context, payload payload.RequestGetAntrian) ([]*model.Antrian2, uint64, error)
	UpdateData(ctx context.Context, tx *sqlx.Tx, payload model.Antrian) error
}

func NewAntrianRepo(common common.IRegistry, master *sqlx.DB) IAntrianRepo {
	return &AntrianRepo{
		common: common,
		master: master,
	}
}

func (r *AntrianRepo) UpdateData(ctx context.Context, tx *sqlx.Tx, payload model.Antrian) error {
	query := "update master_antrian set status=?,tipe_pasien_id=?,loket_id=?, updated_by=?,updated_at=? where id=? "
	fmt.Println("nidzam-query", payload)
	stmx, err := tx.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return err
	}
	res, err := stmx.ExecContext(ctx,
		payload.Status,
		payload.TipePasienId,
		payload.LoketId,
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
func (r *AntrianRepo) FindAll(ctx context.Context, payload payload.RequestGetAntrian) ([]*model.Antrian2, uint64, error) {

	const logCtx = "repository.masterTnc.Get"
	var (
		list              []*model.Antrian2
		totalTransactions uint64

		// args []any
	)

	selectQuery := "SELECT "
	loc := commonTime.LoadTimeZoneAsiaJakarta()
	now := time.Now().In(loc)
	columns := commonDs.GetDbColumns(model.Antrian2{})
	for indexC, kVal := range columns {
		if kVal == "count" || kVal == "tipe_pasien_name" {
			continue
		} else {

			if indexC == len(columns)-1 {
				selectQuery += "ma." + kVal

			} else {
				selectQuery += "ma." + kVal + ","

			}
		}
	}
	selectQuery += " , count(ma.id) as count, mtp.name as tipe_pasien_name FROM master_antrian ma"

	selectQuery = fmt.Sprintf("%s left join master_tipe_pasien mtp on ma.tipe_pasien_id=mtp.id left join master_loket ml on ml.id=ma.loket_id where DATE(ma.created_at)='%s' and ((ma.status='' && ma.loket_id=0) or (ma.status='call' and ml.user_id='%d' )) GROUP BY ma.tipe_pasien_id ORDER BY ma.created_at asc", selectQuery, now.Format("2006-01-02"), payload.UserId)

	fmt.Println(selectQuery, "nidzam-ganteng", payload.UserId)
	err := commonDs.Exec(ctx, r.master, commonDs.NewStatement(&list, selectQuery))
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return nil, 0, err
	}

	return list, totalTransactions, nil

}

func (r *AntrianRepo) FindLastNumber(ctx context.Context, tx *sqlx.Tx, payload model.Antrian) (model.Antrian, error) {
	const logCtx = "repository.masterTnc.GetId"
	var (
		data model.Antrian
	)

	selectQuery := "SELECT "

	columns := commonDs.GetDbColumns(model.Antrian{})
	for indexC, kVal := range columns {
		if indexC == len(columns)-1 {
			selectQuery += kVal

		} else {
			selectQuery += kVal + ","

		}
	}
	selectQuery += " FROM master_antrian where tipe_pasien_id = ? and DATE(created_at)=DATE(?) ORDER BY created_at desc"

	// fmt.Println(selectQuery, countQuery)
	stmx, err := tx.PrepareContext(ctx, selectQuery)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return data, err
	}
	err = stmx.QueryRowContext(ctx, payload.TipePasienId, payload.CreatedAt).Scan(&data.Id, &data.Number, &data.TipePasienId, &data.LoketId, &data.Status, &data.CreatedBy, &data.CreatedAt, &data.UpdatedBy, &data.UpdatedAt)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return data, err
	}

	return data, nil
}

func (r *AntrianRepo) Insert(ctx context.Context, tx *sqlx.Tx, payload model.Antrian) (int, error) {
	query := "insert into master_antrian (number,tipe_pasien_id, created_by,created_at) values (?,?,?,?)"
	stmx, err := tx.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}

	rs, err := stmx.ExecContext(ctx,
		payload.Number,
		payload.TipePasienId,
		payload.CreatedBy,
		payload.CreatedAt,
	)
	intss, err := rs.RowsAffected()
	if err != nil {
		logger.Error(ctx, err.Error(), err)
		return 0, err
	}

	if intss == 0 {
		err = errors.New("nothing insert")
		return 0, err
	}

	return int(intss), nil
}
