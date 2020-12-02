package users

import (
	"fmt"

	"github.com/nytro04/bookstore_users_api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewBadRequestError((fmt.Sprintf("user %d not found", user.Id)))
	}
	return nil
}

func (user User) Save() *errors.RestErr {
	return nil
}
