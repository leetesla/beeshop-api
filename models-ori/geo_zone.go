package models

import (
	"time"
)

type GeoZone struct {
	GeoZoneId    int       `xorm:"not null pk autoincr INT(11)"`
	Name         string    `xorm:"not null VARCHAR(32)"`
	Description  string    `xorm:"not null VARCHAR(255)"`
	DateModified time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	DateAdded    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
