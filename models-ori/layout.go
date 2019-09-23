package models

type Layout struct {
	LayoutId int    `xorm:"not null pk autoincr INT(11)"`
	Name     string `xorm:"not null VARCHAR(64)"`
}
