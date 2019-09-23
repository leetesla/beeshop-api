package models

type OptionValueDescription struct {
	OptionValueId int    `xorm:"not null pk INT(11)"`
	LanguageId    int    `xorm:"not null pk INT(11)"`
	OptionId      int    `xorm:"not null INT(11)"`
	Name          string `xorm:"not null VARCHAR(128)"`
}
