package models

import (
	"time"
)

type Post struct {
	PostId     int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId int       `xorm:"not null index INT(11)"`
	Type       int       `xorm:"not null INT(11)"`
	Title      string    `xorm:"not null VARCHAR(128)"`
	Text       string    `xorm:"not null VARCHAR(1000)"`
	Status     int       `xorm:"not null default 0 index INT(11)"`
	CreatedAt  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	PostType   int       `xorm:"not null default 0 comment('0 : 普通用户发帖；
1: 官方发帖') TINYINT(1)"`
	Top int `xorm:"not null default 0 comment('0 ：不置顶；
1 ：置顶') TINYINT(1)"`
}
