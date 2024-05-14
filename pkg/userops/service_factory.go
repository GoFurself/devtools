package userops

import (
	"errors"
	"strings"
)

type RepoType int

const (
	SQLite RepoType = iota
)

type UserServiceFactoryOption func(*UserServiceFactoryOptions) error

type UserServiceFactoryOptions struct {
	RepoType       RepoType
	DataSourceName string
}

func WithDataSourceName(dataSourceName string) UserServiceFactoryOption {

	return func(o *UserServiceFactoryOptions) error {
		if len(strings.TrimSpace(dataSourceName)) == 0 || len(dataSourceName) > 100 {
			return errors.New("DataSourceName cannot be empty or over 100 characters long")
		}
		o.DataSourceName = dataSourceName
		return nil
	}
}

func UserServiceFactory(repoType RepoType, options ...UserServiceFactoryOption) (UserServiceInterface, error) {

	opts := UserServiceFactoryOptions{}
	for _, opt := range options {
		if err := opt(&opts); err != nil {
			return nil, err
		}
	}

	switch repoType {
	case SQLite:
		if opts.DataSourceName == "" {
			return nil, errors.New("missing DataSourceName for SQLite repo type")
		}

		db, err := newUseropsDBSQLite(opts.DataSourceName)
		if err != nil {
			return nil, err
		}

		repo, err := newUserRepositorySQLite(db)
		if err != nil {
			return nil, err
		}

		return newUserService(repo, NewHasherBCrypt(10)), nil

	default:
		return nil, errors.New("invalid user service type")
	}
}
