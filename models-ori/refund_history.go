package models

import (
	"time"
)

type RefundHistory struct {
	RefundHistoryId   int       `xorm:"not null pk autoincr INT(11)"`
	RefundId          int       `xorm:"not null INT(11)"`
	RefundStatusId    int       `xorm:"not null INT(11)"`
	Notify            int       `xorm:"not null default 0 TINYINT(4)"`
	Comment           string    `xorm:"not null default '' VARCHAR(100)"`
	CreatedAt         time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	AdminRefundStatus string    `xorm:"not null default '1' VARCHAR(45)"`
}
