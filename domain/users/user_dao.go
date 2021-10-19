package users

import (
	"fmt"

	"github.com/nytro04/bookstore_users_api/utils/mysql_utils"

	"github.com/nytro04/bookstore_users_api/datasources/mysql/users_db"

	"github.com/nytro04/bookstore_users_api/utils/errors"
)

// Consist of the entire logic to persist and retrieve the user from the database

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

//we use *User i.e pointer to the user, cos we want to modify
// the user itself and not a copy (User) of the user
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare((queryGetUser))
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, &user.Status, &user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr  {

	stmt, err := users_db.UsersDB.Prepare(queryUpdateUser)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
			if err != nil {
				return mysql_utils.ParseError(err)
			}

			return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryDeleteUser)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		if _, err = stmt.Exec(user.Id); err != nil {
			return mysql_utils.ParseError(err)
		}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr)  {
	stmt, err := users_db.UsersDB.Prepare(queryFindUserByStatus)
	if err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		rows, err := stmt.Query(status)
		if err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		defer rows.Close()

		results := make([]User, 0)
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.Status); err != nil {
	return nil, mysql_utils.ParseError(err)
			}
			results = append(results, user)
		}

		if len(results) == 0 {
			return nil, errors.NewNotFoundError(fmt.Sprintf("No users matching status %s", status))
		}

		return results, nil
}
