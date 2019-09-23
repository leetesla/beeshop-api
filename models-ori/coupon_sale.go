package models

type CouponSale struct {
	CouponSaleId int `xorm:"not null pk autoincr INT(11)"`
	CouponId     int `xorm:"not null INT(11)"`
	SaleId       int `xorm:"not null INT(11)"`
}
