package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {

	r := gin.Default()

	r.Any("/testing", startPage)

	r.Run()
}

func startPage(c *gin.Context) {
	var person Person
	// 只绑定查询字符串
	if c.ShouldBindQuery(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	c.String(200, "Success")
}
