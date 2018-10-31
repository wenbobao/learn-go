package main

import "github.com/gin-gonic/gin"

func main() {

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	router.Run()
	// 自定义端口
	// router.Run(":3000")
}

func getting() {
	
}


