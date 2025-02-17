package model

import "time"

type Loket struct {
	Id        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	UserId    int        `json:"user_id" db:"user_id"`
	Status    string     `json:"status" db:"status"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy string     `json:"updated_by" db:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type LoketAntrian struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	UserId int    `json:"user_id" db:"user_id"`
	Status string `json:"status" db:"status"`
	Number int    `json:"number" db:"number"`

	CreatedBy string     `json:"created_by" db:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy string     `json:"updated_by" db:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type LoketAntrian2 struct {
	Id             int        `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	UserId         int        `json:"user_id" db:"user_id"`
	Status         string     `json:"status" db:"status"`
	Number         int        `json:"number" db:"number"`
	TipePasienName string     `json:"tipe_pasien_name" db:"tipe_pasien_name"`
	CreatedBy      string     `json:"created_by" db:"created_by"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy      string     `json:"updated_by" db:"updated_by"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
}
