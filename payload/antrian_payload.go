package payload

type AntrianPayloadInsert struct {
	TipePasienId int `json:"tipe_pasien_id"`
}

type RequestGetAntrian struct {
	UserId     int  `form:"user_id" validate:"required"`
	RowPerpage uint `form:"row_perpage" validate:"number"`
	Page       uint `form:"page" validate:"number"`
}
