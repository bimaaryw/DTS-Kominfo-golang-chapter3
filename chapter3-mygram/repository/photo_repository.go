package repository

import (
	"errors"

	"chapter3-mygram/model"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photoReqData model.Photo) error
	FindAll() ([]model.Photo, error)
	FindByID(photoID string) (model.Photo, error)
	FindByUserID(userID string) ([]model.Photo, error)
	Update(photoReqData model.Photo) error
	Delete(photoReqData model.Photo) error
}

type PhotoRepositoryImpl struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &PhotoRepositoryImpl{
		DB: db,
	}
}

func (r *PhotoRepositoryImpl) Create(photoReqData model.Photo) error {
	err := r.DB.Create(&photoReqData).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *PhotoRepositoryImpl) FindAll() ([]model.Photo, error) {
	photos := []model.Photo{}

	err := r.DB.Find(&photos).Error
	if err != nil {
		return []model.Photo{}, err
	}

	return photos, nil
}

func (r *PhotoRepositoryImpl) FindByID(photoID string) (model.Photo, error) {
	photo := model.Photo{}

	err := r.DB.Debug().Where("photo_id = ?", photoID).Take(&photo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Photo{}, err
		}

		return model.Photo{}, err
	}

	return photo, nil
}

func (r *PhotoRepositoryImpl) FindByUserID(userID string) ([]model.Photo, error) {
	photos := []model.Photo{}

	err := r.DB.Where("user_id = ?", userID).Find(&photos).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Photo{}, err
		}

		return []model.Photo{}, err
	}

	return photos, nil
}

func (r *PhotoRepositoryImpl) Update(photoReqData model.Photo) error {
	err := r.DB.Save(&model.Photo{
		PhotoID:   photoReqData.PhotoID,
		Title:     photoReqData.Title,
		Caption:   photoReqData.Caption,
		PhotoURL:  photoReqData.PhotoURL,
		UserID:    photoReqData.UserID,
		UpdatedAt: photoReqData.UpdatedAt,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *PhotoRepositoryImpl) Delete(photoReqData model.Photo) error {
	err := r.DB.Delete(&photoReqData).Error
	if err != nil {
		return err
	}

	return nil
}
