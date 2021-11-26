package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-web/enums/jwtEnum"
	Models "gin-web/models"

	"github.com/dgrijalva/jwt-go"

	"time"
)

// create token
func UserJsonWebToken(info string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		`user`: info,
		`exp`:  time.Now().Add(time.Hour * 10).Unix(),
	})
	return token.SignedString([]byte(jwtEnum.UserSecret))
}
func LoginUserGetJWTToken(token string) (*Models.User, error) {
	c, err := LoginUserDecryptJsonWebToken(token)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(`token无效,请重新获取授权！`)
	}
	lo := Models.User{}
	err = json.Unmarshal([]byte(c[`user`].(string)), &lo)
	if err != nil {
		return nil, errors.New(`token无效,请重新获取授权！`)
	}
	return &lo, nil
}

func LoginUserDecryptJsonWebToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtEnum.UserSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		fmt.Println(err)
		return nil, err
	}

}
