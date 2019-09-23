package models

import (
	"time"
)

type Withdrawal struct {
	WithdrawalId       int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId         int       `xorm:"not null INT(11)"`
	Amount             string    `xorm:"not null default 0.00 DECIMAL(15,2)"`
	Account            string    `xorm:"not null VARCHAR(45)"`
	Name               string    `xorm:"not null default '' VARCHAR(15)"`
	WithdrawalStatusId int       `xorm:"not null default 0 INT(11)"`
	CreatedAt          time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt          time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
