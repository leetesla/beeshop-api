package models

type ProductFilter struct {
	ProductId int `xorm:"not null pk INT(11)"`
	FilterId  int `xorm:"not null pk INT(11)"`
}
