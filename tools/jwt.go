package tools

import (
	"errors"
	"fmt"
	//"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v4"
)

func Sign() {
	token := jwt.New(jwt.SigningMethodHS256)
	str, err := token.SignedString("123456")
	fmt.Println(str, err)
}

func VerifyToken(tokenString string) (*AuthInfo, error) {
	//解析传入的token
	//第二个参数是一个回调函数，作用是判断生成token所用的签名算法是否和传入token的签名算法是否一致。
	//算法匹配就返回密钥，用来解析token.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})

	//err不为空，说明token已过期
	if err != nil {
		return nil, err
	}

	//将获取的token中的Claims强转为MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	//判断token是否有效
	if !(ok && token.Valid) {
		return nil, errors.New("cannot convert claim to mapClaim")
	}
	//获取payload中的数据
	info := &AuthInfo{
		ID:       int64(claims["id"].(float64)),
		Username: claims["username"].(string),
	}
	return info, nil
}
