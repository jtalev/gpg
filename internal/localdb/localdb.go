package localdb

import (
	"gpg/portal/internal/database"
	"gpg/portal/internal/user"
	"log"
	"time"
)

type Db struct {
	Ur database.UserRepository
}

func (db *Db) InitDb() {
	mockRepo := mockUserRepo{
		users: []user.User{},
	}
	db.Ur = &mockRepo
	db.hydrateUserTable()
}

func (db *Db) hydrateUserTable() {
	hash, err := user.HashPassword("test")
	if err != nil {
		log.Printf("error hashing password: %v", err)
	}
	user := user.NewUser(
		"1234567",
		1234567,
		"shliddy",
		"Shlid",
		"Dy",
		"test",
		hash,
		true,
		false,
		time.Date(2024, time.July, 3, 6, 6, 6, 6, time.Local),
		time.Date(2024, time.July, 3, 6, 6, 6, 6, time.Local),
	)
	mockRepo, _ := db.Ur.(*mockUserRepo)
	mockRepo.users = append(mockRepo.users, user)
}
