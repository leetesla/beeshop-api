package models

import (
	"time"
)

type CouponHistory struct {
	CouponHistoryId int       `xorm:"not null pk autoincr INT(11)"`
	CouponId        int       `xorm:"not null index INT(11)"`
	Code            string    `xorm:"not null pk default '' VARCHAR(12)"`
	OrderId         int       `xorm:"INT(11)"`
	CustomerId      int       `xorm:"index INT(11)"`
	Amount          string    `xorm:"DECIMAL(15,4)"`
	DateAdded       time.Time `xorm:"not null DATETIME"`
	DateUsed        time.Time `xorm:"DATETIME"`
}
