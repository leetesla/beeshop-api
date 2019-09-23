package models

type TrackCompanys struct {
	Id          int    `xorm:"not null pk autoincr INT(32)"`
	CompanyId   string `xorm:"not null default '' VARCHAR(32)"`
	CompanyName string `xorm:"VARCHAR(32)"`
}
