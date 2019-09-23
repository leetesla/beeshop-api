package models

type DiscountActivity struct {
	Id                int    `xorm:"not null pk autoincr INT(11)"`
	ActivityId        int    `xorm:"comment('活动id主键') INT(11)"`
	ProductId         int    `xorm:"comment('商品id主键') INT(11)"`
	DiscountSalePrice string `xorm:"comment('活动零售价') DECIMAL(15,2)"`
	DiscountVipPrice  string `xorm:"comment('活动vip价') DECIMAL(15,2)"`
}
