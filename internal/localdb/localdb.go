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
	user := user.User{
		Id:           1234567,
		EmployeeId:   1234567,
		Username:     "shliddy",
		FirstName:    "Shlid",
		LastName:     "Dy",
		Email:        "test",
		PasswordHash: "test",

		IsAdmin:     true,
		IsActivated: false,
		CreatedAt:   time.Date(2024, time.July, 3, 6, 6, 6, 6, time.Local),
		ModifiedAt:  time.Date(2024, time.July, 3, 6, 6, 6, 6, time.Local),
	}
	db.u = append(db.u, user)
}

func (db *Db) GetUsers() []user.User {
	users := make([]user.User, 0)
	for i, _ := range db.u {
		log.Println(db.u[i])
		users = append(users, db.u[i])
	}
	return users
}
