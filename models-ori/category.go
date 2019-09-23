package models

import (
	"time"
)

type Category struct {
	CategoryId   int       `xorm:"not null pk autoincr INT(11)"`
	Image        string    `xorm:"VARCHAR(255)"`
	ParentId     int       `xorm:"not null default 0 INT(11)"`
	Hot          int       `xorm:"not null default 0 TINYINT(4)"`
	SortOrder    int       `xorm:"not null default 0 INT(3)"`
	Status       int       `xorm:"not null TINYINT(1)"`
	DateAdded    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	DateModified time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Type         int       `xorm:"not null default 0 comment('分类产品类型，自提，直邮') TINYINT(4)"`
}
