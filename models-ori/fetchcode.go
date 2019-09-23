package models

import (
	"time"
)

type Fetchcode struct {
	Barcode     string    `xorm:"not null pk CHAR(13)"`
	CustomerId  int       `xorm:"not null index INT(11)"`
	Used        int       `xorm:"not null default 0 TINYINT(1)"`
	DateCreated time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	DateUsed    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
