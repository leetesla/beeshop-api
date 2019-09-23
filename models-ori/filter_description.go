package models

type FilterDescription struct {
	FilterId      int    `xorm:"not null pk INT(11)"`
	LanguageId    int    `xorm:"not null pk INT(11)"`
	FilterGroupId int    `xorm:"not null INT(11)"`
	Name          string `xorm:"not null VARCHAR(64)"`
}
