package models

type Information struct {
	InformationId int `xorm:"not null pk autoincr INT(11)"`
	Bottom        int `xorm:"not null default 0 INT(1)"`
	SortOrder     int `xorm:"not null default 0 INT(3)"`
	Status        int `xorm:"not null default 1 TINYINT(1)"`
}
