package models

type OrderAmazonProduct struct {
	OrderProductId      int    `xorm:"not null pk INT(11)"`
	AmazonOrderItemCode string `xorm:"not null VARCHAR(255)"`
}
