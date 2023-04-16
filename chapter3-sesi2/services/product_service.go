package services

import (
	model "chapter3-sesi2/models"
	"chapter3-sesi2/repository"

	"github.com/google/uuid"
)

type ProductService struct {
	productRepository repository.ProductRepository
	userRepository    repository.UserRepository
}

func NewProductService(productRepository repository.ProductRepository, userRepository repository.UserRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		userRepository:    userRepository,
	}
}

func (service *ProductService) CreateProduct(request model.ProductCreateRequest, email string) (model.ProductCreateResponse, error) {
	productID := uuid.New()

	userRes, err := service.userRepository.UserCheck(model.UserLoginRequest{Email: email})
	if err != nil {
		return model.ProductCreateResponse{}, err
	}

	product := model.Product{
		ProductID:   productID,
		Title:       request.Title,
		Description: request.Description,
		UserID:      userRes.UserID,
	}

	result, err := service.productRepository.CreateProduct(product)
	if err != nil {
		return model.ProductCreateResponse{}, err
	}

	response := model.ProductCreateResponse{
		ProductID:   result.ProductID,
		Title:       result.Title,
		Description: result.Description,
		UserID:      result.UserID,
	}

	return response, nil
}

func (service *ProductService) GetProductByUserID(email string) ([]model.ProductResponse, error) {
	userRes, err := service.userRepository.UserCheck(model.UserLoginRequest{Email: email})
	if err != nil {
		return []model.ProductResponse{}, err
	}

	response, err := service.productRepository.GetByUserID(userRes.UserID)
	if err != nil {
		return []model.ProductResponse{}, err
	}

	products := []model.ProductResponse{}
	for _, product := range response {
		product := model.ProductResponse{
			ProductID:   product.ProductID,
			Title:       product.Title,
			Description: product.Description,
			UserID:      product.UserID,
			CreatedAt:   product.CreatedAt.String(),
			UpdatedAt:   product.UpdatedAt.String(),
		}
		products = append(products, product)
	}

	return products, nil
}

func (service *ProductService) GetAllProduct() ([]model.ProductResponse, error) {
	response, err := service.productRepository.GetAllProduct()
	if err != nil {
		return []model.ProductResponse{}, err
	}

	products := []model.ProductResponse{}
	for _, product := range response {
		product := model.ProductResponse{
			ProductID:   product.ProductID,
			Title:       product.Title,
			Description: product.Description,
			UserID:      product.UserID,
			CreatedAt:   product.CreatedAt.String(),
			UpdatedAt:   product.UpdatedAt.String(),
		}
		products = append(products, product)
	}

	return products, nil
}

func (service *ProductService) DeleteProduct(productID string) error {
	result, err := service.productRepository.FindProduct(productID)
	if err != nil {
		return err
	}

	err = service.productRepository.DeleteProduct(*result)
	if err != nil {
		return err
	}

	return nil
}

func (service *ProductService) UpdatedProduct(productID string, request model.ProductUpdateRequest) (model.ProductResponse, error) {
	result, err := service.productRepository.FindProduct(productID)
	if err != nil {
		return model.ProductResponse{}, err
	}

	updatedProductReq := &model.Product{
		ProductID:   result.ProductID,
		Title:       request.Title,
		Description: request.Description,
	}

	updateResult, err := service.productRepository.UpdateProduct(*updatedProductReq)
	if err != nil {
		return model.ProductResponse{}, err
	}

	response := model.ProductResponse{
		ProductID:   updateResult.ProductID,
		Title:       updateResult.Title,
		Description: updateResult.Description,
		UserID:      result.UserID,
		CreatedAt:   result.CreatedAt.String(),
		UpdatedAt:   updateResult.UpdatedAt.String(),
	}

	return response, nil
}
