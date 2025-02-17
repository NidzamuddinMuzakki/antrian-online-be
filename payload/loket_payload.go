package payload

type RequestGetLoket struct {
	Status     string `form:"status"`
	RowPerpage uint   `form:"row_perpage" validate:"number"`
	Page       uint   `form:"page" validate:"number"`
}

type RequestGetLoketById struct {
	Id int `uri:"id"`
}

type RequestUpdateLoketStatus struct {
	Id     int    `uri:"id"`
	UserId string `json:"user_id" validate:"required"`
}
type RequestUpdateLoketUserId struct {
	Id      int    `uri:"id"`
	UserId  string `json:"user_id" validate:"required"`
	UserId2 string `json:"user_id2" validate:"required"`
}

type RequestInsertLoket struct {
	Name   string `json:"name" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}

type RequestUpdateLoket struct {
	Id     int    `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}
