package tools

//
//import (
//	"errors"
//	"fmt"
//	"github.com/dgrijalva/jwt-go"
//	"time"
//)
//
//const (
//	SECRETKEY = "243223ffslsfsldfl412fdsfsdf" //私钥
//)
//
//type CustomClaims struct {
//	UserId int64
//	jwt.StandardClaims
//}
//
//type AuthInfo struct {
//	ID       uint64
//	Username string
//}
//
//func DispatchToken(id int64) string {
//	//生成 token
//	maxAge := 60 * 60 * 24
//	customClaims := &CustomClaims{
//		UserId: id, //用户 id
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
//			Issuer:    "jerry",                                                    // 非必须，也可以填充用户名，
//		},
//	}
//	//采用 HMAC SHA256 加密算法
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
//	tokenString, err := token.SignedString([]byte(SECRETKEY))
//	if err != nil {
//		fmt.Println(err)
//	}
//	return tokenString
//}
//
//func VerifyToken(tokenString string) (*AuthInfo, error) {
//	//解析传入的token
//	//第二个参数是一个回调函数，作用是判断生成token所用的签名算法是否和传入token的签名算法是否一致。
//	//算法匹配就返回密钥，用来解析token.
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(SECRETKEY), nil
//	})
//
//	//err不为空，说明token已过期
//	if err != nil {
//		return nil, err
//	}
//
//	//将获取的token中的Claims强转为MapClaims
//	claims, ok := token.Claims.(jwt.MapClaims)
//	//判断token是否有效
//	if !(ok && token.Valid) {
//		return nil, errors.New("cannot convert claim to mapClaim")
//	}
//	//获取payload中的数据
//	info := &AuthInfo{
//		ID:       uint64(claims["id"].(float64)),
//		Username: claims["username"].(string),
//	}
//	return info, nil
//}
