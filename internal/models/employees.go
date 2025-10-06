package models

import "time"

type Employee struct {
	ID       int       `json:"id" db:"id"`
	FullName string    `json:"full_name" db:"full_name"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Role     Role      `json:"role" db:"role"`
	CreateAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt time.Time `json:"updated_at" db:"updated_at"`
}

type Role string

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)
