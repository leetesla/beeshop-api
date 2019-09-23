package models

import (
	"time"
)

type PostProduct struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	PostId      int       `xorm:"not null default 0 INT(11)"`
	PostReplyId int       `xorm:"not null default 0 INT(11)"`
	ProductId   int       `xorm:"not null INT(11)"`
	Addtime     time.Time `xorm:"DATETIME"`
}
