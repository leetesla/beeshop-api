package models

type Saler struct {
	Id    int    `xorm:"not null pk autoincr INT(11)"`
	Name  string `xorm:"not null default '' comment('名字') VARCHAR(32)"`
	Phone string `xorm:"comment('手机号') VARCHAR(32)"`
}
