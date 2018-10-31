package main

import "github.com/gin-gonic/gin"
import "net/http"

// Multipart/Urlencoded 表单

func main() {

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
    nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
				"status":  "posted",
				"message": message,
				"nick":    nick,
		})
	})

	router.Run()
	// 自定义端口
	// router.Run(":3000")
}



