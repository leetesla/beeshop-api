package models

type Extension struct {
	ExtensionId int    `xorm:"not null pk autoincr INT(11)"`
	Type        string `xorm:"not null VARCHAR(32)"`
	Code        string `xorm:"not null VARCHAR(32)"`
}
