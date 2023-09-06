package routes

import (
	"github.com/gin-gonic/gin"
	"gweb/controller"
	"gweb/mid"
	"log"
	"net/http"
)

var (
	router *gin.Engine
	err    error
)

func init() {
	router = gin.Default()
	if err != nil {
		log.Fatalln(err)
	}
}

// 静态资源路径 ./public
// html模板路径 ./resources
func cfg() {
	gin.ForceConsoleColor()
	router.MaxMultipartMemory = 100 << 20
	router.Static("/public", "./public")
	router.StaticFile("favicon.ico", "./resources/favicon.ico")
	//router.LoadHTMLGlob("resources/templates/*")
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusNotFound,
			"title": "404",
			"msg":   "你访问的资源不存在",
		})
	})
}
func Router() *gin.Engine {
	return router
}

func DefinedRouter() {
	cfg()
	// 首页路由组
	indexGroup := router.Group("/", mid.NewIndexMid().Index())
	{
		index := controller.NewIndexController()
		indexGroup.GET("/", index.Index)
	}
}
