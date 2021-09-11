package resp

import (
	"gobpframe/errs"

	"github.com/gin-gonic/gin"
)

func Failed(ctx *gin.Context, code errs.ErrCode, data interface{}) {
	jsonRespWithMsg(ctx, code, errs.GetErrMsg(code), data)
	ctx.Abort()
}

var FailedWithMsg = jsonRespWithMsg
