package repository

import (
	"errors"
	"fmt"

	"chapter3-mygram/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(commentReqData model.Comment) error
	FindAll() ([]model.Comment, error)
	FindByID(commentID string) (model.Comment, error)
	FindByPhotoID(photoID string) ([]model.Comment, error)
	Update(commentReqData model.Comment) error
	Delete(commentReqData model.Comment) error
}

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{
		DB: db,
	}
}

func (r *CommentRepositoryImpl) Create(commentReqData model.Comment) error {
	err := r.DB.Create(&commentReqData).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepositoryImpl) FindAll() ([]model.Comment, error) {
	comments := []model.Comment{}

	err := r.DB.Find(&comments).Error
	if err != nil {
		return []model.Comment{}, err
	}

	return comments, nil
}

func (r *CommentRepositoryImpl) FindByID(commentID string) (model.Comment, error) {
	comment := model.Comment{}

	err := r.DB.Debug().Where("comment_id = ?", commentID).Take(&comment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Comment{}, err
		}

		return model.Comment{}, err
	}

	return comment, nil
}

func (r *CommentRepositoryImpl) FindByPhotoID(photoID string) ([]model.Comment, error) {
	comments := []model.Comment{}

	err := r.DB.Where("photo_id = ?", photoID).Find(&comments).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Comment{}, err
		}

		return []model.Comment{}, err
	}

	fmt.Println("comments: ", comments)

	return comments, nil
}

func (r *CommentRepositoryImpl) Update(commentReqData model.Comment) error {
	err := r.DB.Save(&model.Comment{
		CommentID: commentReqData.CommentID,
		Message:   commentReqData.Message,
		PhotoID:   commentReqData.PhotoID,
		UserID:    commentReqData.UserID,
		UpdatedAt: commentReqData.UpdatedAt,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepositoryImpl) Delete(commentReqData model.Comment) error {
	err := r.DB.Delete(&commentReqData).Error
	if err != nil {
		return err
	}

	return nil
}
