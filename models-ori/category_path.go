package models

type CategoryPath struct {
	CategoryId int `xorm:"not null pk INT(11)"`
	PathId     int `xorm:"not null pk INT(11)"`
	Level      int `xorm:"not null INT(11)"`
}
