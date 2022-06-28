package JwtToken

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type CustomJwtStruct struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func CreateToken(c CustomJwtStruct) (string, error) {
	// 签名
	SigningKey := []byte(viper.GetString("jwt.sign-key"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(SigningKey)
}

func ParseToken(tokenString string) (*CustomJwtStruct, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomJwtStruct{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.sign-key")), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token错误")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("token无效")
			} else {
				return nil, fmt.Errorf("invalid token")
			}
		}
	}

	if claims, ok := token.Claims.(*CustomJwtStruct); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
