package validation

import (
	"gpg/portal/internal/database"
	"gpg/portal/internal/user"
)

func ValidateLogin(repo database.UserRepository, username, password string) ValidationResult {
	result := ValidationResult{IsValid: true, Msg: ""}

	u, err := repo.GetUserByUsername(username)
	if err != nil {
		result.IsValid = false
		result.Msg = "invalid username or password"
		return result
	}
	if !user.CheckPasswordHash(u.PasswordHash, password) {
		result.IsValid = false
		result.Msg = "invalid username or password"
		return result
	}

	return result
}
