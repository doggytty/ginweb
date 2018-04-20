package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
func ParamTest (c *gin.Context) {
	name := c.Param("name")
	password := c.Param("password")
	//c.String(http.StatusOK, "参数:%s %s  test3 OK", name, password)
	result := map[string]string {name: password}
	c.JSON(200, result)
}


// func1: 处理最基本的GET
func func1 (c *gin.Context)  {
	// 回复一个200OK,在client的http-get的resp的body中获取数据
	c.String(http.StatusOK, "test1 OK")
}
// func2: 处理最基本的POST
func func2 (c *gin.Context) {
	// 回复一个200 OK, 在client的http-post的resp的body中获取数据
	c.String(http.StatusOK, "test2 OK")
}



