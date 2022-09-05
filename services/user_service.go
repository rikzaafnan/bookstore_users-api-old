package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/crypto_utils"
	"bookstore_users-api/utils/errors"
)

var (
	UserService UserServiceInterface = &userService{}
)

type userService struct {
}

type UserServiceInterface interface {
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	GetUser(userID int64) (*users.User, *errors.RestErr)
	UpdateUser(user users.User, userID int64) (*users.User, *errors.RestErr)
	DeleteUser(userID int64) *errors.RestErr
	Search(status string) (users.Users, *errors.RestErr)
	FindAll() (users.Users, *errors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {

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

	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) GetUser(userID int64) (*users.User, *errors.RestErr) {

	result := &users.User{ID: userID}

	if err := result.Get(userID); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *userService) UpdateUser(user users.User, userID int64) (*users.User, *errors.RestErr) {

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

	currentUser, err := s.GetUser(userID)
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

func (s *userService) DeleteUser(userID int64) *errors.RestErr {

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

func (s *userService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	users, err := dao.FindByStatus(status)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) FindAll() (users.Users, *errors.RestErr) {
	dao := &users.User{}
	users, err := dao.FindAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
