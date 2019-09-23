package models

import (
	"time"
)

type TaxClass struct {
	TaxClassId   int       `xorm:"not null pk autoincr INT(11)"`
	Title        string    `xorm:"not null VARCHAR(32)"`
	Description  string    `xorm:"not null VARCHAR(255)"`
	DateAdded    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	DateModified time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
