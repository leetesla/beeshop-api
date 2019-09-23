package models

type ProductToCategory struct {
	ProductId  int `xorm:"not null pk INT(11)"`
	CategoryId int `xorm:"not null pk INT(11)"`
}
