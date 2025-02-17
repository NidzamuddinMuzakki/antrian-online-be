package model

import "time"

type Antrian struct {
	Id           int        `json:"id" db:"id"`
	Number       int        `json:"number" db:"number"`
	TipePasienId int        `json:"tipe_pasien_id" db:"tipe_pasien_id"`
	LoketId      int        `json:"loket_id" db:"loket_id"`
	Status       string     `json:"status" db:"status"`
	CreatedBy    string     `json:"created_by" db:"created_by"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy    string     `json:"updated_by" db:"updated_by"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
}

type Antrian2 struct {
	Id             int    `json:"id" db:"id"`
	Number         int    `json:"number" db:"number"`
	TipePasienId   int    `json:"tipe_pasien_id" db:"tipe_pasien_id"`
	TipePasienName string `json:"tipe_pasien_name" db:"tipe_pasien_name"`

	LoketId   int        `json:"loket_id" db:"loket_id"`
	Status    string     `json:"status" db:"status"`
	Count     int        `json:"count" db:"count"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy string     `json:"updated_by" db:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
