package models

type OrderDownload struct {
	OrderDownloadId int    `xorm:"not null pk autoincr INT(11)"`
	OrderId         int    `xorm:"not null INT(11)"`
	OrderProductId  int    `xorm:"not null INT(11)"`
	Name            string `xorm:"not null VARCHAR(64)"`
	Filename        string `xorm:"not null VARCHAR(128)"`
	Mask            string `xorm:"not null VARCHAR(128)"`
	Remaining       int    `xorm:"not null default 0 INT(3)"`
}
