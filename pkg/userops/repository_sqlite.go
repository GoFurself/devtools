package userops

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type UserRepositorySQLite struct {
	dal                *UseropsDBSQLite
	preparedStatements map[string]*sql.Stmt
}

func newUserRepositorySQLite(dal *UseropsDBSQLite) (UserRepositoryInterface, error) {
	repo := &UserRepositorySQLite{
		dal:                dal,
		preparedStatements: make(map[string]*sql.Stmt),
	}

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
