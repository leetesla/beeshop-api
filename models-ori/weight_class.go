package models

type WeightClass struct {
	WeightClassId int    `xorm:"not null pk autoincr INT(11)"`
	Value         string `xorm:"not null default 0.00000000 DECIMAL(15,8)"`
}
