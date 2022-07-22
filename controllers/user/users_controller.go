package user

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// 	Todo:Handle errors
	// 	return
	// }
	// if err = json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	// 	Todo:Handle errors
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(http.StatusBadRequest, restErr)
		// 	todo: handle json errors
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// 	todo : handle user creation errors
		return
	}
	fmt.Println(result)
	fmt.Println(user)

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
