package models

import (
	"time"
)

type OrderHistory struct {
	OrderHistoryId int       `xorm:"not null pk autoincr INT(11)"`
	OrderId        int       `xorm:"not null INT(11)"`
	OrderStatusId  int       `xorm:"not null INT(5)"`
	Notify         int       `xorm:"not null default 0 TINYINT(1)"`
	NotifyMessage  string    `xorm:"TEXT"`
	Comment        string    `xorm:"not null TEXT"`
	DateAdded      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	StoreHouseId   int       `xorm:"default 1 TINYINT(4)"`
}
