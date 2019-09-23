package models

type OrderOption struct {
	OrderOptionId        int    `xorm:"not null pk autoincr INT(11)"`
	OrderId              int    `xorm:"not null INT(11)"`
	OrderProductId       int    `xorm:"not null INT(11)"`
	ProductOptionId      int    `xorm:"not null INT(11)"`
	ProductOptionValueId int    `xorm:"not null default 0 INT(11)"`
	Name                 string `xorm:"not null VARCHAR(255)"`
	Value                string `xorm:"not null TEXT"`
	Type                 string `xorm:"not null VARCHAR(32)"`
}
