package userops

import (
	"testing"
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
