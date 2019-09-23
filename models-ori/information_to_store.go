package models

type InformationToStore struct {
	InformationId int `xorm:"not null pk INT(11)"`
	StoreId       int `xorm:"not null pk INT(11)"`
}
