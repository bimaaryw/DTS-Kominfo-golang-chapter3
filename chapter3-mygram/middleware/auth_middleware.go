package middleware

import (
	"errors"
	"net/http"
	"strings"

	"chapter3-mygram/helper"
	"chapter3-mygram/model"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if bearerIsExist := strings.Contains(auth, "Bearer"); !bearerIsExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: "Unauthorized",
		})
		return
	}

	token := strings.Split(auth, " ")
	if len(token) < 2 {
		err := errors.New("Must provide Authorization header with format `Bearer {token}`")

		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	claims, err := helper.VerifyAccessToken(token[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	c.Set("username", claims.Username)
	c.Set("userID", claims.UserID)

	c.Next()
}
