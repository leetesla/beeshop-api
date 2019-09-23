package models

import (
	"time"
)

type RefundProduct struct {
	RefundProductId int       `xorm:"not null pk autoincr INT(11)"`
	RefundId        int       `xorm:"not null index INT(11)"`
	OrderProductId  int       `xorm:"not null default 0 INT(11)"`
	OrderId         int       `xorm:"not null index INT(11)"`
	ProductId       int       `xorm:"not null INT(11)"`
	Price           string    `xorm:"not null DECIMAL(15,2)"`
	Quantity        int       `xorm:"not null INT(11)"`
	Comment         string    `xorm:"not null default '' VARCHAR(100)"`
	CreatedAt       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
