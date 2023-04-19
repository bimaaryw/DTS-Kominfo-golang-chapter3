package helper

import (
	"chapter3-mygram/model"

	"github.com/go-playground/validator/v10"
)

func UserRegisterValidator(requestData model.UserRegisterRequest) []error {
	validate = validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func UserLoginValidator(requestData model.UserLoginRequest) []error {
	validate = validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func PhotoCreateValidator(requestData model.PhotoCreateRequest) []error {
	validate = validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func PhotoUpdateValidator(requestData model.PhotoUpdateRequest) []error {
	validate = validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func SocialCreateValidator(requestData model.SocialMediaCreateRequest) []error {
	validate = validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func SocialUpdateValidator(requestData model.SocialMediaUpdateRequest) []error {
	validate = validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func CommentCreateValidator(requestData model.CommentCreateRequest) []error {
	validate = validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func CommentUpdateValidator(requestData model.CommentUpdateRequest) []error {
	validate = validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}
