package services

import (
	"chapter3-sesi2/helpers"
	model "chapter3-sesi2/models"
	"chapter3-sesi2/repository"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (service *UserService) CreateNewUser(request model.UserRegisterRequest) (model.UserRegisterResponse, error) {
	userID := uuid.New()
	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	var role bool
	if request.Role == "admin" {
		role = true
	} else {
		role = false
	}

	user := model.User{
		UserID:   userID,
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
		Role:     role,
	}

	result, err := service.UserRepository.CreateUser(user)
	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	var roleString string
	if result.Role == true {
		roleString = "admin"
	} else {
		roleString = "user"
	}
	response := model.UserRegisterResponse{
		UserID: result.UserID,
		Name:   result.Name,
		Email:  result.Email,
		Role:   roleString,
	}

	return response, nil
}

func (service *UserService) Login(request model.UserLoginRequest) (string, error) {
	result, err := service.UserRepository.UserCheck(request)
	if err != nil {
		return "", err
	}

	isMatch := helpers.PasswordIsMatch(request.Password, result.Password)
	if isMatch == false {
		return "", errors.New(fmt.Sprintf("Invalid username or password"))
	}

	myClaim := helpers.MyClaims{
		User: &model.User{
			Email: result.Email,
			Role:  result.Role,
		},
	}
	jwtToken, err := helpers.GenerateToken(myClaim)

	return jwtToken, err
}
