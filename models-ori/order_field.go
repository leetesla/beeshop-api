package models

type OrderField struct {
	OrderId            int    `xorm:"not null pk INT(11)"`
	CustomFieldId      int    `xorm:"not null pk INT(11)"`
	CustomFieldValueId int    `xorm:"not null pk INT(11)"`
	Name               int    `xorm:"not null INT(128)"`
	Value              string `xorm:"not null TEXT"`
	SortOrder          int    `xorm:"not null INT(3)"`
}
