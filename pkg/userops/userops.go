package userops

import (
	"errors"
	"strings"
)

type ServiceType int

const (
	SQLite ServiceType = iota
)

type UserServiceFactoryOption func(*UserServiceFactoryOptions) error

type UserServiceFactoryOptions struct {
	ServiceType    ServiceType
	DataSourceName string
	Hasher         Hasher
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

func WithHasher(h Hasher) UserServiceFactoryOption {

	return func(o *UserServiceFactoryOptions) error {
		if h == nil {
			return errors.New("Hasher cannot be nil")
		}
		o.Hasher = h
		return nil
	}
}

func NewUseropsService(st ServiceType, options ...UserServiceFactoryOption) (UseropsServiceInterface, error) {

	opts := UserServiceFactoryOptions{}
	for _, opt := range options {
		if err := opt(&opts); err != nil {
			return nil, err
		}
	}

	switch st {
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

		if opts.Hasher == nil {
			opts.Hasher = NewHasherBCrypt(10)
		}

		return newUseropsService(repo, opts.Hasher), nil

	default:
		return nil, errors.New("invalid user service type")
	}
}
