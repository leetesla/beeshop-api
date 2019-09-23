package models

import (
	"time"
)

type PostImage struct {
	PostImageId int       `xorm:"not null pk autoincr INT(11)"`
	PostId      int       `xorm:"not null index INT(11)"`
	Image       string    `xorm:"not null VARCHAR(255)"`
	CreatedAt   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	PostReplyId int       `xorm:"not null default 0 INT(11)"`
}
