package service

import (
	"errors"
	"time"

	"chapter3-mygram/helper"
	"chapter3-mygram/model"
	"chapter3-mygram/repository"
)

type CommentService interface {
	Create(commentReqData model.CommentCreateRequest, userID string, photoID string) (*model.CommentResponse, error)
	GetAll() ([]model.CommentResponse, error)
	GetOne(commentID string) (model.CommentResponse, error)
	UpdateComment(commentReqData model.CommentUpdateRequest, userID string, commentID string) (*model.CommentResponse, error)
	Delete(commentID string, userID string) (model.CommentDeleteResponse, error)
}

type CommentServiceIml struct {
	commentRepository repository.CommentRepository
	photoRepository   repository.PhotoRepository
}

func NewCommentService(commentRepo repository.CommentRepository, photoRepo repository.PhotoRepository) CommentService {
	return &CommentServiceIml{
		commentRepository: commentRepo,
		photoRepository:   photoRepo,
	}
}

func (s *CommentServiceIml) Create(commentReqData model.CommentCreateRequest, userID string, photoID string) (*model.CommentResponse, error) {
	_, err := s.photoRepository.FindByID(photoID)
	if err != nil {
		return nil, err
	}

	commentID := helper.GenerateID()
	newComment := model.Comment{
		CommentID: commentID,
		Message:   commentReqData.Message,
		PhotoID:   photoID,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.commentRepository.Create(newComment)
	if err != nil {
		return nil, err
	}

	return &model.CommentResponse{
		CommentID: newComment.CommentID,
		Message:   newComment.Message,
		PhotoID:   newComment.PhotoID,
		UserID:    newComment.UserID,
		CreatedAt: newComment.CreatedAt,
		UpdatedAt: newComment.UpdatedAt,
	}, nil
}

func (s *CommentServiceIml) GetAll() ([]model.CommentResponse, error) {
	commentsResult, err := s.commentRepository.FindAll()
	if err != nil {
		return []model.CommentResponse{}, err
	}

	commentsResponse := []model.CommentResponse{}
	for _, commentRes := range commentsResult {
		commentsResponse = append(commentsResponse, model.CommentResponse(commentRes))
	}

	return commentsResponse, nil
}

func (s *CommentServiceIml) GetOne(commentID string) (model.CommentResponse, error) {
	commentResult, err := s.commentRepository.FindByID(commentID)
	if err != nil {
		return model.CommentResponse{}, err
	}

	return model.CommentResponse(commentResult), nil
}

func (s *CommentServiceIml) UpdateComment(commentReqData model.CommentUpdateRequest, userID string, commentID string) (*model.CommentResponse, error) {
	findCommentResponse, err := s.commentRepository.FindByID(commentID)
	if err != nil {
		return nil, err
	}

	if userID != findCommentResponse.UserID {
		return nil, errors.New("Unauthorized")
	}

	updatedCommentReq := model.Comment{
		CommentID: findCommentResponse.CommentID,
		Message:   commentReqData.Message,
		PhotoID:   findCommentResponse.PhotoID,
		UserID:    userID,
		UpdatedAt: time.Now(),
	}

	err = s.commentRepository.Update(updatedCommentReq)
	if err != nil {
		return nil, err
	}

	return &model.CommentResponse{
		CommentID: commentID,
	}, nil
}

func (s *CommentServiceIml) Delete(commentlID string, userID string) (model.CommentDeleteResponse, error) {
	findCommentResponse, err := s.commentRepository.FindByID(commentlID)
	if err != nil {
		return model.CommentDeleteResponse{}, err
	}

	if userID != findCommentResponse.UserID {
		return model.CommentDeleteResponse{}, errors.New("Unauthorized")
	}

	err = s.commentRepository.Delete(model.Comment{CommentID: commentlID})
	if err != nil {
		return model.CommentDeleteResponse{}, err
	}

	return model.CommentDeleteResponse{
		CommentID: commentlID,
	}, nil
}
