package models

type CategoryToLayout struct {
	CategoryId int `xorm:"not null pk INT(11)"`
	StoreId    int `xorm:"not null pk INT(11)"`
	LayoutId   int `xorm:"not null INT(11)"`
}
