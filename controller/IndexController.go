package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func NewIndexController() *IndexController {
	return new(IndexController)
}
func (t *IndexController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world 这是一个基于Golang的web程序",
	})
}
