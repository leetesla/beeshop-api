package models

import (
	"time"
)

type Deposit struct {
	DepositId       int       `xorm:"not null pk autoincr INT(11)"`
	DepositNo       string    `xorm:"not null default '' VARCHAR(32)"`
	CustomerId      int       `xorm:"not null index INT(11)"`
	Amount          string    `xorm:"not null default 0.00 DECIMAL(15,2)"`
	Bonus           string    `xorm:"not null default 0.00 DECIMAL(15,2)"`
	Description     string    `xorm:"not null default '' VARCHAR(255)"`
	DepositStatusId int       `xorm:"not null default 1 index INT(11)"`
	CreatedAt       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
