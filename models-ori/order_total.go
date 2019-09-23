package models

type OrderTotal struct {
	OrderTotalId int    `xorm:"not null pk autoincr INT(10)"`
	OrderId      int    `xorm:"not null index INT(11)"`
	Code         string `xorm:"not null VARCHAR(32)"`
	Title        string `xorm:"not null VARCHAR(255)"`
	Text         string `xorm:"not null VARCHAR(255)"`
	Value        string `xorm:"not null default 0.00 DECIMAL(15,2)"`
	SortOrder    int    `xorm:"not null INT(3)"`
}
