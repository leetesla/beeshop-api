package models

import (
	"time"
)

type Download struct {
	DownloadId int       `xorm:"not null pk autoincr INT(11)"`
	Filename   string    `xorm:"not null VARCHAR(128)"`
	Mask       string    `xorm:"not null VARCHAR(128)"`
	Remaining  int       `xorm:"not null default 0 INT(11)"`
	DateAdded  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
