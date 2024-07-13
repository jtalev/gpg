package database

import (
	"gpg/portal/internal/user"
)

type UserRepository interface {
	GetUserByUsername(username string) (user.User, error)
}

type Db struct {
	Ur UserRepository
}

func NewDb(ur UserRepository) *Db {
	return &Db{
		Ur: ur,
	}
}