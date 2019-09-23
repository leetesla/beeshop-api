package models

import (
	"time"
)

type StockLiquidation struct {
	LiquidationId int       `xorm:"not null pk autoincr INT(11)"`
	ProductId     int       `xorm:"not null INT(11)"`
	WarehouseId   int       `xorm:"not null INT(11)"`
	Stock         int       `xorm:"not null comment('管家记录的库存') INT(11)"`
	RealStock     int       `xorm:"not null comment('清点出的实际库存') INT(11)"`
	Addtime       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
