package mwares

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"gobpframe/errs"
	"gobpframe/resp"
)

// JWT - jwt middleware
func JWT(ctx *gin.Context) {
	code := errs.ErrAuthReject
	tokens := strings.Split(ctx.Request.Header.Get("Authorization"), " ")

	if len(tokens) != 2 {
		resp.Failed(ctx, code, nil)
		return
	}

	code = errs.ErrAuthTokenInvalid
	token, err := jwt.Parse(tokens[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("smartbow"), nil
	})

	if err != nil {
		if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
			code = errs.ErrAuthTokenExpired
		}
		resp.Failed(ctx, code, nil)
		return
	}

	user, ok := token.Claims.(jwt.MapClaims)
	_ = user
	if !ok || !token.Valid {
		resp.Failed(ctx, code, nil)
		return
	}

	ctx.Next()
}
