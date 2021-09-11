package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Render(ctx *gin.Context, tpl string, data map[string]interface{}) {
	ctx.HTML(http.StatusOK, tpl, data)
}
