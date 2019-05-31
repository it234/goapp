package jwt

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

const jwtKey = "adtkls"

// 创建token
func CreateToken(m map[string]string, keys ...string) string {
	key := jwtKey
	if len(keys) > 0 {
		key = keys[0]
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	for index, val := range m {
		claims[index] = val
	}
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

// 解析token
func ParseToken(tokenString string, keys ...string) (map[string]string, bool) {
	key := jwtKey
	if len(keys) > 0 {
		key = keys[0]
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapData:=make(map[string]string)
		for index, val := range claims {
			mapData[index] = fmt.Sprintf("%v", val)
		}
		return mapData, true
	} else {
		return nil, false
	}
}
