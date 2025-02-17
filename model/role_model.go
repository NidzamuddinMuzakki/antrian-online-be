package model

import "time"

type Role struct {
	Id        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	RoleUuid  string     `json:"role_uuid" db:"role_uuid"`
	Status    string     `json:"status" db:"status"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy string     `json:"updated_by" db:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
