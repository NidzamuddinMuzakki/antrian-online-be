package payload

type RequestLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
