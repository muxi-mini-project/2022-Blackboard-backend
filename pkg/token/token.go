package token

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

const (
	ErrorReasonServerBusy = "服务器繁忙"
	ErrorReasonReLogin    = "请重新登录"
)

type Jwt struct {
	StudentID string `json:"student_id"`
	jwt.StandardClaims
}

//ResolveToke resloves token
func ResolveToken(strToken string) (*Jwt, error) {
	token, err := jwt.ParseWithClaims(strToken, &Jwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("blackboard"), nil

	})
	if err != nil {
		return nil, errors.New(ErrorReasonServerBusy + ",或token解析失败")
	}
	claims, ok := token.Claims.(*Jwt)
	if !ok {
		return nil, errors.New(ErrorReasonReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReasonReLogin)
	}
	return claims, nil
}
