package user

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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

	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		// 	todo : handle user creation errors
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("x-Public") == "true"))
}

func GetUser(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError(fmt.Sprintf("invalid user id %d", userID))
		c.JSON(err.Status, err)
		return
	}

	result, saveErr := services.UserService.GetUser(userID)
	if saveErr != nil {
		// 	todo : handle user creation errors
		c.JSON(http.StatusBadRequest, saveErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("x-Public") == "true"))
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

	result, saveErr := services.UserService.UpdateUser(user, userID)
	if saveErr != nil {
		// 	todo : handle user creation errors
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("x-Public") == "true"))
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

	// c.JSON(http.StatusCreated, userID)
}

func DeleteUser(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError(fmt.Sprintf("invalid user id %d", userID))
		c.JSON(err.Status, err)
		return
	}

	saveErr := services.UserService.DeleteUser(userID)
	if saveErr != nil {

		// 	todo : handle user creation errors
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {

	status := c.Query("status")

	log.Println(status)

	if status != "" {
		users, err := services.UserService.Search(status)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}

		// result := make([]interface{}, len(users))
		// for index, user := range users {
		// 	result[index] = user.Marshall(c.GetHeader("x-Public") == "true")
		// }

		c.JSON(http.StatusOK, users.Marshall(c.GetHeader("x-Public") == "true"))
		return
	} else {
		users, err := services.UserService.FindAll()
		if err != nil {
			c.JSON(err.Status, err)
			return
		}

		c.JSON(http.StatusOK, users.Marshall(c.GetHeader("x-Public") == "true"))
		return
	}

}
