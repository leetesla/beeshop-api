package models

import (
	"time"
)

type CustomerHistory struct {
	CustomerHistoryId int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId        int       `xorm:"not null INT(11)"`
	Comment           string    `xorm:"not null TEXT"`
	DateAdded         time.Time `xorm:"not null DATETIME"`
}
