package index

import (
	"github.com/gin-gonic/gin"
	"gobpframe/resp"
)

type controller struct{}

func NewController() *controller {
	return &controller{}
}

func (ctrl *controller) Index(c *gin.Context) {
	name := c.DefaultQuery("name", "friend")
	resp.Render(c, "views/index.html", map[string]interface{}{"name": name})
}
