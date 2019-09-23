package models

type CustomerGroupDescription struct {
	CustomerGroupId int    `xorm:"not null pk INT(11)"`
	LanguageId      int    `xorm:"not null pk INT(11)"`
	Name            string `xorm:"not null VARCHAR(32)"`
	Description     string `xorm:"not null TEXT"`
}
