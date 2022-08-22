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

func UpdateUser(user users.User, userID int64) (*users.User, *errors.RestErr) {

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

	currentUser, err := GetUser(userID)
	if err != nil {
		return nil, err
	}

	currentUser.Email = user.Email
	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName

	if err := currentUser.Update(userID); err != nil {
		return nil, err
	}

	return currentUser, nil
}

func DeleteUser(userID int64) *errors.RestErr {

	result := &users.User{ID: userID}

	if err := result.Get(userID); err != nil {
		return err
	}

	rowAffected, err := result.Delete(userID)
	if err != nil {
		return err
	}

	if rowAffected <= 0 {
		return errors.NewBadRequestError("no row affected")
	}

	return nil

}
