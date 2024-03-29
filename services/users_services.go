package services

import (
	"github.com/nytro04/bookstore_users_api/domain/users"
	"github.com/nytro04/bookstore_users_api/utils/bcrypt_utils"
	"github.com/nytro04/bookstore_users_api/utils/date_utils"
	"github.com/nytro04/bookstore_users_api/utils/errors"
)

var (
	UserService userServiceInferface = &userService{}
)

type userService struct {
}

type userServiceInferface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	var errPassword error
	user.Password, errPassword = bcrypt_utils.HashedPassword(user.Password)
	if errPassword != nil {
		return nil, errors.NewBadRequestError("failed to hash password")
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}


	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *userService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	user := &users.User{}
	return user.FindByStatus(status)
}
