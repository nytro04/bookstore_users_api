package users

import (
	"net/http"

	"github.com/nytro04/bookstore_users_api/utils/errors"

	"github.com/nytro04/bookstore_users_api/services"

	"github.com/nytro04/bookstore_users_api/domain/users"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user users.User

	//replaced with c.ShouldBindJSON()
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	//TODO: handle error
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	//TODO: handle error
	//	return
	//}      l

	// this handles reading and unmarshalling of request body
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
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
	c.String(http.StatusNotImplemented, "implement me")

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")

}
