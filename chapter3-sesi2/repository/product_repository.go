package repository

import (
	"errors"
	model "chapter3-sesi2/models"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (repository *ProductRepository) CreateProduct(product model.Product) (*model.Product, error) {
	newProduct := model.Product{
		ProductID:   product.ProductID,
		Title:       product.Title,
		Description: product.Description,
		UserID:      product.UserID,
	}

	err := repository.DB.Create(&newProduct).Error
	if err != nil {
		log.Fatal("error")
		return nil, err
	}

	return &newProduct, nil
}

func (repository *ProductRepository) FindProduct(productID string) (*model.Product, error) {
	productResult := model.Product{}

	err := repository.DB.Debug().Where("product_id = ?", productID).Take(&productResult).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}

	return &productResult, nil
}

func (repository *ProductRepository) GetByUserID(userID uuid.UUID) ([]model.Product, error) {
	products := make([]model.Product, 0)
	tx := repository.DB.Where("user_id = ?", userID).Find(&products)
	return products, tx.Error
}

func (repository *ProductRepository) GetAllProduct() ([]model.Product, error) {
	products := []model.Product{}

	err := repository.DB.Find(&products).Error
	if err != nil {
		return []model.Product{}, err
	}

	return products, nil
}

func (repository *ProductRepository) DeleteProduct(product model.Product) error {
	err := repository.DB.Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProductRepository) UpdateProduct(product model.Product) (model.Product, error) {
	productUpdated := product

	err := repository.DB.Model(&productUpdated).Updates(model.Product{
		Title:       product.Title,
		Description: product.Description,
	}).Error
	if err != nil {
		return model.Product{}, err
	}

	return productUpdated, nil
}
