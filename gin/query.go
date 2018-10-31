package main

import "github.com/gin-gonic/gin"
import "net/http"

// 查询字符串参数
func main() {

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	router := gin.Default()

	// 请求响应匹配的 URL： /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run()
	// 自定义端口
	// router.Run(":3000")
}



