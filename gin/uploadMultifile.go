package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"
import "log"

// 上传多文件
func main() {

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
				log.Println(file.Filename)

				// 上传文件到指定的 dst.
				// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	router.Run()
	// 自定义端口
	// router.Run(":3000")
}



