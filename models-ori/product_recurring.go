package models

type ProductRecurring struct {
	ProductId int `xorm:"not null pk INT(11)"`
	StoreId   int `xorm:"not null pk INT(11)"`
}
