package models

import (
	"time"
)

type Return struct {
	ReturnId       int       `xorm:"not null pk autoincr INT(11)"`
	OrderId        int       `xorm:"not null INT(11)"`
	ProductId      int       `xorm:"not null INT(11)"`
	CustomerId     int       `xorm:"not null INT(11)"`
	Firstname      string    `xorm:"not null VARCHAR(32)"`
	Lastname       string    `xorm:"not null VARCHAR(32)"`
	Email          string    `xorm:"not null VARCHAR(96)"`
	Telephone      string    `xorm:"not null VARCHAR(32)"`
	Product        string    `xorm:"not null VARCHAR(255)"`
	Model          string    `xorm:"not null VARCHAR(64)"`
	Quantity       int       `xorm:"not null INT(4)"`
	Opened         int       `xorm:"not null TINYINT(1)"`
	ReturnReasonId int       `xorm:"not null INT(11)"`
	ReturnActionId int       `xorm:"not null INT(11)"`
	ReturnStatusId int       `xorm:"not null INT(11)"`
	Comment        string    `xorm:"TEXT"`
	DateOrdered    time.Time `xorm:"not null DATE"`
	DateAdded      time.Time `xorm:"not null DATETIME"`
	DateModified   time.Time `xorm:"not null DATETIME"`
}
