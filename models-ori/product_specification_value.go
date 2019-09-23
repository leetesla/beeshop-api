package models

type ProductSpecificationValue struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	ProductId  int    `xorm:"not null INT(11)"`
	SpecValue1 string `xorm:"not null default '' VARCHAR(64)"`
	SpecValue2 string `xorm:"not null default '' VARCHAR(64)"`
	SpecValue3 string `xorm:"not null default '' VARCHAR(64)"`
	SpecValue4 string `xorm:"not null default '' VARCHAR(64)"`
	SpecValue5 string `xorm:"not null default '' VARCHAR(64)"`
}
