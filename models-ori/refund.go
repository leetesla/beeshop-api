package models

import (
	"time"
)

type Refund struct {
	RefundId          int       `xorm:"not null pk autoincr INT(11)"`
	OrderId           int       `xorm:"not null index INT(11)"`
	Total             string    `xorm:"not null default 0.00 DECIMAL(15,2)"`
	RefundStatusId    int       `xorm:"not null default 1 INT(11)"`
	CreatedAt         time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt         time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	SubTotal          string    `xorm:"not null default 0.00 DECIMAL(15,2)"`
	Shipping          float32   `xorm:"not null default 0 FLOAT"`
	AdminRefundStatus int       `xorm:"not null default 1 comment('后台退款状态') INT(11)"`
	RefundType        string    `xorm:"not null default '0' comment('退款类型:0 用户的退款; 1 后台添加的退款') VARCHAR(45)"`
	OrderStatusId     int       `xorm:"not null default 0 INT(11)"`
}
