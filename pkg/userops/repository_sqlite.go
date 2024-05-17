package userops

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
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

type UserRepositorySQLite struct {
	dal                *UseropsDBSQLite
	preparedStatements map[string]*sql.Stmt
}

func newUserRepositorySQLite(dal *UseropsDBSQLite) (UserRepositoryInterface, error) {
	repo := &UserRepositorySQLite{
		dal:                dal,
		preparedStatements: make(map[string]*sql.Stmt),
	}

	_, err := dal.DB().Exec(`CREATE TABLE IF NOT EXISTS Users (
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

	rows, err := dal.DB().Query(`SELECT name FROM sqlite_master WHERE type='table' AND name='Users';`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	statements := map[string]string{
		"create":     `INSERT INTO Users (FirstName, LastName, Email, Password, Role, Created, LastLogin, Enabled, Metadata) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		"getByID":    `SELECT ID, FirstName, LastName, Email, Password, Role, Created, LastLogin, Enabled, Metadata FROM Users WHERE ID = ?;`,
		"getByEmail": `SELECT ID, FirstName, LastName, Email, Password, Role, Created, LastLogin, Enabled, Metadata FROM Users WHERE Email = ?;`,
		"getAll":     `SELECT ID, FirstName, LastName, Email, Password, Role, Created, LastLogin, Enabled, Metadata FROM Users;`,
		"update":     `UPDATE Users SET FirstName = ?, LastName = ?, Email = ?, Password = ?, Role = ?, Created = ?, LastLogin = ?, Enabled = ?, Metadata = ? WHERE ID = ?;`,
		"delete":     `DELETE FROM Users WHERE ID = ?;`,
		"enable":     `UPDATE Users SET Enabled = 1 WHERE ID = ?;`,
		"disable":    `UPDATE Users SET Enabled = 0 WHERE ID = ?;`,
	}

	for key, query := range statements {
		stmt, err := dal.DB().Prepare(query)
		if err != nil {
			dal.Close()
			return nil, fmt.Errorf("error preparing %s statement: %w", key, err)
		}
		repo.preparedStatements[key] = stmt
	}

	return repo, nil
}

func (r *UserRepositorySQLite) CreateUser(user *user) error {
	res, err := r.preparedStatements["create"].Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Role, user.Created, user.LastLogin, user.Enabled, user.Metadata)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = uint64(id)
	return nil
}

func (r *UserRepositorySQLite) GetUserByID(id uint64) (*user, error) {
	res, err := r.preparedStatements["getByID"].Query(id)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	if res.Next() {
		user := &user{}
		err = res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role, &user.Created, &user.LastLogin, &user.Enabled, &user.Metadata)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

// GetUserByEmail retrieves a user by email
// It returns the user or nil if the user does not exist
// Errors are returned if there are issues with the database
func (r *UserRepositorySQLite) GetUserByEmail(email string) (*user, error) {
	res, err := r.preparedStatements["getByEmail"].Query(email)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	if res.Next() {
		user := &user{}
		err = res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role, &user.Created, &user.LastLogin, &user.Enabled, &user.Metadata)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

func (r *UserRepositorySQLite) GetUsers() ([]*user, error) {
	res, err := r.preparedStatements["getAll"].Query()
	if err != nil {
		return nil, err
	}
	defer res.Close()

	users := []*user{}
	for res.Next() {
		user := &user{}
		err = res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role, &user.Created, &user.LastLogin, &user.Enabled, &user.Metadata)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepositorySQLite) UpdateUser(user *user) error {
	res, err := r.preparedStatements["update"].Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Role, user.Created, user.LastLogin, user.Enabled, user.Metadata, user.ID)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

func (r *UserRepositorySQLite) DeleteUser(user *user) error {
	res, err := r.preparedStatements["delete"].Exec(user.ID)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("no rows deleted")
	}
	return nil
}

func (r *UserRepositorySQLite) EnableUser(user *user) error {
	res, err := r.preparedStatements["enable"].Exec(user.ID)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

func (r *UserRepositorySQLite) DisableUser(user *user) error {
	res, err := r.preparedStatements["disable"].Exec(user.ID)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("no rows updated")
	}
	return nil
}
