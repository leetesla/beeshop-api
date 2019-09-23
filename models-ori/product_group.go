package models

type ProductGroup struct {
	Id      int    `xorm:"not null pk autoincr INT(11)"`
	Name    string `xorm:"not null VARCHAR(128)"`
	SpecId1 int    `xorm:"not null default 0 INT(11)"`
	SpecId2 int    `xorm:"not null default 0 INT(11)"`
	SpecId3 int    `xorm:"not null default 0 INT(11)"`
	SpecId4 int    `xorm:"not null default 0 INT(11)"`
	SpecId5 int    `xorm:"not null default 0 INT(11)"`
}
