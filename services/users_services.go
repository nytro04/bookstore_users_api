package services

import (
	"github.com/nytro04/bookstore_users_api/domain/users"
	"github.com/nytro04/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return &user, nil
}
