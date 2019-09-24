package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var Db *gorm.DB

func GetDb() (*gorm.DB, error) {
	var err error

	Db, err = gorm.Open("mysql", "root:jahaE8eTstFSla_@(116.62.213.177)/beeshop?charset=utf8&parseTime=True&loc=Local")
	// defer db.Close()

	if err != nil {
		// fmt.Println("连接失败")
		fmt.Println(err)
	}

	return Db, err
}
