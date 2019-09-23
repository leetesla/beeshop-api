package models

import (
	"time"
)

type Notification struct {
	NotificationId int       `xorm:"not null pk autoincr INT(11)"`
	TaskId         string    `xorm:"not null default '' VARCHAR(60)"`
	CustomerId     int       `xorm:"not null default 0 index INT(11)"`
	Type           int       `xorm:"not null TINYINT(4)"`
	Title          string    `xorm:"not null default '' VARCHAR(45)"`
	Description    string    `xorm:"not null default '' VARCHAR(255)"`
	Link           string    `xorm:"not null default '' VARCHAR(255)"`
	Status         int       `xorm:"not null default 1 TINYINT(4)"`
	CreatedAt      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
