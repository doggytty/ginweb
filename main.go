package main

import (
	"github.com/gin-gonic/gin"
	"github.com/doggytty/ginweb/router"
	"net/http"
)

func main(){
	//gin.SetMode(gin.ReleaseMode)
	// 注册一个默认的路由器
	engine := gin.Default()
	//加载模板
	engine.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	//获取当前文件的相对路径
	engine.Static("/assets", "./assets")
	// 配置路由
	router.InitEngine(engine)
	// 绑定端口是8888
	//engine.Run(":8888")
	//s := &http.Server{
	//	Addr:           ":8080",
	//	Handler:        router,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	http.ListenAndServe(":8888", engine)
}
