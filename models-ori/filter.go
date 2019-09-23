package models

type Filter struct {
	FilterId      int `xorm:"not null pk autoincr INT(11)"`
	FilterGroupId int `xorm:"not null INT(11)"`
	SortOrder     int `xorm:"not null INT(3)"`
}
