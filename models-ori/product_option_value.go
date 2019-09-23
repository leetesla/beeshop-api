package models

type ProductOptionValue struct {
	ProductOptionValueId int    `xorm:"not null pk autoincr INT(11)"`
	ProductOptionId      int    `xorm:"not null INT(11)"`
	ProductId            int    `xorm:"not null INT(11)"`
	OptionId             int    `xorm:"not null INT(11)"`
	OptionValueId        int    `xorm:"not null INT(11)"`
	Quantity             int    `xorm:"not null INT(3)"`
	Subtract             int    `xorm:"not null TINYINT(1)"`
	Price                string `xorm:"not null DECIMAL(15,4)"`
	PricePrefix          string `xorm:"not null VARCHAR(1)"`
	Points               int    `xorm:"not null INT(8)"`
	PointsPrefix         string `xorm:"not null VARCHAR(1)"`
	Weight               string `xorm:"not null DECIMAL(15,8)"`
	WeightPrefix         string `xorm:"not null VARCHAR(1)"`
}
