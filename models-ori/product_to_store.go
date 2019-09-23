package models

type ProductToStore struct {
	ProductId int `xorm:"not null pk INT(11)"`
	StoreId   int `xorm:"not null pk default 0 INT(11)"`
}
