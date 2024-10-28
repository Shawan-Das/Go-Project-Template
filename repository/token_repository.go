package repository

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tools/iservice/dto"
	"github.com/tools/iservice/util"
)

type TokenRepository struct{}

func (tokenRepository *TokenRepository) CreateToken(username string) dto.CreateTokenOutput {
	if username == "" {
		return dto.CreateTokenOutput{
			Message:   "Token creation failed. Username is empty",
			IsSuccess: false,
		}
	}
	mySigningKey := []byte(util.ViperReturnStringConfigVariableFromLocalConfigJSON("ACCESS_SECRET"))
	type MyCustomClaims struct {
		Preferred_username string `json:"preferred_username"`
		jwt.StandardClaims
	}
	claims := MyCustomClaims{
		username,
		jwt.StandardClaims{
			Audience:  username,
			ExpiresAt: time.Now().Add(10 * time.Minute * 60).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Issuer:    "TBS_COMMON_API",
			Subject:   username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return dto.CreateTokenOutput{
			Message:      "Token creation failed. Something went wrong. Please login again.",
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

func (tokenRepository *TokenRepository) ValidateToken(tokenString string) dto.ValidateTokenOutput {
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
		return []byte(util.ViperReturnStringConfigVariableFromLocalConfigJSON("ACCESS_SECRET")), nil
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
