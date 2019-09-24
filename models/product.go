package models

import (
	"time"
)

type Product struct {
	ProductId            int       `json:"product_id" xorm:"not null pk autoincr INT(11)"`
	Type                 int       `xorm:"not null default 0 comment('商品类型：0普通商品，5云伴分销商品') TINYINT(3)"`
	Model                string    `xorm:"not null VARCHAR(64)"`
	Upc                  string    `xorm:"not null VARCHAR(20)"`
	Ean                  string    `xorm:"not null VARCHAR(14)"`
	Jan                  string    `xorm:"not null VARCHAR(13)"`
	Mpn                  string    `xorm:"not null VARCHAR(64)"`
	Location             string    `xorm:"not null VARCHAR(128)"`
	Quantity             int       `xorm:"not null default 0 INT(4)"`
	StockStatusId        int       `xorm:"not null INT(11)"`
	Image                string    `xorm:"VARCHAR(255)"`
	MfgrId               int       `xorm:"not null INT(11)"`
	Shipping             int       `xorm:"not null default 0 TINYINT(1)"`
	Status               int       `xorm:"not null default 0 TINYINT(1)"`
	Price                string    `xorm:"not null default 0.00 DECIMAL(15,2)"`
	SalePrice            string    `xorm:"DECIMAL(15,2)"`
	VipPrice             string    `xorm:"not null default 0.00 DECIMAL(15,2)"`
	OfflineSalePrice     string    `xorm:"not null default 0.00 DECIMAL(15,2)"`
	OfflineVipPrice      string    `xorm:"not null default 0.00 DECIMAL(15,2)"`
	Reward               int       `xorm:"not null default 0 INT(8)"`
	TaxClassId           int       `xorm:"not null INT(11)"`
	DateAvailable        time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Weight               string    `xorm:"not null default 0.00000000 DECIMAL(15,8)"`
	WeightClassId        int       `xorm:"not null default 0 INT(11)"`
	Length               string    `xorm:"not null default 0.00000000 DECIMAL(15,8)"`
	Width                string    `xorm:"not null default 0.00000000 DECIMAL(15,8)"`
	Height               string    `xorm:"not null default 0.00000000 DECIMAL(15,8)"`
	LengthClassId        int       `xorm:"not null default 0 INT(11)"`
	Subtract             int       `xorm:"not null default 1 TINYINT(1)"`
	Minimum              int       `xorm:"not null default 1 INT(11)"`
	SortOrder            int       `xorm:"not null default 0 INT(11)"`
	Viewed               int       `xorm:"not null default 0 INT(5)"`
	Description          string    `xorm:"CHAR(255)"`
	ProductShortDesc     string    `xorm:"CHAR(255)"`
	DisplayPurchasedBase int       `xorm:"not null default 0 INT(11)"`
	StoreHouseId         int       `xorm:"default 1 TINYINT(4)"`
	GroupId              int       `xorm:"not null default 0 comment('属于哪个group组') INT(32)"`
	CreatedAt            int       `xorm:"not null INT(11)"`
	UpdatedAt            int       `xorm:"not null INT(11)"`
}

func (p Product) TableName() string {
	return "product"
	//if u.Role == "admin" {
	//	return "admin_users"
	//} else {
	//	return "users"
	//}
	// return "profiles"
}