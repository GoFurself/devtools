package userops

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type hasherBCrypt struct {
	cost int
}

// GenerateFromPassword does not accept passwords longer than 72 bytes, which is the longest password bcrypt will operate on
func NewHasherBCrypt(cost int) *hasherBCrypt {

	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		cost = bcrypt.DefaultCost
	}

	return &hasherBCrypt{cost: cost}
}

func (h *hasherBCrypt) HashPassword(password string) (string, error) {

	if len(password) > 72 {
		return "", errors.New("password too long, max 72 bytes allowed by bcrypt")
	}

	if len(password) == 0 {
		return "", errors.New("password empty")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (h *hasherBCrypt) CheckPassword(password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
