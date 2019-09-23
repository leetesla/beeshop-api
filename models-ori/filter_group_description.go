package models

type FilterGroupDescription struct {
	FilterGroupId int    `xorm:"not null pk INT(11)"`
	LanguageId    int    `xorm:"not null pk INT(11)"`
	Name          string `xorm:"not null VARCHAR(64)"`
}
