package models

import (
	"time"
)

type DepositPlan struct {
	DepositPlanId int       `xorm:"not null pk autoincr INT(11)"`
	Name          string    `xorm:"not null VARCHAR(50)"`
	Enabled       int       `xorm:"not null default 0 TINYINT(4)"`
	CreatedAt     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
