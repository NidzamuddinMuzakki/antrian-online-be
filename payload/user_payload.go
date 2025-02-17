package payload

type RequestLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RequestCreateUser struct {
	Username string `json:"username" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate:"required"`
	UserId   string `json:"user_id" validate:"required"`
}

type RequestUpdateUserStatus struct {
	Id     int    `json:"id" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}

type RequestUpdateUserFindById struct {
	Id     int    `uri:"id" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}

type RequestUpdateUserPassword struct {
	Id       int    `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
	UserId   string `json:"user_id" validate:"required"`
}

type RequestUpdateUser struct {
	Id       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Role     string `json:"role" validate:"required"`
	UserId   string `json:"user_id" validate:"required"`
}

type RequestGetUser struct {
	UserId     string `json:"user_id" validate:"required"`
	Status     string `form:"status"`
	Username   string `form:"username"`
	Role       string `form:"role"`
	RowPerpage uint   `form:"row_perpage" validate:"number"`
	Page       uint   `form:"page" validate:"number"`
}
