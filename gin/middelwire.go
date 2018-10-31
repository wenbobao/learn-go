package main

import "github.com/gin-gonic/gin"

func main() {

	// Default 已经连接了 Logger 和 Recovery 中间件
	// r := gin.Default()

	// 没有中间件
	r := gin.New()

	// 全局中间件
	// Logger 中间件将写日志到 gin.DefaultWriter ,即使你设置 GIN_MODE=release 。
	// 默认 gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件从任何 panic 恢复，如果出现 panic，它会写一个 500 错误。
	r.Use(gin.Recovery())
		
	r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
					"message": "pong",
			})
	})
  r.Run()
}


