package main

import "github.com/gin-gonic/gin"
import "net/http"

// 查询字符串参数
func main() {

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	router := gin.Default()

	// 简单组： v1
	v1 := router.Group("/v1")
	{
		v1.GET("/welcome", func(c *gin.Context) {
			firstname := c.DefaultQuery("firstname", "Guest")
			lastname := c.Query("lastname")
			c.String(http.StatusOK, "V1 %s %s", firstname, lastname)
		})
	}

	// 简单组： v2
	v2 := router.Group("/v2")
	{
		v2.GET("/welcome", func(c *gin.Context) {
			firstname := c.DefaultQuery("firstname", "Guest")
			lastname := c.Query("lastname")
			c.String(http.StatusOK, "V2 %s %s", firstname, lastname)
		})
	}

	router.Run()
	// 自定义端口
	// router.Run(":3000")
}



