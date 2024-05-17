package userops

import "time"

type user struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string
	Role      UserRole  `json:"role"`
	Created   time.Time `json:"created"`
	LastLogin time.Time `json:"last_login"`
	Enabled   bool      `json:"enabled"`
	Metadata  string    `json:"metadata"`
}

type UserRole uint8

const (
	RoleUser UserRole = iota
	RoleAdmin
)

func NewUser(email string, password string, options ...UserOptions) *user {
	user := &user{
		Email:    email,
		Password: password,
	}
	for _, option := range options {
		option(user)
	}
	return user
}

type UserOptions func(*user)

func WithFirstName(firstName string) UserOptions {
	return func(user *user) {
		user.FirstName = firstName
	}
}

func WithLastName(lastName string) UserOptions {
	return func(user *user) {
		user.LastName = lastName
	}
}

func WithRole(role UserRole) UserOptions {
	return func(user *user) {
		user.Role = role
	}
}
