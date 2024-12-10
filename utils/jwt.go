package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"picbed/models"
	"time"
)

var privateSigningKey = []byte("mongielee")

type CustomClaims struct {
	Username string `json:"username"`
	UserId   int64  `json:"userId"`
	jwt.RegisteredClaims
}

// GenRegisteredClaims 生成JWT
func GenRegisteredClaims(user *models.User, expired time.Duration) (string, error) {
	claims := &CustomClaims{
		Username: user.Username,
		UserId:   user.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "picbed-logic",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expired)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(privateSigningKey)
}

func GenerateAccessToken(user *models.User) (string, error) {
	return GenRegisteredClaims(user, time.Hour*12)
}

func GenerateRefreshToken(user *models.User) (string, error) {
	return GenRegisteredClaims(user, time.Hour*24)
}

// ValidRegisteredClaims 校验JWT有效性，不解析
func ValidRegisteredClaims(tokenStr string) bool {
	parse, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return privateSigningKey, nil
	})
	if err != nil {
		return false
	}
	return parse.Valid
}

// ParseJWTToken 解析JWT
func ParseJWTToken(tokenStr string) (*CustomClaims, error) {
	parse, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateSigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parse.Claims.(*CustomClaims); ok && parse.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid parse")
}
