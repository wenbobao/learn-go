package main

import "github.com/gin-gonic/gin"
import "fmt"

// Post
func main() {

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	router.Run()
	// 自定义端口
	// router.Run(":3000")
}



