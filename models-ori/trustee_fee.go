package models

import (
	"time"
)

type TrusteeFee struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	OrderId    int       `xorm:"not null INT(11)"`
	CustomerId int       `xorm:"not null INT(11)"`
	Remark     string    `xorm:"not null default '' VARCHAR(200)"`
	Amount     string    `xorm:"not null default 0.00 comment('应缴') DECIMAL(15,2)"`
	Paid       string    `xorm:"not null default 0.00 comment('实缴') DECIMAL(15,2)"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Addtime    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
