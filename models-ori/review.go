package models

import (
	"time"
)

type Review struct {
	ReviewId     int       `xorm:"not null pk autoincr INT(11)"`
	ProductId    int       `xorm:"not null index INT(11)"`
	CustomerId   int       `xorm:"not null INT(11)"`
	Author       string    `xorm:"not null VARCHAR(64)"`
	Text         string    `xorm:"not null TEXT"`
	Rating       int       `xorm:"not null default 5 INT(1)"`
	Status       int       `xorm:"not null default 0 TINYINT(1)"`
	DateAdded    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	DateModified time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	OrderId      int       `xorm:"INT(32)"`
	Image1       string    `xorm:"VARCHAR(64)"`
	Image2       string    `xorm:"VARCHAR(64)"`
	Image3       string    `xorm:"VARCHAR(64)"`
	Image4       string    `xorm:"VARCHAR(64)"`
	Image5       string    `xorm:"VARCHAR(64)"`
}
