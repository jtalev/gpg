package validation

import (
	"gpg/portal/internal/localdb"
	"testing"
)

var db = localdb.Db{}

func TestValidateLogin(t *testing.T) {
	db.InitDb()
	expected := ValidationResult{IsValid: true, Msg: ""}
	result := ValidateLogin(db.Ur, "shliddy", "test")
	if result != expected {
		t.Errorf("login validation failed: want=%v, got=%v", expected, result)
	}

	expected.IsValid = false
	expected.Msg = "invalid username or password"
	result = ValidateLogin(db.Ur, "shliddy", "password")
	if result != expected {
		t.Errorf("unexpected validation: want=%v, got=%v", expected, result)
	}
	result = ValidateLogin(db.Ur, "thisusernamedoesntexist", "test")
	if result != expected {
		t.Errorf("unexpected validation: want=%v, got=%v", expected, result)
	}
}