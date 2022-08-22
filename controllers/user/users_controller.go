package user

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError(fmt.Sprintf("invalid user id %d", userID))
		c.JSON(err.Status, err)
		return
	}

	result, saveErr := services.GetUser(userID)
	if saveErr != nil {
		// 	todo : handle user creation errors
		c.JSON(http.StatusBadRequest, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func UpdateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError(fmt.Sprintf("invalid user id %d", userID))
		c.JSON(err.Status, err)
		return
	}

	result, saveErr := services.UpdateUser(user, userID)
	if saveErr != nil {
		// 	todo : handle user creation errors
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func PatchEmailUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError(fmt.Sprintf("invalid user id %d", userID))
		c.JSON(err.Status, err)
		return
	}

	// result, saveErr := services.UpdateUser(user, userID)
	// if saveErr != nil {
	// 	// 	todo : handle user creation errors
	// 	c.JSON(http.StatusBadRequest, saveErr)
	// 	return
	// }

	panic("implement patch user")

	c.JSON(http.StatusCreated, userID)
}

func DeleteUser(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError(fmt.Sprintf("invalid user id %d", userID))
		c.JSON(err.Status, err)
		return
	}

	saveErr := services.DeleteUser(userID)
	if saveErr != nil {

		// 	todo : handle user creation errors
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, "deleted")
}
