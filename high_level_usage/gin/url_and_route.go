package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type Person struct {
	//约束 i的格式必须为 uuid
	//ID   string `uri:"id" binding:"required,uuid"`
	ID   string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	server := gin.Default()
	//路由分组
	goodsGroup := server.Group("/goods")

	{
		goodsGroup.GET("", goodsList)
		goodsGroup.GET("/:id", goodsDetail)
		goodsGroup.POST("", createGoods)
		//静态资源
		goodsGroup.GET("/resource/*action")
	}
	personGroup := server.Group("/person")

	{
		personGroup.GET("/:name/:id", personDetail)
		personGroup.GET("/query", queryPerson)
		personGroup.POST("/post_param", postParam)
	}
	server.Run()
}

//获取post参数
func postParam(context *gin.Context) {
	//获取form参数
	firstname := context.PostForm("firstname")
	lastname := context.DefaultPostForm("lastname", "yan")
	//获取body参数
	//body, _ := ioutil.ReadAll(context.Request.Body)
	var person Person
	context.ShouldBindBodyWith(&person, binding.JSON)
	//获取query参数
	age := context.Query("age")
	context.JSON(http.StatusOK, gin.H{
		"firstname": firstname,
		"lastname":  lastname,
		"age":       age,
		"body":      person,
	})
}

func queryPerson(context *gin.Context) {
	firstname := context.DefaultQuery("firstname", "body")
	lastname := context.Query("lastname")
	context.JSON(http.StatusOK, gin.H{
		"first_name": firstname,
		"last_name":  lastname,
	})
}
func personDetail(context *gin.Context) {
	var person Person
	//绑定参数
	if err := context.ShouldBindUri(&person); err != nil {
		context.Status(http.StatusNotFound)
		return
	}
	context.JSON(http.StatusOK, person)
}
func createGoods(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func goodsDetail(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func goodsList(context *gin.Context) {
	context.JSON(http.StatusOK, []gin.H{
		{"name": "goodsList"},
	})
}
