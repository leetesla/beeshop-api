package models

type MfgrToStore struct {
	MfgrId  int `xorm:"not null pk INT(11)"`
	StoreId int `xorm:"not null pk INT(11)"`
}
