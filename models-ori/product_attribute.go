package models

type ProductAttribute struct {
	ProductId   int    `xorm:"not null pk INT(11)"`
	AttributeId int    `xorm:"not null pk INT(11)"`
	LanguageId  int    `xorm:"not null pk INT(11)"`
	Text        string `xorm:"not null TEXT"`
}
