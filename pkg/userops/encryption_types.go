package userops

type Hasher interface {
	HashPassword(password string) (string, error)
	CheckPassword(password string, hash string) (bool, error)
}
