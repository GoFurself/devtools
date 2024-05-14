package userops

import (
	"errors"
)

type UserRepositoryStatic struct {
	Users map[uint64]*user
}

func newUserRepositoryStatic() *UserRepositoryStatic {
	return &UserRepositoryStatic{Users: make(map[uint64]*user)}
}

func (ur *UserRepositoryStatic) CreateUser(user *user) error {
	ur.Users[user.ID] = user
	return nil
}

func (ur *UserRepositoryStatic) GetUserByID(id uint64) (*user, error) {
	user, ok := ur.Users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (ur *UserRepositoryStatic) GetUserByEmail(email string) (*user, error) {
	for _, user := range ur.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (ur *UserRepositoryStatic) GetUsers(filter string) ([]*user, error) {
	var users []*user
	for _, user := range ur.Users {
		if user.FirstName == filter || user.LastName == filter || user.Email == filter {
			users = append(users, user)
		}
	}
	return users, nil
}

func (ur *UserRepositoryStatic) UpdateUser(user *user) error {
	ur.Users[user.ID] = user
	ur.Users[user.ID] = user
	ur.Users[user.ID] = user
	ur.Users[user.ID] = user
	ur.Users[user.ID] = user
	ur.Users[user.ID] = user
	return nil
}

func (ur *UserRepositoryStatic) DeleteUser(user *user) error {
	delete(ur.Users, user.ID)
	return nil
}

func (ur *UserRepositoryStatic) DisableUser(user *user) error {
	user.Enabled = false
	return nil
}
