package models

type ProductToLayout struct {
	ProductId int `xorm:"not null pk INT(11)"`
	StoreId   int `xorm:"not null pk INT(11)"`
	LayoutId  int `xorm:"not null INT(11)"`
}
