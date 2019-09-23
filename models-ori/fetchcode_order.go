package models

type FetchcodeOrder struct {
	Id      int    `xorm:"not null pk autoincr INT(11)"`
	Barcode string `xorm:"not null index CHAR(13)"`
	OrderId int    `xorm:"not null index INT(11)"`
}
