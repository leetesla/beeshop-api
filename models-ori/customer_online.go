package models

import (
	"time"
)

type CustomerOnline struct {
	Ip         string    `xorm:"not null pk VARCHAR(40)"`
	CustomerId int       `xorm:"not null INT(11)"`
	Url        string    `xorm:"not null TEXT"`
	Referer    string    `xorm:"not null TEXT"`
	DateAdded  time.Time `xorm:"not null DATETIME"`
}
