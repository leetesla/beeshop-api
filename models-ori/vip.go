package models

import (
	"time"
)

type Vip struct {
	CustomerId int       `xorm:"not null pk INT(11)"`
	Name       string    `xorm:"not null default '' VARCHAR(45)"`
	Telephone  string    `xorm:"not null default '' VARCHAR(45)"`
	CardNo     string    `xorm:"not null default '' VARCHAR(45)"`
	Type       int       `xorm:"not null comment('vip类型(公司or个人)') INT(11)"`
	CreatedAt  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	ExpiredAt  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
