package models

type ProductRelated struct {
	ProductId              int    `xorm:"not null pk INT(11)"`
	ProductRelateAttribute string `xorm:"VARCHAR(200)"`
}
