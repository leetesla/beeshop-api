package models

type OrderAmazon struct {
	OrderId       int    `xorm:"not null pk INT(11)"`
	AmazonOrderId string `xorm:"not null index VARCHAR(255)"`
	FreeShipping  int    `xorm:"not null default 0 TINYINT(4)"`
}
