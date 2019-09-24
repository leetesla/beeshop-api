package controllers

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"beeshop/models"
	"github.com/gin-gonic/gin"
	"beeshop/app"
)

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := app.Db.Where("user_id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		resp := map[string]interface{}{
			"code":0,
			"data":[]string{"first","s"},
			"msg":"success",
		}
		c.JSON(200, resp)
	}
}
