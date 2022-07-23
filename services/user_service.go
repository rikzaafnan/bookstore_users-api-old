package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	// user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	// if user.Email == "" {
	// 	return nil, errors.NewBadRequestError("invalid email address")
	// }

	// cara function
	// if err := users.Validate(&user); err != nil {
	// 	return nil, err
	// }

	// cara method
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {

	result := &users.User{ID: userID}

	if err := result.Get(userID); err != nil {
		return nil, err
	}

	return result, nil

}
