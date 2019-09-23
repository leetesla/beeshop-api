package models

type SalerToCustomer struct {
	SalerId    int `xorm:"not null pk INT(11)"`
	CustomerId int `xorm:"not null pk INT(11)"`
}
