package userops

type UserServiceInterface interface {
	CreateUser(user *user) error
	GetUserByID(id uint64) (*user, error)
	GetUserByEmail(email string) (*user, error)
	GetUsers(filter string) ([]*user, error)
	UpdateUser(user *user) error
	DeleteUser(user *user) error
	DisableUser(user *user) error
	AuthenticateUser(user *user, password string) (bool, error)
	HashPassword(password string) (string, error)
	ValidatePassword(password string, hash string) (bool, error)
}
