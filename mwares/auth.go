package mwares

import (
	"gobpframe/errs"
	"gobpframe/resp"
	"strings"

	"github.com/gin-gonic/gin"
)

const loggedIn = false

func CheckAuth(ctx *gin.Context) {
	if strings.HasPrefix(ctx.Request.URL.Path, "/statics/") {
		ctx.Next()
		return
	}

	if !loggedIn {
		resp.Failed(ctx, errs.ErrAuthReject, nil)
		return
	}
	ctx.Next()
}
