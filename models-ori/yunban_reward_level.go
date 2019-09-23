package models

type YunbanRewardLevel struct {
	Id      int    `xorm:"not null pk autoincr INT(11)"`
	Limit   int    `xorm:"not null INT(11)"`
	Percent string `xorm:"not null default 0.00 DECIMAL(15,2)"`
}
