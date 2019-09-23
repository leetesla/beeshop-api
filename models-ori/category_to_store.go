package models

type CategoryToStore struct {
	CategoryId int `xorm:"not null pk INT(11)"`
	StoreId    int `xorm:"not null pk INT(11)"`
}
