package models

import (
	"time"
)

type PostLike struct {
	PostLikeId int       `xorm:"not null pk autoincr INT(11)"`
	PostId     int       `xorm:"not null index INT(11)"`
	ReplyerId  int       `xorm:"not null index INT(11)"`
	CreatedAt  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
