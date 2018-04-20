package router

import (
	"github.com/gin-gonic/gin"
	"github.com/doggytty/ginweb/controllers"
)

func InitEngine(router *gin.Engine)  {
	// 最基本的用法
	router.GET("/test1", func1)
	router.POST("/test2", func2)
	router.GET("/demo/:name/:password", controllers.ParamTest)

	//router.GET("/someGet", getting)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)
}
