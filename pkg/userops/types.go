package userops

import "database/sql"

type UseropsDB interface {
	DB() *sql.DB
	Close() error
}

type Hasher interface {
	HashPassword(password string) (string, error)
	CheckPassword(password string, hash string) (bool, error)
}

type UserRepositoryInterface interface {
	CreateUser(user *user) error
	GetUserByID(id uint64) (*user, error)
	GetUserByEmail(email string) (*user, error)
	GetUsers() ([]*user, error)
	UpdateUser(user *user) error
	DeleteUser(user *user) error
	DisableUser(user *user) error
	EnableUser(user *user) error
}

type UseropsServiceInterface interface {
	CreateUser(user *user) error
	GetUserByID(id uint64) (*user, error)
	GetUserByEmail(email string) (*user, error)
	GetUsers(filter string) ([]*user, error)
	UpdateUser(user *user) error
	DeleteUser(user *user) error
	DisableUser(user *user) error
	EnableUser(user *user) error
	AuthenticateUser(user *user, password string) (bool, error)
	HashPassword(password string) (string, error)
	ValidatePassword(password string, hash string) (bool, error)
}
