package payload

type RequestGetRole struct {
	Status     string `form:"status"`
	Name       string `form:"name"`
	RowPerpage uint   `form:"row_perpage" validate:"number"`
	Page       uint   `form:"page" validate:"number"`
	UserId     string `json:"user_id" validate:"required"`
}

type RequestCreateRole struct {
	Name   string `json:"name" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}

type RequestUpdateRoleStatus struct {
	Uuid   string `json:"uuid" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}

type RequestRoleFindByUUID struct {
	Uuid   string `uri:"uuid" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}

type RequestUpdateRole struct {
	Uuid   string `json:"uuid" validate:"required"`
	Name   string `json:"name" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}
