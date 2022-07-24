package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"bookstore_users-api/utils/date_utils"
	"bookstore_users-api/utils/errors"
	"fmt"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get(userID int64) *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[userID]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", userID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]

	if current != nil {

		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s allready exists", user.Email))
		}

		return errors.NewBadRequestError(fmt.Sprintf("user %d allready exists", user.ID))
	}

	user.DateCreated = date_utils.GetNowString()

	return nil
}
