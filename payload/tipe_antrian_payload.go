package payload

type RequestGetTipePasien struct {
	Status     string `form:"status"`
	RowPerpage uint   `form:"row_perpage" validate:"number"`
	Page       uint   `form:"page" validate:"number"`
}
type RequestGetTipePasienById struct {
	Id int `uri:"id"`
}

type RequestUpdateStatus struct {
	Id     int    `uri:"id"`
	UserId string `json:"user_id" validate:"required"`
}

type RequestInsertTipePasien struct {
	Name   string `json:"name" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}

type RequestUpdateTipePasien struct {
	Id     int    `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}
