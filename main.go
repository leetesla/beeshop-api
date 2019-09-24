package main

//xorm reverse mysql "root:jahaE8eTstFSla_@(116.62.213.177)/beeshop?charset=utf8" templates/goxorm
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"beeshop/controllers"
	"beeshop/app"
)

func main() {
	var err error
	app.Db,err=app.GetDb()

	defer app.Db.Close()
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}

	// 读取
	router := gin.Default()

	router.GET("/user/:id", controllers.GetUser)
	router.GET("/product/:id", controllers.GetProduct)

	router.Run(":9999")
}