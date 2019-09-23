package models

import (
	"time"
)

type CustomerTransaction struct {
	CustomerTransactionId int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId            int       `xorm:"not null index INT(11)"`
	OrderId               int       `xorm:"not null index INT(11)"`
	Description           string    `xorm:"not null TEXT"`
	Amount                string    `xorm:"not null DECIMAL(15,2)"`
	DateAdded             time.Time `xorm:"not null DATETIME"`
}
