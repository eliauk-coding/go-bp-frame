package mwares

import (
	"fmt"
	"gobpframe/errs"
	"gobpframe/resp"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Catch - catch exception
func Catch(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("ERROR: %v\nSTACK: %v", r, string(debug.Stack()))
			resp.Failed(ctx, errs.ErrSvrInternal, msg)
			ctx.Abort()
		}
	}()
	ctx.Next()
}
