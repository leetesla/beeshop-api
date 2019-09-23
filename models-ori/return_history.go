package models

import (
	"time"
)

type ReturnHistory struct {
	ReturnHistoryId int       `xorm:"not null pk autoincr INT(11)"`
	ReturnId        int       `xorm:"not null INT(11)"`
	ReturnStatusId  int       `xorm:"not null INT(11)"`
	Notify          int       `xorm:"not null TINYINT(1)"`
	Comment         string    `xorm:"not null TEXT"`
	DateAdded       time.Time `xorm:"not null DATETIME"`
}
