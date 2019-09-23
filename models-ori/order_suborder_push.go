package models

import (
	"time"
)

type OrderSuborderPush struct {
	OrderId      int       `xorm:"not null pk INT(11)"`
	StoreHouseId int       `xorm:"not null pk TINYINT(4)"`
	PushStatus   int       `xorm:"TINYINT(4)"`
	ErrorCode    int       `xorm:"INT(4)"`
	ErrorMsg     string    `xorm:"VARCHAR(200)"`
	DateModify   time.Time `xorm:"DATETIME"`
}
