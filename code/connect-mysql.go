package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil{
    fmt.Println("Error1", err)
	}
	fmt.Println("connect success")
	db.Exec("insert into user (id , name) values (?, ?)", "2", "AAA")
	db.Exec("insert into user (id , name) values (?, ?)", "3", "BBB")
}
