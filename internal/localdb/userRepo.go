package localdb

import (
	"errors"
	"gpg/portal/internal/user"
)

type mockUserRepo struct {
	users []user.User
}

func (m *mockUserRepo) GetUserByUsername(username string) (user.User, error) {
	for _, u := range m.users {
		if u.Username == username {
			return u, nil
		}
	}
	return user.User{}, errors.New("user not found")
}

func (m *mockUserRepo) GetUserById(id string) (user.User, error) {
	for _, u := range m.users {
		if u.Id == id {
			return u, nil
		}
	}
	return user.User{}, errors.New("user not found")
}
