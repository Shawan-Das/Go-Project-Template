package repository

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tools/system/dto"
	"github.com/tools/system/model"
	"github.com/tools/system/util"
)

type Token_Repository struct{}

func (tokenrepository *Token_Repository) CreateTokenRepo(username string) dto.CreateTokenOutput {
	if username == "" {
		return dto.CreateTokenOutput{
			Message:   "Token creation failed. Username is empty",
			IsSuccess: false,
		}
	}

	mySigningKey := []byte(util.ViperReturnStringConfigVariableFromLocalConfigJSON("access_secret"))
	type MyCustomClaims struct {
		Preferred_username string `json:"preferred_username"`
		jwt.StandardClaims
	}

	claims := MyCustomClaims{
		username,
		jwt.StandardClaims{
			Audience:  username,
			ExpiresAt: time.Now().Add(time.Duration(util.ViperReturnIntegerConfigVariableFromLocalConfigJSON("timeout")) * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Issuer:    "Issuer",
			Subject:   username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return dto.CreateTokenOutput{
			Message:      "Token creation failed",
			IsSuccess:    false,
			Token:        ss,
			ErrorMessage: err.Error(),
		}
	}
	return dto.CreateTokenOutput{
		Message:   "Token creation successful",
		IsSuccess: true,
		Token:     ss,
	}
}

func (tokenRepository *Token_Repository) ValidateToken_V2(tokenString string) dto.ValidateTokenOutput {
	if tokenString == "" {
		return dto.ValidateTokenOutput{
			Message:   "Token is empty. Please login again",
			IsSuccess: false,
		}
	}
	type MyCustomClaims struct {
		Preferred_username string `json:"preferred_username"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(util.ViperReturnStringConfigVariableFromLocalConfigJSON("JWT_TOKEN_SECRET")), nil
	})
	if err != nil {
		return dto.ValidateTokenOutput{
			Message:      "Token validation failed. Please login again.",
			IsSuccess:    false,
			ErrorMessage: err.Error(),
		}
	}
	if _, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return dto.ValidateTokenOutput{
			Message:   "Token validated successfully",
			IsSuccess: true,
		}
	}
	return dto.ValidateTokenOutput{
		Message:   "Token validation failed. Something went wrong. Please login again",
		IsSuccess: false,
	}
}

func (tokenRepository *Token_Repository) TokenTimeExtend(input model.Logins) dto.ResponseDto {
	var res dto.ResponseDto
	if input.Username == "" {
		res.Message = "Token extension failed. Username is empty"
		res.IsSuccess = false
		return res
	}

	CreateToken := tokenRepository.CreateTokenRepo(input.Username)
	if !CreateToken.IsSuccess {
		res.Message = "Error while creating token"
		res.IsSuccess = false
		res.Payload = nil
		res.StatusCode = http.StatusInternalServerError
		return res
	}
	// database connection
	db := util.CreateConnectionUsingGormToCommonSchema()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	// if no action then close the connection

	input.Token_value = CreateToken.Token
	result4 := db.Where(&model.Logins{Username: input.Username}).Updates(&input)
	if result4.RowsAffected == 0 {
		res.Message = result4.Error.Error()
		res.IsSuccess = false
		res.Payload = nil
		res.StatusCode = http.StatusInternalServerError
		return res
	}

	type tempOutPut struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	}

	var temp tempOutPut
	temp.Username = input.Username
	temp.Token = input.Token_value

	res.Message = "Token extended successfully"
	res.IsSuccess = true
	res.Payload = temp
	res.StatusCode = http.StatusOK

	return res
}

func (tokenRepository *TokenRepository) TokenTimeExtend_V2(input model.Logins) dto.ResponseDto {
	var res dto.ResponseDto
	if input.Username == "" {
		res.Message = "Token extension failed. Username is empty"
		res.IsSuccess = false
		return res
	}

	// database connection
	db := util.CreateConnectionUsingGormToCommonSchema()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	// if no action then close the connection

	check1 := db.Where("username = ?", input.Username).First(&input)
	if check1.RowsAffected == 0 {
		res.Message = "For security reasons, kindly <b>log in</b> again as the current token lacks authorization. Thank you."
		res.IsSuccess = false
		res.Payload = nil
		res.StatusCode = http.StatusNotFound
		return res
	}

	if input.Token_value == "" {
		res.Message = "Token extension failed. Token is empty"
		res.IsSuccess = false
		return res
	}

	tokenStr := input.Token_value

	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(util.ViperReturnStringConfigVariableFromLocalConfigJSON("access_secret")), nil // Replace with your own secret key.
	})
	preferred_username := token.Claims.(jwt.MapClaims)["preferred_username"].(string)

	if preferred_username != input.Username {
		res.Message = "For security reasons, kindly <b>log in</b> again as the current token lacks authorization. Thank you."
		res.IsSuccess = false
		res.Payload = nil
		res.StatusCode = http.StatusNotFound
		return res
	}

	// Get the existing expiration time from the token.
	expTime := token.Claims.(jwt.MapClaims)["exp"].(float64)

	// TODO: new code
	// Calculate the new expiration time by adding local-configue json minutes.
	extensionMinutes := util.ViperReturnIntegerConfigVariableFromLocalConfigJSON("timeout")
	extensionDuration := time.Duration(extensionMinutes) * time.Minute
	newExpTime := time.Unix(int64(expTime), 0).Add(extensionDuration).Unix()
	// // TODO: old code
	// // Calculate the new expiration time by adding 30 minutes.
	// newExpTime := time.Unix(int64(expTime), 0).Add(1 * time.Minute).Unix()

	// Create a new JWT token with the updated expiration time.
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"preferred_username": input.Username,
		"exp":                newExpTime,
		"nbf":                time.Now().Unix(),
	})

	// Sign the new token with the same secret key as the original token.
	newTokenStr, err := newToken.SignedString([]byte(util.ViperReturnStringConfigVariableFromLocalConfigJSON("access_secret")))
	if err != nil {
		res.Message = "Token extension failed. Please login again."
		res.IsSuccess = false
		res.Payload = nil
		return res
	}

	input.Token_value = newTokenStr
	result4 := db.Where(&model.Logins{Username: input.Username}).Updates(&input)
	if result4.RowsAffected == 0 {
		res.Message = result4.Error.Error()
		res.IsSuccess = false
		res.Payload = nil
		res.StatusCode = http.StatusInternalServerError
		return res
	}
	type tempOutPut struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	}

	var temp tempOutPut
	temp.Username = input.Username
	temp.Token = input.Token_value
	res.Message = "Token extended successfully"
	res.IsSuccess = true
	res.Payload = temp
	res.StatusCode = http.StatusOK

	return res
}

// func (tokenRepository *TokenRepository) GeninfoRepo(input model.Comp_year_token_input_dto) dto.ResponseDto {
// 	var res dto.ResponseDto

// 	if input.Geninfo == "" {
// 		res.Message = "Token creation failed. Geninfo is empty"
// 		res.IsSuccess = false
// 		return res
// 	}

// 	db := util.CreateConnectionUsingGormToCommonSchema()
// 	sqlDB, _ := db.DB()
// 	defer sqlDB.Close()

// 	// _ = db.Raw("SELECT MAX(id) FROM common.comp_year_token WHERE username = ?;",input.Username).First(&input.Id)
// 	var check model.Logins
// 	_ = db.Where("id = ?", input.User_id).First(&check)

// 	var output2 model.Comp_year_token

// 	result := db.Where("username = ?", check.Username).First(&output2)
// 	if result.RowsAffected == 0 {
// 		res.Message = result.Error.Error()
// 		res.IsSuccess = false
// 		res.Payload = nil
// 		res.StatusCode = http.StatusInternalServerError
// 		return res
// 	}

// 	tokenStr := output2.Geninfo

// 	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(util.ViperReturnStringConfigVariableFromLocalConfigJSON("access_secret")), nil // Replace with your own secret key.
// 	})
// 	username := token.Claims.(jwt.MapClaims)["username"].(string)
// 	compcode := token.Claims.(jwt.MapClaims)["compcode"].(float64)
// 	yearcode := token.Claims.(jwt.MapClaims)["yearcode"].(float64)
// 	var output dto.ModuleLoad_Swag_dto
// 	output.Username = username
// 	output.Compcode = int(compcode)
// 	output.Yearcode = int(yearcode)

// 	res.Message = "Geninfo Successfully Get"
// 	res.IsSuccess = true
// 	res.Payload = output
// 	res.StatusCode = http.StatusOK

// 	return res
// }
