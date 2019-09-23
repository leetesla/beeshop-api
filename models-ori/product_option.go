package models

type ProductOption struct {
	ProductOptionId int    `xorm:"not null pk autoincr INT(11)"`
	ProductId       int    `xorm:"not null INT(11)"`
	OptionId        int    `xorm:"not null INT(11)"`
	OptionValue     string `xorm:"not null TEXT"`
	Required        int    `xorm:"not null TINYINT(1)"`
}
