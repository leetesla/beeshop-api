package models

import (
	"time"
)

type CustomerIp struct {
	CustomerIpId int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId   int       `xorm:"not null INT(11)"`
	Ip           string    `xorm:"not null index VARCHAR(40)"`
	DateAdded    time.Time `xorm:"not null DATETIME"`
}
