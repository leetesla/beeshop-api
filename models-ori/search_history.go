package models

import (
	"time"
)

type SearchHistory struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId int       `xorm:"not null default 0 INT(11)"`
	Keyword    string    `xorm:"default '' comment('搜索关键词') VARCHAR(45)"`
	Addtime    time.Time `xorm:"comment('搜索时间') DATETIME"`
}
