package user

import "testing"

func TestHashPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}
	if !CheckPasswordHash(hash, password) {
		t.Errorf("error mismatched hash and password")
	}

	expected := false
	hash, _ = HashPassword("password1")
	result := CheckPasswordHash(hash, password)
	if result != expected {
		t.Errorf("error hash and password should be mismatched: want=%v, got=%v", expected, result)
	}
}