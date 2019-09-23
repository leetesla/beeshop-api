package models

type LengthClass struct {
	LengthClassId int    `xorm:"not null pk autoincr INT(11)"`
	Value         string `xorm:"not null DECIMAL(15,8)"`
}
