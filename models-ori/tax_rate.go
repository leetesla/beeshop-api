package models

import (
	"time"
)

type TaxRate struct {
	TaxRateId    int       `xorm:"not null pk autoincr INT(11)"`
	GeoZoneId    int       `xorm:"not null default 0 INT(11)"`
	Name         string    `xorm:"not null VARCHAR(32)"`
	Rate         string    `xorm:"not null default 0.0000 DECIMAL(15,4)"`
	Type         string    `xorm:"not null CHAR(1)"`
	DateAdded    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	DateModified time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
