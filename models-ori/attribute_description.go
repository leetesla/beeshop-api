package models

type AttributeDescription struct {
	AttributeId int    `xorm:"not null pk INT(11)"`
	LanguageId  int    `xorm:"not null pk INT(11)"`
	Name        string `xorm:"not null VARCHAR(64)"`
}
