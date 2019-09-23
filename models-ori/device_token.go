package models

type DeviceToken struct {
	DeviceToken string `xorm:"not null pk VARCHAR(64)"`
	TokenType   int    `xorm:"not null TINYINT(4)"`
	CustomerId  int    `xorm:"not null index INT(11)"`
}
