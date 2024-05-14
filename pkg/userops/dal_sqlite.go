package userops

import (
	"database/sql"
	"time"
)

type UseropsDBSQLite struct {
	db             *sql.DB
	dataSourceName string
}

func newUseropsDBSQLite(dataSourceName string) (*UseropsDBSQLite, error) {

	sqlDB, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	_, err = sqlDB.Exec(`CREATE TABLE IF NOT EXISTS Users (
		ID INTEGER PRIMARY KEY, -- SQLite uses INTEGER for 64-bit auto-increment IDs
		FirstName TEXT,
		LastName TEXT,
		Email TEXT UNIQUE, -- Assuming email should be unique
		Password TEXT,
		Role INTEGER, -- Using INTEGER to store the uint8
		Created DATETIME, -- Using DATETIME for the time.Time fields
		LastLogin DATETIME,
		Enabled BOOLEAN, -- BOOLEAN in SQLite is stored as INTEGER (0 or 1)
		Metadata TEXT -- Storing metadata as TEXT, assuming JSON or similar format	
	);`)
	if err != nil {
		return nil, err
	}

	rows, err := sqlDB.Query(`SELECT name FROM sqlite_master WHERE type='table' AND name='Users';`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return &UseropsDBSQLite{
		db:             sqlDB,
		dataSourceName: dataSourceName,
	}, nil
}

func (u *UseropsDBSQLite) DB() *sql.DB {
	return u.db
}

func (s *UseropsDBSQLite) Close() error {
	return s.db.Close()
}
