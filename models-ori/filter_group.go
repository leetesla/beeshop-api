package models

type FilterGroup struct {
	FilterGroupId int `xorm:"not null pk autoincr INT(11)"`
	SortOrder     int `xorm:"not null INT(3)"`
}
