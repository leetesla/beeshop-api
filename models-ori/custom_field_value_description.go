package models

type CustomFieldValueDescription struct {
	CustomFieldValueId int    `xorm:"not null pk INT(11)"`
	LanguageId         int    `xorm:"not null pk INT(11)"`
	CustomFieldId      int    `xorm:"not null INT(11)"`
	Name               string `xorm:"not null VARCHAR(128)"`
}
