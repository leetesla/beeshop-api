package models

import (
	"time"
)

type Reward struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId int       `xorm:"not null INT(11)"`
	DateAdded  time.Time `xorm:"comment('reward时间') DATETIME"`
	Total      string    `xorm:"not null default 0.00 comment('reward金额') DECIMAL(15,2)"`
}
