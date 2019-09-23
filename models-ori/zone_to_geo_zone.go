package models

import (
	"time"
)

type ZoneToGeoZone struct {
	ZoneToGeoZoneId int       `xorm:"not null pk autoincr INT(11)"`
	CountryId       int       `xorm:"not null INT(11)"`
	ZoneId          int       `xorm:"not null default 0 INT(11)"`
	GeoZoneId       int       `xorm:"not null INT(11)"`
	DateAdded       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	DateModified    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
