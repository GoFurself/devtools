package userops

import (
	"errors"
	"time"
)

type UserService struct {
	Repo   UserRepositoryInterface
	Hasher Hasher
}

func newUserService(repo UserRepositoryInterface, hasher Hasher) *UserService {
	return &UserService{
		Repo:   repo,
		Hasher: hasher,
	}
}

func (us *UserService) CreateUser(user *user) error {

	user.Created = time.Now()
	user.Enabled = true
	user.Metadata = "{}"

	hashedPassword, err := us.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return us.Repo.CreateUser(user)
}

func (us *UserService) GetUserByID(id uint64) (*user, error) {
	return us.Repo.GetUserByID(id)
}

func (us *UserService) GetUserByEmail(email string) (*user, error) {
	return us.Repo.GetUserByEmail(email)
}

func (us *UserService) GetUsers(filter string) ([]*user, error) {
	return us.Repo.GetUsers()
}

func (us *UserService) UpdateUser(user *user) error {
	return us.Repo.UpdateUser(user)
}

func (us *UserService) DeleteUser(user *user) error {
	return us.Repo.DeleteUser(user)
}

func (us *UserService) EnableUser(user *user) error {
	return us.Repo.EnableUser(user)
}

func (us *UserService) DisableUser(user *user) error {
	return us.Repo.DisableUser(user)
}

func (us *UserService) AuthenticateUser(user *user, password string) (bool, error) {
	valid, err := us.ValidatePassword(password, user.Password)
	if err != nil {
		return false, err
	}
	if !valid {
		return false, errors.New("invalid password")
	}
	return true, nil
}

func (us *UserService) ValidatePassword(password string, hash string) (bool, error) {
	valid, err := us.Hasher.CheckPassword(password, hash)
	if err != nil {
		return false, err
	}
	return valid, nil
}

func (us *UserService) HashPassword(password string) (string, error) {
	hash, err := us.Hasher.HashPassword(password)
	if err != nil {
		return "", err
	}
	return hash, nil
}
