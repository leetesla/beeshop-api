package models

import (
	"time"
)

type TrusteeFeeDetail struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	TrusteeFeeId int       `xorm:"not null INT(11)"`
	Amount       string    `xorm:"not null default 0.00 DECIMAL(15,2)"`
	Addtime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
