package service

import (
	
	"github.com/gin-gonic/gin"
	"github.com/tools/iservice/dto"
	"github.com/tools/iservice/model"
	"github.com/tools/iservice/repository"
)

type ServicePartyService struct{}

// @Summary Get New ServiceParty Id
// @Description Get New ServiceParty Id
// @Produce json
// @Success 200 {object} dto.ResponseDtoV2
// @Router /iservice/serive-party/get-service-party-id [get]
// @Tags ServiceParty Entry Service
// @Security BearerAuth
func (sps *ServicePartyService) GetNewIdForServiceParty(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	repo := new(repository.ServicePartyRepo)
	output := repo.GetServicePartyCode()
	c.JSON(output.StatusCode, output)
}

// @Summary Get All ServiceParty
// @Description Get All ServiceParty
// @Produce json
// @Success 200 {object} dto.ResponseDtoV2
// @Router /iservice/service-party/get-all-service-party [get]
// @Tags ServiceParty Entry Service
// @Security BearerAuth
func (sps *ServicePartyService) GetAllServiceParty(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	repo := new(repository.ServicePartyRepo)
	output := repo.GetAllServiceParty()
	c.JSON(output.StatusCode, output)
}

// @Summary Get All ServiceParty
// @Description Get All ServiceParty
// @Produce json
// @Success 200 {object} dto.ResponseDtoV2
// @Router /iservice/service-party/get-all-service-party-by-name [get]
// @Tags ServiceParty Entry Service
// @Security BearerAuth
func (sps *ServicePartyService) GetAllServicePartyByName(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	repo := new(repository.ServicePartyRepo)
	output := repo.GetAllServicePartyByName()
	c.JSON(output.StatusCode, output)
}

// @Summary Get One ServiceParty
// @Description Get One ServiceParty
// @Produce json
// @Param buyer body dto.PartyCode true "ServiceParty ID to get ServiceParty details"
// @Success 200 {array} dto.ResponseDtoV2
// @Router /iservice/service-party/get-one-service-party [post]
// @Tags ServiceParty Entry Service
// @Security BearerAuth
func (sps *ServicePartyService) GetOneServiceParty(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	var input dto.PartyCode
	c.ShouldBind(&input)
	repo := new(repository.ServicePartyRepo)
	output := repo.GetOneParty(input)
	c.JSON(output.StatusCode, output)
}

// @Summary Create One ServiceParty
// @Description Create One ServiceParty
// @Produce json
// @Param buyer body model.ServiceParty true "ServiceParty ID to get ServiceParty details"
// @Success 200 {array} dto.ResponseDtoV2
// @Router /iservice/service-party/create-one-service-party [post]
// @Tags ServiceParty Entry Service
// @Security BearerAuth
func (sps *ServicePartyService) CreateOneServiceParty(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	var input model.ServiceParty
	c.ShouldBind(&input)
	repo := new(repository.ServicePartyRepo)
	output := repo.CreateParty(input)
	c.JSON(output.StatusCode, output)
}

// @Summary Update One ServiceParty
// @Description Update One ServiceParty
// @Produce json
// @Param buyer body dto.UpdateParty true "ServiceParty ID to get ServiceParty details"
// @Success 200 {array} dto.ResponseDtoV2
// @Router /iservice/service-party/update-service-party [post]
// @Tags ServiceParty Entry Service
// @Security BearerAuth
func (sps *ServicePartyService) UpdateOneServiceParty(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	var input dto.UpdateParty
	c.ShouldBind(&input)
	repo := new(repository.ServicePartyRepo)
	output := repo.UpdateParty(input)
	c.JSON(output.StatusCode, output)
}

// @Summary Create One ServiceParty
// @Description Create One ServiceParty
// @Produce json
// @Param buyer body dto.DeleteParty true "ServiceParty ID to get ServiceParty details"
// @Success 200 {array} dto.ResponseDto
// @Router /iservice/service-party/delete-service-party [post]
// @Tags ServiceParty Entry Service
// @Security BearerAuth
func (sps *ServicePartyService) DeleteServiceParty(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	var input dto.DeleteParty
	c.ShouldBind(&input)
	repo := new(repository.ServicePartyRepo)
	output := repo.DeleteOneParty(input)
	c.JSON(output.StatusCode, output)
}


func (sps *ServicePartyService) AddRouters(router *gin.Engine) {
	router.GET("/iservice/service-party/get-service-party-id", sps.GetNewIdForServiceParty)
	router.GET("/iservice/service-party/get-all-service-party", sps.GetAllServiceParty)
	router.GET("/iservice/service-party/get-all-service-party-by-name", sps.GetAllServicePartyByName)
	router.POST("/iservice/service-party/get-one-service-party", sps.GetOneServiceParty)
	router.POST("/iservice/service-party/create-one-service-party", sps.CreateOneServiceParty)
	router.POST("/iservice/service-party/update-service-party", sps.UpdateOneServiceParty)
	router.POST("/iservice/service-party/delete-service-party", sps.DeleteServiceParty)

}
