package models

import (
	"time"
)

type StockLiquidationHistory struct {
	Id            int       `xorm:"not null pk autoincr INT(11)"`
	LiquidationId int       `xorm:"not null INT(11)"`
	Stock         int       `xorm:"not null INT(11)"`
	RealStock     int       `xorm:"not null INT(11)"`
	Addtime       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
