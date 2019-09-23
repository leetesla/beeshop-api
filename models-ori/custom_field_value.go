package models

type CustomFieldValue struct {
	CustomFieldValueId int `xorm:"not null pk autoincr INT(11)"`
	CustomFieldId      int `xorm:"not null INT(11)"`
	SortOrder          int `xorm:"not null INT(3)"`
}
