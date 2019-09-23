package models

type TaxRule struct {
	TaxRuleId  int    `xorm:"not null pk autoincr INT(11)"`
	TaxClassId int    `xorm:"not null INT(11)"`
	TaxRateId  int    `xorm:"not null INT(11)"`
	Based      string `xorm:"not null VARCHAR(10)"`
	Priority   int    `xorm:"not null default 1 INT(5)"`
}
