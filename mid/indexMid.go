package mid

import (
	"github.com/gin-gonic/gin"
	"gweb/log"
)

type IndexMid struct{}

func NewIndexMid() *IndexMid {
	return new(IndexMid)
}

func (t *IndexMid) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		log.Println("请求前置操作...")
		c.Next()
		// 请求后
	}
}
