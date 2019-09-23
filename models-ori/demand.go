package models

import (
	"time"
)

type Demand struct {
	DemandId     int       `xorm:"not null pk autoincr INT(11)"`
	Type         int       `xorm:"not null default 0 comment('0 货物；1 物料') TINYINT(1)"`
	Status       int       `xorm:"not null default 0 TINYINT(1)"`
	Promoter     int       `xorm:"not null comment('需求发起人') INT(11)"`
	Accepter     int       `xorm:"not null comment('需求接受人') INT(11)"`
	WarehouseOut int       `xorm:"not null default 0 INT(11)"`
	WarehouseIn  int       `xorm:"not null INT(11)"`
	Remark       string    `xorm:"not null default '' VARCHAR(200)"`
	Addtime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
