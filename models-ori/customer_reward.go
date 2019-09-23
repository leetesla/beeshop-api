package models

import (
	"time"
)

type CustomerReward struct {
	CustomerRewardId int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId       int       `xorm:"not null default 0 INT(11)"`
	OrderId          int       `xorm:"not null default 0 unique INT(11)"`
	Amount           string    `xorm:"not null default 0.00 DECIMAL(15,2)"`
	CreatedAt        time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	ModifiedAt       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Status           int       `xorm:"not null default 0 TINYINT(4)"`
	Points           int       `xorm:"not null default 0 INT(11)"`
}
