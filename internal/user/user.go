package user

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int    `json:"id"`
	EmployeeId   int    `json:"employee_id"`
	Username     string `json:"user_name"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`

	IsAdmin     bool      `json:"is_admin"`
	IsActivated bool      `json:"is_activated"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Printf("error: %v", err)
		return false
	}
	return true
}
