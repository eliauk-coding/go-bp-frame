package resp

import (
	"gobpframe/errs"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code errs.ErrCode `json:"code"` // error code((0: success, otherwise: failed with error code))
	Msg  string       `json:"msg"`  // error message
	Data interface{}  `json:"data"` // response data
}

type Page struct {
	PageIndex int   `json:"pageIndex"` // 分页索引
	PageSize  int   `json:"pageSize"`  // 每页数量
	Total     int64 `json:"total"`     // 总数
}

type ResponseWithPage struct {
	Response
	Page Page `json:"page"`
}

func jsonRespWithMsg(ctx *gin.Context, code errs.ErrCode, msg string, data interface{}) {
	if ctx.IsAborted() {
		return
	}
	res := data
	if data == nil {
		res = map[string]interface{}{}
	}
	ctx.JSON(200, Response{Code: code, Msg: msg, Data: res})
}

func JsonResp(ctx *gin.Context, code errs.ErrCode, data interface{}) {
	jsonRespWithMsg(ctx, code, errs.GetErrMsg(code), data)
}
