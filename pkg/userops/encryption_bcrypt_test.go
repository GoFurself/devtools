package userops

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcryptHasher(t *testing.T) {

	hasher := NewHasherBCrypt(10)
	password := "password"

	hash, err := hasher.HashPassword(password)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if hash == "" {
		t.Errorf("Expected hash, got empty string")
	}
	if len(hash) < 60 {
		t.Errorf("Expected hash to be at least 60 characters, got %d", len(hash))
	}
}

func TestBcryptChecker(t *testing.T) {

	hasher := NewHasherBCrypt(10)

	password := "password"
	hash, err := hasher.HashPassword(password)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	valid, err := hasher.CheckPassword(password, hash)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if !valid {
		t.Errorf("Expected valid password, got invalid")
	}
}

func TestBcryptCheckerInvalid(t *testing.T) {

	hasher := NewHasherBCrypt(10)

	password := "password"
	hash, err := hasher.HashPassword(password)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	valid, err := hasher.CheckPassword("invalid", hash)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if valid {
		t.Errorf("Expected invalid password, got valid")
	}
}

// BCrypt does not accept passwords longer than 72 bytes
func TestToLongPassword(t *testing.T) {

	hasher := NewHasherBCrypt(10)
	password := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" // 73 bytes
	_, err := hasher.HashPassword(password)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	password = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" // 72 bytes
	_, err = hasher.HashPassword(password)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	password = "" // 0 bytes
	_, err = hasher.HashPassword(password)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestCosts(t *testing.T) {

	hasher := NewHasherBCrypt(1)
	// cost should be set to default
	if hasher.cost != bcrypt.DefaultCost {
		t.Errorf("Expected cost 10, got %d", hasher.cost)
	}

	hasher = NewHasherBCrypt(100)
	// cost should be set to max
	if hasher.cost != bcrypt.DefaultCost {
		t.Errorf("Expected cost 10, got %d", hasher.cost)
	}
}
