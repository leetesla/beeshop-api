package models

type TaxRateToCustomerGroup struct {
	TaxRateId       int `xorm:"not null pk INT(11)"`
	CustomerGroupId int `xorm:"not null pk INT(11)"`
}
