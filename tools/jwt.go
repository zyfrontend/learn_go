package tools

import (
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
	// 生成 token MapClaims方式
	//t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"exp":      time.Now().Unix() + 5, // 过期时间
	//	"iss":      id,                    // 签发人id
	//	"nbf":      time.Now().Unix() - 5, // 签发时间
	//	"username": "zy",                  // 签发人
	//})
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 加密
	str, err := t.SignedString(mysigningKey)
	if err != nil {
		fmt.Println(err)
	}
	return str
}

func VerifyToken(tokenString string) {
	//token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mysigningKey, nil
	})
	if err != nil {
		fmt.Println("err", err)
		return
	}

	//fmt.Println(token.Claims.(*jwt.MapClaims))
	fmt.Println(token.Claims.(*MyClaims))
}
