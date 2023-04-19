package controller

import (
	"net/http"

	"chapter3-mygram/helper"
	"chapter3-mygram/model"
	"chapter3-mygram/service"

	"github.com/gin-gonic/gin"
)

type SocialController interface {
	CreateSocial(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type SocialControllerImpl struct {
	socialService service.SocialService
}

func NewSocialController(service service.SocialService) SocialController {
	return &SocialControllerImpl{
		socialService: service,
	}
}

// CreateSocialMedia godoc
//
//	@Summary		create social media
//	@Description	create social media for a particular user
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.SocialCreateReq	true	"request is required"
//	@Success		200		{object}	model.SuccessResponse{data=model.SocialResponse}
//	@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/social-media [post]
func (c *SocialControllerImpl) CreateSocial(ctx *gin.Context) {
	var request model.SocialMediaCreateRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.SocialCreateValidator(request)
	if validateErrs != nil {
		errResponseStr := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponseStr[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponseStr,
		})
		return
	}

	response, err := c.socialService.Create(request, userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Social media created successfully",
		Data:    response,
	})
}

// GetAllSocialMedia godoc
//
//	@Summary		get all social media
//	@Description	get all social media
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	model.SuccessResponse{data=[]model.SocialResponse}
//	@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/social-media [get]
func (c *SocialControllerImpl) GetAll(ctx *gin.Context) {
	response, err := c.socialService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Get all social media successfully",
		Data:    response,
	})
}

// GetOneSocialMedia godoc
//
//		@Summary		get one social media
//		@Description	get one social media
//		@Tags			Social Media
//		@Accept			json
//		@Produce		json
//	 @Param          social_media_id   path      string  true  "social_media_id"
//		@Success		200		{object}	model.SuccessResponse{data=model.SocialResponse}
//		@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//		@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//		@Security		BearerAuth
//		@Router			/social-media/{social_media_id} [get]
func (c *SocialControllerImpl) GetOne(ctx *gin.Context) {
	socialID := ctx.Param("social_media_id")

	response, err := c.socialService.GetOne(socialID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Get social media successfully",
		Data:    response,
	})
}

// UpdateSocialMedia godoc
//
//		@Summary		update social media
//		@Description	update social media
//		@Tags			Social Media
//		@Accept			json
//		@Produce		json
//	 @Param          social_media_id   path      string  true  "social_media_id"
//		@Param			request	body		model.SocialUpdateReq	true	"request is required"
//		@Success		200		{object}	model.SuccessResponse{data=model.SocialUpdateRes}
//		@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//		@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//		@Security		BearerAuth
//		@Router			/social-media/{social_media_id} [put]
func (c *SocialControllerImpl) UpdateSocialMedia(ctx *gin.Context) {
	var request model.SocialMediaUpdateRequest
	socialID := ctx.Param("social_media_id")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.SocialUpdateValidator(request)
	if validateErrs != nil {
		errResponseStr := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponseStr[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponseStr,
		})
		return
	}

	response, err := c.socialService.UpdateSocialMedia(request, userID.(string), socialID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Social media updated successfully",
		Data: model.SocialMediaUpdateResponse{
			ID: response.ID,
		},
	})
}

// DeleteSocialMedia godoc
//
//		@Summary		delete social media
//		@Description	delete social media
//		@Tags			Social Media
//		@Accept			json
//		@Produce		json
//	 @Param          social_media_id   path      string  true  "social_media_id"
//		@Success		200		{object}	model.SuccessResponse{data=model.SocialDeleteRes}
//		@Failure		400		{object}	model.ErrorResponse{errors=interface{}}
//		@Failure		500		{object}	model.ErrorResponse{errors=interface{}}
//		@Security		BearerAuth
//		@Router			/social-media/{social_media_id} [delete]
func (c *SocialControllerImpl) DeleteSocialMedia(ctx *gin.Context) {
	socialID := ctx.Param("social_media_id")

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	response, err := c.socialService.Delete(socialID, userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "social media deleted successfully",
		Data: model.SocialMediaDeleteResponse{
			ID: response.ID,
		},
	})
}
