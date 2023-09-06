package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Server 实例化服务器
func Server(port string, engine *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    port,
		Handler: engine,
	}
}

// Redirect 请求内跳转
func Redirect(c *gin.Context, code int, url string) {
	c.Redirect(code, url)
}

// Forward 路由跳
func Forward(c *gin.Context, r *gin.Engine, url string) {
	c.Request.URL.Path = url
	r.HandleContext(c)
}
