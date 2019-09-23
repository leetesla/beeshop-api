package models

type CustomFieldToCustomerGroup struct {
	CustomFieldId   int `xorm:"not null pk INT(11)"`
	CustomerGroupId int `xorm:"not null pk INT(11)"`
}
