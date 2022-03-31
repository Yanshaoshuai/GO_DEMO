package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//实例化一个gin的默认server对象 默认server对象会开启Logger和Recovery插件
	r := gin.Default()
	//New()返回的server默认没有Logger和Recovery插件
	newServer := gin.New()
	newServer.Use(gin.Logger(), gin.Recovery())
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8083") //默认8080
}
