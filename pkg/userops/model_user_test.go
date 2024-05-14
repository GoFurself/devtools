package userops

import (
	"testing"
)

func TestNewUserWithUserOptions(t *testing.T) {

	firstname := "Einar"
	lastname := "Nystedt"
	email := "einar.nystedt@centria.fi"
	password := "password"

	user := NewUser(email, password,
		WithFirstName(firstname),
		WithLastName(lastname),
		WithRole(1))

	if user.FirstName != "Einar" {
		t.Errorf("Expected first name Einar, got %s", user.FirstName)
	}
	if user.LastName != "Nystedt" {
		t.Errorf("Expected last name %s, got %s", email, user.LastName)
	}
	if user.Email != email {
		t.Errorf("Expected email %s, got %s", email, user.Email)
	}
	if user.Password != password {
		t.Errorf("Expected password %s, got %s", password, user.Password)
	}
	if user.Role != 1 {
		t.Errorf("Expected role 1, got %d", user.Role)
	}
}
