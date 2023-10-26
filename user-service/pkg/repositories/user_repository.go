package repositories

import (
	"github.com/nghiack7/micro-expense-manager/user-service/pkg/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateNewUser(user *models.User) error
	UpdateUser(user *models.User) error
	GetUserByID(id uint) (user *models.User, err error)
	GetUserByUserNameOrEmail(name string) (user *models.User, err error)
	DeleteUserByID(id uint) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (r *userRepositoryImpl) CreateNewUser(user *models.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImpl) UpdateUser(user *models.User) error {
	var oldUser models.User
	err := r.db.Where("id=?", user.ID).First(&oldUser).Error
	if err != nil {
		return err
	}
	user.ID = oldUser.ID
	err = r.db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImpl) GetUserByID(id uint) (*models.User, error) {
	var user *models.User

	err := r.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepositoryImpl) GetUserByUserNameOrEmail(name string) (user *models.User, err error) {
	err = r.db.Where("username=? or email=?", name, name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) DeleteUserByID(id uint) error {
	err := r.db.Where("id=?", id).Delete(&models.User{}).Error
	if err != nil {
		return err
	}

	return nil
}
