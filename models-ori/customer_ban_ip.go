package models

type CustomerBanIp struct {
	CustomerBanIpId int    `xorm:"not null pk autoincr INT(11)"`
	Ip              string `xorm:"not null index VARCHAR(40)"`
}
