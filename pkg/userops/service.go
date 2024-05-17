package userops

import (
	"errors"
	"time"
)

type UseropsService struct {
	Repo   UserRepositoryInterface
	Hasher Hasher
}

func newUseropsService(repo UserRepositoryInterface, hasher Hasher) UseropsServiceInterface {
	return &UseropsService{
		Repo:   repo,
		Hasher: hasher,
	}
}

func (use UseropsService) CreateUser(user *user) error {

	user.Created = time.Now()
	user.Enabled = true
	user.Metadata = "{}"

	hashedPassword, err := use.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return use.Repo.CreateUser(user)
}

func (use UseropsService) GetUserByID(id uint64) (*user, error) {
	return use.Repo.GetUserByID(id)
}

func (use UseropsService) GetUserByEmail(email string) (*user, error) {
	return use.Repo.GetUserByEmail(email)
}

func (use UseropsService) GetUsers(filter string) ([]*user, error) {
	return use.Repo.GetUsers()
}

func (use UseropsService) UpdateUser(user *user) error {

	hashedPassword, err := use.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return use.Repo.UpdateUser(user)
}

func (use UseropsService) DeleteUser(user *user) error {
	return use.Repo.DeleteUser(user)
}

func (use UseropsService) EnableUser(user *user) error {
	return use.Repo.EnableUser(user)
}

func (use UseropsService) DisableUser(user *user) error {
	return use.Repo.DisableUser(user)
}

func (use UseropsService) AuthenticateUser(user *user, password string) (bool, error) {
	valid, err := use.ValidatePassword(password, user.Password)
	if err != nil {
		return false, err
	}
	if !valid {
		return false, errors.New("invalid password")
	}
	return true, nil
}

func (use UseropsService) ValidatePassword(password string, hash string) (bool, error) {
	valid, err := use.Hasher.CheckPassword(password, hash)
	if err != nil {
		return false, err
	}
	return valid, nil
}

func (use UseropsService) HashPassword(password string) (string, error) {
	hash, err := use.Hasher.HashPassword(password)
	if err != nil {
		return "", err
	}
	return hash, nil
}
