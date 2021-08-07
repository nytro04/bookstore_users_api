package users

import (
	"github.com/nytro04/bookstore_users_api/utils/mysql_utils"

	"github.com/nytro04/bookstore_users_api/datasources/mysql/users_db"

	"github.com/nytro04/bookstore_users_api/utils/date_utils"

	"github.com/nytro04/bookstore_users_api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?)"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created from users WHERE id=?"
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

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
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

	user.DateCreated = date_utils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
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
