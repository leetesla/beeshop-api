package models

type CouponCategory struct {
	CouponId   int `xorm:"not null pk INT(11)"`
	CategoryId int `xorm:"not null pk INT(11)"`
}
