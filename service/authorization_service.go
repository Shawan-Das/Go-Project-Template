package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tools/system/repository"
)

type AuthorizationService struct{}

func (authorizationService *AuthorizationService) ValidateAuthorization(c *gin.Context) {
	authorizationRepository := new(repository.AuthorizationRepository)
	output := authorizationRepository.ValidateAuthorization(c)
	if !output.IsSuccess {
		c.AbortWithStatusJSON(http.StatusBadRequest, output)
	} else {
		c.JSON(http.StatusOK, output)
	}
}

func (authorizationService *AuthorizationService) ValidateAuthorization_V2(c *gin.Context) {
	authorizationRepository := new(repository.AuthorizationRepository)
	output := authorizationRepository.ValidateAuthorization_V2(c)
	if !output.IsSuccess {
		c.AbortWithStatusJSON(http.StatusBadRequest, output)
	} else {
		c.JSON(http.StatusOK, output)
	}
}

// AddRouters add api end points specific to this service
func (authorizationService *AuthorizationService) AddRouters(router *gin.Engine) {
	router.POST("/ValidateToken", authorizationService.ValidateAuthorization)
	router.POST("/ValidateToken_V2", authorizationService.ValidateAuthorization_V2)
}
