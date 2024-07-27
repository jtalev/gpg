package database

import (
	"gpg/portal/internal/user"
)

type UserRepository interface {
	GetUserByUsername(username string) (user.User, error)
}

type Db struct {
	UserRepo UserRepository
}

func NewDb(userRepo UserRepository) *Db {
	return &Db{
		UserRepo: userRepo,
	}
}