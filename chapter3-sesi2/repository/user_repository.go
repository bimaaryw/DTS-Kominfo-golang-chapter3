package repository

import (
	"errors"
	model "chapter3-sesi2/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repository *UserRepository) CreateUser(user model.User) (*model.User, error) {
	newUser := model.User{
		UserID:   user.UserID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}

	err := repository.DB.Create(&newUser).Error
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (repository *UserRepository) UserCheck(user model.UserLoginRequest) (*model.User, error) {
	userResult := model.User{}

	err := repository.DB.Debug().Where("email = ?", user.Email).Take(&userResult).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}

	return &userResult, nil
}
