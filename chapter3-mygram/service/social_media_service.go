package service

import (
	"errors"
	"time"

	"chapter3-mygram/helper"
	"chapter3-mygram/model"
	"chapter3-mygram/repository"
)

type SocialService interface {
	Create(socialReqData model.SocialMediaCreateRequest, userID string) (*model.SocialMediaResponse, error)
	GetAll() ([]model.SocialMediaResponse, error)
	GetOne(socialID string) (model.SocialMediaResponse, error)
	UpdateSocialMedia(socialReqData model.SocialMediaUpdateRequest, userID string, socialID string) (*model.SocialMediaResponse, error)
	Delete(socialID string, userID string) (model.SocialMediaResponse, error)
}

type SocialServiceIml struct {
	socialRepository repository.SocialMediaRepository
}

func NewSocialService(socialRepo repository.SocialMediaRepository) SocialService {
	return &SocialServiceIml{
		socialRepository: socialRepo,
	}
}

func (s *SocialServiceIml) Create(socialReqData model.SocialMediaCreateRequest, userID string) (*model.SocialMediaResponse, error) {
	socialID := helper.GenerateID()
	newSocial := model.SocialMedia{
		ID:             socialID,
		Name:           socialReqData.Name,
		SocialMediaURL: socialReqData.SocialMediaURL,
		UserID:         userID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := s.socialRepository.Create(newSocial)
	if err != nil {
		return nil, err
	}

	return &model.SocialMediaResponse{
		ID:             newSocial.ID,
		Name:           newSocial.Name,
		SocialMediaURL: newSocial.SocialMediaURL,
		UserID:         newSocial.UserID,
		CreatedAt:      newSocial.CreatedAt,
		UpdatedAt:      newSocial.UpdatedAt,
	}, nil
}

func (s *SocialServiceIml) GetAll() ([]model.SocialMediaResponse, error) {
	photosResult, err := s.socialRepository.FindAll()
	if err != nil {
		return []model.SocialMediaResponse{}, err
	}

	socialsResponse := []model.SocialMediaResponse{}
	for _, socialRes := range photosResult {
		socialsResponse = append(socialsResponse, model.SocialMediaResponse(socialRes))
	}

	return socialsResponse, nil
}

func (s *SocialServiceIml) GetOne(socialID string) (model.SocialMediaResponse, error) {
	socialsResult, err := s.socialRepository.FindByID(socialID)
	if err != nil {
		return model.SocialMediaResponse{}, err
	}

	return model.SocialMediaResponse(socialsResult), nil
}

func (s *SocialServiceIml) UpdateSocialMedia(socialReqData model.SocialMediaUpdateRequest, userID string, socialID string) (*model.SocialMediaResponse, error) {
	findSocialMediaResponse, err := s.socialRepository.FindByID(socialID)
	if err != nil {
		return nil, err
	}

	if userID != findSocialMediaResponse.UserID {
		return nil, errors.New("Unauthorized")
	}

	updatedSocialReq := model.SocialMedia{
		ID:             socialID,
		Name:           socialReqData.Name,
		SocialMediaURL: socialReqData.SocialMediaURL,
		UserID:         userID,
		UpdatedAt:      time.Now(),
	}

	err = s.socialRepository.Update(updatedSocialReq)
	if err != nil {
		return nil, err
	}

	return &model.SocialMediaResponse{
		ID: socialID,
	}, nil
}

func (s *SocialServiceIml) Delete(socialID string, userID string) (model.SocialMediaResponse, error) {
	findSocialResponse, err := s.socialRepository.FindByID(socialID)
	if err != nil {
		return model.SocialMediaResponse{}, err
	}

	if userID != findSocialResponse.UserID {
		return model.SocialMediaResponse{}, errors.New("Unauthorized")
	}

	err = s.socialRepository.Delete(model.SocialMedia{ID: socialID})
	if err != nil {
		return model.SocialMediaResponse{}, err
	}

	return model.SocialMediaResponse{
		ID: socialID,
	}, nil
}
