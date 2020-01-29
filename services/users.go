package services

import (
	"github.com/apiotrowski312/isOnline-utils-go/crypto_utils"
	"github.com/apiotrowski312/isOnline-utils-go/date_utils"
	"github.com/apiotrowski312/isOnline-utils-go/rest_errors"
	"github.com/apiotrowski312/isOnline-users-api/domain/users"
)

var (
	UsersService userServiceInterace = &userService{}
)

type userService struct{}

type userServiceInterace interface {
	CreateUser(users.User) (*users.User, rest_errors.RestErr)
	GetUser(int64) (*users.User, rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	DeleteUser(int64) rest_errors.RestErr
	SearchUser(string) (users.Users, rest_errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, rest_errors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) GetUser(userId int64) (*users.User, rest_errors.RestErr) {
	results := &users.User{Id: userId}

	if err := results.Get(); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, rest_errors.RestErr) {
	current, err := UsersService.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
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
		if user.Password != "" {
			current.Password = crypto_utils.GetMd5(user.Password)
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Password = crypto_utils.GetMd5(user.Password)
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *userService) DeleteUser(userId int64) rest_errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *userService) SearchUser(status string) (users.Users, rest_errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *userService) LoginUser(request users.LoginRequest) (*users.User, rest_errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
