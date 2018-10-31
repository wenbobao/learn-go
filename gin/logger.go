package main

import "github.com/gin-gonic/gin"
import "os"
import "io"

func main() {
	// 禁用控制台颜色，当你将日志写入到文件的时候，你不需要控制台颜色。
	gin.DisableConsoleColor()

	// 写入日志的文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果你需要同时写入日志文件和控制台上显示，使用下面代码
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
		
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
					"message": "pong",
			})
	})
  r.Run()
}


