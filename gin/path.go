package main

import "github.com/gin-gonic/gin"
import "net/http"

// Path中的参数
func main() {

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	router := gin.Default()

	// 这个处理器将去匹配 /user/john ， 但是它不会去匹配 /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 然而，这个将会匹配 /user/john 并且也会匹配 /user/john/send
	// 如果没有其他的路由匹配 /user/john ， 它将重定向到 /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
	})

	router.Run()
	// 自定义端口
	// router.Run(":3000")
}



