package models

import (
	"time"
)

type Requisition struct {
	RequisitionId int       `xorm:"not null pk autoincr INT(11)"`
	BillNo        string    `xorm:"not null VARCHAR(45)"`
	WarehouseOut  int       `xorm:"not null INT(11)"`
	WarehouseIn   int       `xorm:"not null INT(11)"`
	Status        int       `xorm:"not null INT(11)"`
	CreateTime    time.Time `xorm:"not null DATETIME"`
	Remark        string    `xorm:"default '' VARCHAR(200)"`
	Addtime       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('记录添加时间') DATETIME"`
}
