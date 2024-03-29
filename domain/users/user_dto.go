package users

import (
	"strings"

	"github.com/nytro04/bookstore_users_api/utils/errors"
)

const (
	StatusActive = "active"
)

// the dto consist of the data that is going to be transferred from the persistence layer to the application and backwards

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.NewBadRequestError("Please provide a valid email address")
	}

	user.Password = strings.TrimSpace(user.Password)

	if user.Password == "" {
		return errors.NewBadRequestError("User password field is required")
	}
	return nil
}
