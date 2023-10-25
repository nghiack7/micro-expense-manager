package services

import (
	"github.com/nghiack7/micro-expense-manager/user-service/pkg/models"
	"github.com/nghiack7/micro-expense-manager/user-service/pkg/repositories"
)

type UserService interface {
	FindUserById(uint) (*models.User, error)
	FindUserByEmail(string) (*models.User, error)
	UpdateUser(*models.User) (*models.User, error)
	CreateUser(*models.User) (*models.User, error)
	DeleteUserByID(uint) error
}

type userServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewUserServices() UserService {
	return &userServiceImpl{}
}
func (u *userServiceImpl) FindUserById(id uint) (*models.User, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userServiceImpl) FindUserByEmail(name string) (*models.User, error) {
	user, err := u.userRepo.GetUserByUserNameOrEmail(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userServiceImpl) UpdateUser(user *models.User) (*models.User, error) {
	err := u.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userServiceImpl) CreateUser(user *models.User) (*models.User, error) {
	err := u.userRepo.CreateNewUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userServiceImpl) DeleteUserByID(id uint) error {
	return u.userRepo.DeleteUserByID(id)

}
