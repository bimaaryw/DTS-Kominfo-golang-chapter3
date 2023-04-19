package service

import (
	"errors"
	"fmt"

	"chapter3-mygram/helper"
	"chapter3-mygram/model"
	"chapter3-mygram/repository"
)

type UserService interface {
	Register(userReqData model.UserRegisterRequest) (*model.UserRegisterResponse, error)
	Login(userReqData model.UserLoginRequest) (*string, error)
	GetProfile(userID string) (model.User, error)
}

type UserServiceIml struct {
	userRepository  repository.UserRepository
	photoRepository repository.PhotoRepository
	socalRepository repository.SocialMediaRepository
}

func NewUserService(userRepo repository.UserRepository, photoRepo repository.PhotoRepository, socialRepo repository.SocialMediaRepository) UserService {
	return &UserServiceIml{
		userRepository:  userRepo,
		photoRepository: photoRepo,
		socalRepository: socialRepo,
	}
}

func (s *UserServiceIml) Register(userReqData model.UserRegisterRequest) (*model.UserRegisterResponse, error) {
	userID := helper.GenerateID()
	hashedPassword, err := helper.Hash(userReqData.Password)
	if err != nil {
		return nil, err
	}

	newUser := model.User{
		UserID:   userID,
		Username: userReqData.Username,
		Email:    userReqData.Email,
		Password: hashedPassword,
		Age:      userReqData.Age,
	}

	err = s.userRepository.Create(newUser)
	if err != nil {
		return nil, err
	}

	return &model.UserRegisterResponse{
		UserID:   newUser.UserID,
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
		Age:      newUser.Age,
	}, nil
}

func (s *UserServiceIml) Login(userReqData model.UserLoginRequest) (*string, error) {
	userResponse, err := s.userRepository.FindByUsername(userReqData.Username)
	if err != nil {
		return nil, err
	}

	isMatch := helper.PasswordIsMatch(userReqData.Password, userResponse.Password)
	if isMatch == false {
		return nil, errors.New(fmt.Sprintf("Invalid username or password"))
	}

	token, err := helper.GenerateToken(userResponse)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *UserServiceIml) GetProfile(userID string) (model.User, error) {
	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return model.User{}, err
	}

	photos, err := s.photoRepository.FindByUserID(userID)
	if err != nil {
		return model.User{}, err
	}

	socials, err := s.socalRepository.FindByUserID(userID)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		UserID:      userID,
		Username:    user.Username,
		Email:       user.Email,
		Age:         user.Age,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Photos:      photos,
		SocialMedia: socials,
	}, nil
}
