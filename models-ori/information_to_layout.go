package models

type InformationToLayout struct {
	InformationId int `xorm:"not null pk INT(11)"`
	StoreId       int `xorm:"not null pk INT(11)"`
	LayoutId      int `xorm:"not null INT(11)"`
}
