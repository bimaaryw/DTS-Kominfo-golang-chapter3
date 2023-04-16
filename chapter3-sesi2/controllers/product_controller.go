package controllers

import (
	model "chapter3-sesi2/models"
	"chapter3-sesi2/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{
		productService: service,
	}
}

func (controller *ProductController) CreateProduct(c *gin.Context) {
	var newProduct model.ProductCreateRequest

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	email, emailIsExist := c.Get("email")
	if !emailIsExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: "Unauthorized",
		})
		return
	}

	response, err := controller.productService.CreateProduct(newProduct, email.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.ProductCreateResponse{
		ProductID:   response.ProductID,
		Title:       response.Title,
		Description: response.Description,
		UserID:      response.UserID,
	})
}

func (controller *ProductController) GetProduct(c *gin.Context) {
	role, roleIsExist := c.Get("role")
	if !roleIsExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: "Unauthorized",
		})
		return
	}

	if role.(bool) == false {
		fmt.Println("Log in disini")
		email, emailIsExist := c.Get("email")
		if !emailIsExist {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Err: "Unauthorized",
			})
			return
		}

		response, err := controller.productService.GetProductByUserID(email.(string))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Err: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	} else {
		response, err := controller.productService.GetAllProduct()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Err: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func (controller *ProductController) DeleteProduct(c *gin.Context) {
	productID := c.Param("product_id")

	err := controller.productService.DeleteProduct(productID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Deleted",
	})
}

func (controller *ProductController) UpdateProduct(c *gin.Context) {
	productID := c.Param("product_id")
	var updatedProductReq model.ProductUpdateRequest

	if err := c.ShouldBindJSON(&updatedProductReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	updatedProductRes, err := controller.productService.UpdatedProduct(productID, updatedProductReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedProductRes)
}
