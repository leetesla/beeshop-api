package models

import (
	"time"
)

type OrderRecurringTransaction struct {
	OrderRecurringTransactionId int       `xorm:"not null pk autoincr INT(11)"`
	OrderRecurringId            int       `xorm:"not null INT(11)"`
	Created                     time.Time `xorm:"not null DATETIME"`
	Amount                      string    `xorm:"not null DECIMAL(10,4)"`
	Type                        string    `xorm:"not null VARCHAR(255)"`
}
