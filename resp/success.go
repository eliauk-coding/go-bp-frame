package resp

import (
	"gobpframe/errs"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data interface{}) {
	jsonRespWithMsg(ctx, errs.Success, errs.GetErrMsg(errs.Success), data)
}

func SuccessWithMsg(ctx *gin.Context, msg string, data interface{}) {
	jsonRespWithMsg(ctx, errs.Success, msg, data)
}
