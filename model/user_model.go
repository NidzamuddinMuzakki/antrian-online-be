package model

import "time"

type User struct {
	Id        int        `json:"id" db:"id"`
	Username  string     `json:"username" db:"username"`
	Password  string     `json:"password" db:"password"`
	Role      string     `json:"role" db:"role"`
	Status    string     `json:"status" db:"status"`
	CreatedBy string     `json:"created_by" db:"created_by"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedBy string     `json:"updated_by" db:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
