package models

type CategoryDescription struct {
	CategoryId      int    `xorm:"not null pk INT(11)"`
	LanguageId      int    `xorm:"not null pk INT(11)"`
	Name            string `xorm:"not null index VARCHAR(255)"`
	Description     string `xorm:"not null TEXT"`
	MetaDescription string `xorm:"not null VARCHAR(255)"`
	MetaKeyword     string `xorm:"not null VARCHAR(255)"`
}
