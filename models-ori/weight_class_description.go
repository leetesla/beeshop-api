package models

type WeightClassDescription struct {
	WeightClassId int    `xorm:"not null pk autoincr INT(11)"`
	LanguageId    int    `xorm:"not null pk INT(11)"`
	Title         string `xorm:"not null VARCHAR(32)"`
	Unit          string `xorm:"not null VARCHAR(4)"`
}
