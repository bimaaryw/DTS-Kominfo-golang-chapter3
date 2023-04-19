package repository

import (
	"errors"

	"chapter3-mygram/model"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(photoReqData model.SocialMedia) error
	FindAll() ([]model.SocialMedia, error)
	FindByID(socialID string) (model.SocialMedia, error)
	FindByUserID(userID string) ([]model.SocialMedia, error)
	Update(socialReqData model.SocialMedia) error
	Delete(photoReqData model.SocialMedia) error
}

type SocialMediaRepositoryImpl struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &SocialMediaRepositoryImpl{
		DB: db,
	}
}

func (r *SocialMediaRepositoryImpl) Create(socialReqData model.SocialMedia) error {
	err := r.DB.Create(&socialReqData).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *SocialMediaRepositoryImpl) FindAll() ([]model.SocialMedia, error) {
	socials := []model.SocialMedia{}

	err := r.DB.Find(&socials).Error
	if err != nil {
		return []model.SocialMedia{}, err
	}

	return socials, nil
}

func (r *SocialMediaRepositoryImpl) FindByID(socialID string) (model.SocialMedia, error) {
	social := model.SocialMedia{}

	err := r.DB.Debug().Where("id = ?", socialID).Take(&social).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SocialMedia{}, err
		}

		return model.SocialMedia{}, err
	}

	return social, nil
}

func (r *SocialMediaRepositoryImpl) FindByUserID(userID string) ([]model.SocialMedia, error) {
	socials := []model.SocialMedia{}

	err := r.DB.Debug().Where("user_id = ?", userID).Find(&socials).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.SocialMedia{}, err
		}

		return []model.SocialMedia{}, err
	}

	return socials, nil
}

func (r *SocialMediaRepositoryImpl) Update(socialReqData model.SocialMedia) error {
	err := r.DB.Save(&model.SocialMedia{
		ID:             socialReqData.ID,
		Name:           socialReqData.Name,
		SocialMediaURL: socialReqData.SocialMediaURL,
		UserID:         socialReqData.UserID,
		UpdatedAt:      socialReqData.UpdatedAt,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *SocialMediaRepositoryImpl) Delete(socialReqData model.SocialMedia) error {
	err := r.DB.Delete(&socialReqData).Error
	if err != nil {
		return err
	}

	return nil
}
