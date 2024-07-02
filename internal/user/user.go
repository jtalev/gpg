package user

import (
	"time"
)

type User struct {
	Id           int       `json:"id"`
	EmployeeId   int       `json:"employee_id"`
	Username     string    `json:"user_name"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Role         string    `json:"role"`
	IsActivated  bool      `json:"is_activated"`
	CreatedAt    time.Time `json:"created_at"`
	ModifiedAt   time.Time `json:"modified_at"`
}
