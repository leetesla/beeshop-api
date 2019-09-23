package models

type NotificationRead struct {
	NotificationId int `xorm:"not null pk INT(11)"`
	CustomerId     int `xorm:"not null pk index INT(11)"`
	Deleted        int `xorm:"not null default 0 TINYINT(4)"`
}
