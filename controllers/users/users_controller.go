package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/nytro04/bookstore_users_api/services"

	"github.com/nytro04/bookstore_users_api/domain/users"
	"github.com/nytro04/bookstore_users_api/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Println(user)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, getUserErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if getUserErr != nil {
		err := errors.NewBadRequestError("user Id shoud be a number")
		c.JSON(err.Status, err)
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}
	c.JSON(http.StatusOK, user)
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me")
// }
