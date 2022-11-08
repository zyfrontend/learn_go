package tools

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

var mysigningKey = []byte("292929290100101293231")

func DispatchToken(id int64) string {
	c := MyClaims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,       // 签发时间
			ExpiresAt: time.Now().Unix() + 60*60*24, // 过期时间
			Issuer:    "zy",                         // 签发人
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 加密
	str, err := t.SignedString(mysigningKey)
	if err != nil {
		fmt.Println(err)
	}
	return str
}

func VerifyToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mysigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
