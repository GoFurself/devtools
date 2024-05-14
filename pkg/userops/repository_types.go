package userops

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
