package models

import (
	"time"
)

type PostReply struct {
	PostReplyId int       `xorm:"not null pk autoincr INT(11)"`
	PostId      int       `xorm:"not null index INT(11)"`
	Text        string    `xorm:"not null VARCHAR(255)"`
	ReplyerId   int       `xorm:"not null index INT(11)"`
	ReplyeeId   int       `xorm:"not null INT(11)"`
	CreatedAt   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Status      int       `xorm:"not null default 0 index INT(11)"`
	Top         int       `xorm:"not null default 0 comment('0 ：不置顶；
1 ：置顶') TINYINT(1)"`
}
