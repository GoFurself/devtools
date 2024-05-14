package userops

import "database/sql"

type UseropsDB interface {
	DB() *sql.DB
	Close() error
}
