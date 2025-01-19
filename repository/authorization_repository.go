package repository

import (
	"strings"

	"github.com/gin-gonic/gin"
	// "github.com/tools/common/model"
	"github.com/tools/system/dto"
	// "github.com/tools/system/model/customreport"
	"github.com/tools/system/util"
)

type AuthorizationRepository struct {
	tokenRepository *TokenRepository
}

func (authorizationrepository *AuthorizationRepository) ValidateAuthorization_V2(c *gin.Context) dto.ValidateAuthorizationOutput {
	url := c.Request.URL
	uri := url.RequestURI()

	config, err := util.LoadConfig(util.ConfigDetails)
	if err != nil {
		return dto.ValidateAuthorizationOutput{
			Message:   "Unable to read local-config.json file.",
			IsSuccess: false,
		}
	}

	authorizationUrls, ok := config["bypassAuth_V2"].(map[string]interface{})
	if !ok {
		return dto.ValidateAuthorizationOutput{
			Message:   "Invalid authorizationUrls format in config.",
			IsSuccess: false,
		}
	}

	equalFoldUrls, equalFoldOk := authorizationUrls["reqBody"].([]interface{})
	containsUrls, containsOk := authorizationUrls["params"].([]interface{})

	if !equalFoldOk || !containsOk {
		return dto.ValidateAuthorizationOutput{
			Message:   "Invalid authorizationUrls format in config.",
			IsSuccess: false,
		}
	}

	equalFoldStrUrls := util.ToStringSlice(equalFoldUrls)
	containsStrUrls := util.ToStringSlice(containsUrls)

	if util.ContainsString(equalFoldStrUrls, uri) || util.ContainsSubstring(containsStrUrls, uri) {
		return dto.ValidateAuthorizationOutput{
			Message:   "Authorization successful.",
			IsSuccess: true,
		}
	}

	authHeader := c.Request.Header.Get("Authorization")
	if len(authHeader) == 0 || !strings.HasPrefix(authHeader, "Bearer") {
		return dto.ValidateAuthorizationOutput{
			Message:   "Invalid authotization header. Please login again.",
			IsSuccess: false,
		}
	}

	authArray := strings.Split(authHeader, " ")
	if len(authArray) != 2 {
		return dto.ValidateAuthorizationOutput{
			Message:   "Unable to parse authorization token. Please login again.",
			IsSuccess: false,
		}
	}

	validateTokenOutput := authorizationrepository.tokenRepository.ValidateToken(authArray[1])
	if !validateTokenOutput.IsSuccess {
		return dto.ValidateAuthorizationOutput{
			Message:   "Authorization failed: Token is invalid. Please login again.",
			IsSuccess: false,
			Payload:   validateTokenOutput,
		}
	}

	return dto.ValidateAuthorizationOutput{
		Message:   "Authorization successful. Token is valid.",
		IsSuccess: true,
	}
}

// TODO: OLD CODE
func (authorizationRepository *AuthorizationRepository) ValidateAuthorization(c *gin.Context) dto.ValidateAuthorizationOutput {
	url := c.Request.URL
	uri := url.RequestURI()
	if strings.EqualFold(uri, "/login") || strings.EqualFold(uri, "/CreateToken") {
		return dto.ValidateAuthorizationOutput{
			Message:   "Authorization successful.",
			IsSuccess: true,
		}
	}
	authHeader := c.Request.Header.Get("Authorization")
	if len(authHeader) == 0 || !strings.HasPrefix(authHeader, "Bearer") {
		return dto.ValidateAuthorizationOutput{
			Message:   "Invalid authotization header. Please login again.",
			IsSuccess: false,
		}
	}
	authArray := strings.Split(authHeader, " ")
	if len(authArray) != 2 {
		return dto.ValidateAuthorizationOutput{
			Message:   "Unable to parse authorization token. Please login again.",
			IsSuccess: false,
		}
	}
	validateTokenOutput := authorizationRepository.tokenRepository.ValidateToken(authArray[1])
	if !validateTokenOutput.IsSuccess {
		return dto.ValidateAuthorizationOutput{
			Message:   "Authorization failed: Token is invalid. Please login again.",
			IsSuccess: false,
			Payload:   validateTokenOutput,
		}
	}
	// token to db validation required
	// db := util.CreateConnectionUsingGormToCommonSchema()
	// sqlDB, _ := db.DB()
	// defer sqlDB.Close()
	// var login model.Login
	// result := db.Raw("select * from common.login where token = ?", authArray[1]).First(&login)
	// if result.RowsAffected == 0 {
	// 	return customreport.ValidateAuthorizationOutput{
	// 		Message:   "Authorization failed: Token doesn't exists.",
	// 		IsSuccess: false,
	// 	}
	// }
	return dto.ValidateAuthorizationOutput{
		Message:   "Authorization successful. Token is valid.",
		IsSuccess: true,
	}
}
