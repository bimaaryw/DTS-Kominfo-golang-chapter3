package service

import (
	"errors"
	"time"

	"chapter3-mygram/helper"
	"chapter3-mygram/model"
	"chapter3-mygram/repository"
)

type PhotoService interface {
	Create(photoReqData model.PhotoCreateRequest, userID string) (*model.PhotoCreateResponse, error)
	GetAll() ([]model.PhotoResponse, error)
	GetOne(photoID string) (model.PhotoResponse, error)
	UpdatePhoto(photoReqData model.PhotoUpdateRequest, userID string, photoID string) (*model.PhotoResponse, error)
	Delete(photoID string, userID string) (model.PhotoResponse, error)
}

type PhotoServiceIml struct {
	photoRepository   repository.PhotoRepository
	commentRepository repository.CommentRepository
}

func NewPhotoService(photoRepo repository.PhotoRepository, commentRepo repository.CommentRepository) PhotoService {
	return &PhotoServiceIml{
		photoRepository:   photoRepo,
		commentRepository: commentRepo,
	}
}

func (s *PhotoServiceIml) Create(photoReqData model.PhotoCreateRequest, userID string) (*model.PhotoCreateResponse, error) {
	photoID := helper.GenerateID()
	newPhoto := model.Photo{
		PhotoID:   photoID,
		Title:     photoReqData.Title,
		Caption:   photoReqData.Caption,
		PhotoURL:  photoReqData.PhotoURL,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.photoRepository.Create(newPhoto)
	if err != nil {
		return nil, err
	}

	return &model.PhotoCreateResponse{
		PhotoID:   newPhoto.PhotoID,
		Title:     newPhoto.Title,
		Caption:   newPhoto.Caption,
		PhotoURL:  newPhoto.PhotoURL,
		UserID:    newPhoto.UserID,
		CreatedAt: newPhoto.CreatedAt,
		UpdatedAt: newPhoto.UpdatedAt,
	}, nil
}

func (s *PhotoServiceIml) GetAll() ([]model.PhotoResponse, error) {
	photosResult, err := s.photoRepository.FindAll()
	if err != nil {
		return []model.PhotoResponse{}, err
	}

	photosResponse := []model.PhotoResponse{}
	for _, photoRes := range photosResult {
		photosResponse = append(photosResponse, model.PhotoResponse(photoRes))
	}

	return photosResponse, nil
}

func (s *PhotoServiceIml) GetOne(photoID string) (model.PhotoResponse, error) {
	photosResult, err := s.photoRepository.FindByID(photoID)
	if err != nil {
		return model.PhotoResponse{}, err
	}

	comments := []model.Comment{}
	commentsResponse, err := s.commentRepository.FindByPhotoID(photoID)
	for _, comment := range commentsResponse {
		comments = append(comments, model.Comment(comment))
	}
	if err != nil {
		return model.PhotoResponse{}, err
	}

	return model.PhotoResponse{
		PhotoID:   photosResult.PhotoID,
		Title:     photosResult.Title,
		Caption:   photosResult.Caption,
		PhotoURL:  photosResult.PhotoURL,
		UserID:    photosResult.UserID,
		Comments:  comments,
		CreatedAt: photosResult.CreatedAt,
		UpdatedAt: photosResult.UpdatedAt,
	}, nil
}

func (s *PhotoServiceIml) UpdatePhoto(photoReqData model.PhotoUpdateRequest, userID string, photoID string) (*model.PhotoResponse, error) {
	findPhotoResponse, err := s.photoRepository.FindByID(photoID)
	if err != nil {
		return nil, err
	}

	if userID != findPhotoResponse.UserID {
		return nil, errors.New("Unauthorized")
	}

	updatedPhotoReq := model.Photo{
		PhotoID:   photoID,
		Title:     photoReqData.Title,
		Caption:   photoReqData.Caption,
		PhotoURL:  photoReqData.PhotoURL,
		UserID:    userID,
		UpdatedAt: time.Now(),
	}

	err = s.photoRepository.Update(updatedPhotoReq)
	if err != nil {
		return nil, err
	}

	return &model.PhotoResponse{
		PhotoID: photoID,
	}, nil
}

func (s *PhotoServiceIml) Delete(photoID string, userID string) (model.PhotoResponse, error) {
	findPhotoResponse, err := s.photoRepository.FindByID(photoID)
	if err != nil {
		return model.PhotoResponse{}, err
	}

	if userID != findPhotoResponse.UserID {
		return model.PhotoResponse{}, errors.New("Unauthorized")
	}

	err = s.photoRepository.Delete(model.Photo{PhotoID: photoID})
	if err != nil {
		return model.PhotoResponse{}, err
	}

	return model.PhotoResponse{
		PhotoID: photoID,
	}, nil
}
