package controllers

import (
	"fmt"
	model "chapter3-sesi2/models"
	"chapter3-sesi2/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (controller *UserController) Registration(c *gin.Context) {
	var newUser model.UserRegisterRequest

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	response, err := controller.userService.CreateNewUser(newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.UserRegisterResponse{
		UserID: response.UserID,
		Name:   response.Name,
		Email:  response.Email,
		Role:   response.Role,
	})
}

func (controller *UserController) Login(c *gin.Context) {
	var request model.UserLoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	response, err := controller.userService.Login(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: fmt.Sprintf("Invalid email or password"),
		})
		return
	}

	c.JSON(http.StatusOK, model.UserLoginResponse{
		Token: response,
	})
}
