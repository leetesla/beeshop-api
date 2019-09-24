package controllers

import (
	"beeshop/models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"beeshop/app"
	"beeshop/app"
)

func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product

	if err := app.Db.Where("product_id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, product)
	}
}
