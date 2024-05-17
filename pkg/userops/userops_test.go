package userops_test

import (
	"os"
	"testing"

	"github.com/GoFurself/devtools/pkg/userops"
)

// TestUseropsServiceWithEmailAndPassword tests the userops service with email and password
// It creates a new user, checks if the user was created and if the password is valid
func TestUseropsServiceWithEmailAndPassword(t *testing.T) {

	email := "test@test.if"
	dataSourceName := "test_1.db"
	cost := 10

	// Create service
	service, err := userops.NewUseropsService(userops.SQLite, userops.WithDataSourceName(dataSourceName), userops.WithHasher(userops.NewHasherBCrypt(cost)))
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	// Create user
	user := userops.NewUser(email, "password")
	err = service.CreateUser(user)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	// Check if id was set
	if user.ID == 0 {
		t.Errorf("Expected id to be set, got 0")
	}

	// Check if user was created: excists in db
	user, err = service.GetUserByEmail(email)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	// Check if user was created: email is correct
	if user.Email != email {
		t.Errorf("Expected email %s, got %s", email, user.Email)
	}

	// Check if user was created: password is correct
	valid, err := service.AuthenticateUser(user, "password")
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if !valid {
		t.Errorf("Expected valid password, got invalid")
	}

	// Update user
	user.Email = "test2@test@.if"
	user.Password = "password2"
	err = service.UpdateUser(user)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	// Check if user was updated: use id to get user
	user, err = service.GetUserByID(user.ID)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	// Check if user was updated: email is correct
	if user.Email != "test2@test@.if" {
		t.Errorf("Expected email test2@test@.if, got %s", user.Email)
	}

	// Check if user was updated: password is correct
	valid, err = service.AuthenticateUser(user, "password2")
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if !valid {
		t.Errorf("Expected valid password, got invalid")
	}

	// User should be enabled by default
	if !user.Enabled {
		t.Errorf("Expected user to be enabled, got disabled")
	}

	// Disable user
	err = service.DisableUser(user)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	// Enable user
	err = service.EnableUser(user)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	// Delete user
	err = service.DeleteUser(user)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	// Check if user was deleted
	user, err = service.GetUserByEmail(email)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if user != nil {
		t.Errorf("Expected user to be deleted, got user")
	}

	// Remove the database file
	if os.Remove(dataSourceName) != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}
