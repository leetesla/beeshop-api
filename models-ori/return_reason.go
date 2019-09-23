package models

type ReturnReason struct {
	ReturnReasonId int    `xorm:"not null pk autoincr INT(11)"`
	LanguageId     int    `xorm:"not null pk default 0 INT(11)"`
	Name           string `xorm:"not null VARCHAR(128)"`
}
