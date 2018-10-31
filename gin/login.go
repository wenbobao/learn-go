package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 从JSON绑定
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {

	r := gin.Default()

	// 绑定 JSON 的示例 ({"user": "manu", "password": "123"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err == nil {

			if json.User == "bob" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 一个 HTML 表单绑定的示例 (user=manu&password=123)
	r.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if err := c.ShouldBind(&form); err == nil {

			if form.User == "bob" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	r.Run()
}
