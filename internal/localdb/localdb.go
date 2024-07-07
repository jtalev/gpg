package localdb

import (
	"gpg/portal/internal/user"
	"log"
	"time"
)

type Db struct {
	u []user.User
}

func (db *Db) InitDb() {
	db.u = []user.User{}
	db.hydrateUserTable()
}

func (db *Db) hydrateUserTable() {
	hash, err := user.HashPassword("test")
	if err != nil {
		log.Printf("error hashing password: %v", err)
	}
	user := user.User{
		Id:           1234567,
		EmployeeId:   1234567,
		Username:     "shliddy",
		FirstName:    "Shlid",
		LastName:     "Dy",
		Email:        "test",
		PasswordHash: hash,

		IsAdmin:     true,
		IsActivated: false,
		CreatedAt:   time.Date(2024, time.July, 3, 6, 6, 6, 6, time.Local),
		ModifiedAt:  time.Date(2024, time.July, 3, 6, 6, 6, 6, time.Local),
	}
	db.u = append(db.u, user)
}

func (db *Db) GetUsers() []user.User {
	return db.u
}

func (db *Db) GetUserById(id int) *user.User {
	for _, u := range db.u {
		if u.Id == id {
			return &u
		}
	}
	log.Println("user not found")
	return nil
}
