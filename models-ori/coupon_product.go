package models

type CouponProduct struct {
	CouponProductId int `xorm:"not null pk autoincr INT(11)"`
	CouponId        int `xorm:"not null INT(11)"`
	ProductId       int `xorm:"not null INT(11)"`
}
