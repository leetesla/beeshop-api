package main

//xorm reverse mysql "root:jahaE8eTstFSla_@(116.62.213.177)/beeshop?charset=utf8" templates/goxorm
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"beeshop/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := gorm.Open("mysql", "root:jahaE8eTstFSla_@(116.62.213.177)/beeshop?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		fmt.Println("连接数据库成功")
	}

	app.db=db

	// 读取
	router := gin.Default()

	router.GET("/stores/:id", store)
	router.GET("/user/:id", GetUser)
	router.GET("/store/:id", GetStoreHouse)
	router.GET("/product/:id", GetProduct)

	router.Run(":9999")
}

func store(c *gin.Context) {
	id := c.Params.ByName("id")
	db, err := gorm.Open("mysql", "root:jahaE8eTstFSla_@(116.62.213.177)/beeshop?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		fmt.Println("连接数据库成功")
	}

	var storeHouse models.StoreHouse

	if err := db.Where("store_house_id=?", id).First(&storeHouse).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, storeHouse)
	}
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := app.db.Where("user_id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

func GetStoreHouse(c *gin.Context) {
	// id := c.Params.ByName("id")
	var storeHouse models.StoreHouse

	// db.Where("store_id=?", id).First(&storeHouse)
	app.db.Where("store_house_id=?", 2).First(&storeHouse)
	fmt.Println(storeHouse)

	//if err := db.Where("store_id=?", id).First(&storeHouse).Error; err != nil {
	//	c.AbortWithStatus(404)
	//	fmt.Println(err)
	//} else {
	//	c.JSON(200, storeHouse)
	//}
	c.JSON(200, storeHouse)
}

func GetProduct(c *gin.Context) {
	// id := c.Params.ByName("id")
	var product models.Product
	app.db.First(&product, 2) // find product with id 1
	//db.First(&product, "code = ?", "L1212") // find product with code l1212
	fmt.Println(product)
	// c.JSON(200, product)

	// db.Where("store_id=?", id).First(&storeHouse)
	// db.First(&storeHouse)
	// fmt.Println(storeHouse)

	//if err := db.Where("store_id=?", id).First(&storeHouse).Error; err != nil {
	//	c.AbortWithStatus(404)
	//	fmt.Println(err)
	//} else {
	//	c.JSON(200, storeHouse)
	//}
}

type App struct {
	db *gorm.DB
}

var app = App{}