package models

type OrderTotalTax struct {
	OrderTotalId int    `xorm:"not null pk default 0 INT(11)"`
	Code         string `xorm:"VARCHAR(255)"`
	Tax          string `xorm:"not null DECIMAL(10,4)"`
}
