package models

type InformationDescription struct {
	InformationId int    `xorm:"not null pk INT(11)"`
	LanguageId    int    `xorm:"not null pk INT(11)"`
	Title         string `xorm:"not null VARCHAR(64)"`
	Description   string `xorm:"not null TEXT"`
	CategoryId    int    `xorm:"not null default 1 INT(11)"`
	HeaderImage   string `xorm:"VARCHAR(255)"`
	Recommend     int    `xorm:"default 0 TINYINT(4)"`
}
