package models

import (
	"time"
)

type BannerImage struct {
	BannerImageId int       `xorm:"not null pk autoincr INT(11)"`
	BannerId      int       `xorm:"not null INT(11)"`
	Link          string    `xorm:"not null VARCHAR(255)"`
	Image         string    `xorm:"not null VARCHAR(255)"`
	SortOrder     int       `xorm:"not null default 0 INT(11)"`
	DateStart     time.Time `xorm:"DATETIME"`
	DateEnd       time.Time `xorm:"DATETIME"`
}
