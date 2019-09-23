package models

import (
	"time"
)

type Security struct {
	Id    int       `xorm:"not null pk autoincr comment('序列号') INT(11)"`
	Code  string    `xorm:"default '' comment('防伪码') VARCHAR(32)"`
	Total int       `xorm:"default 00000000000 comment('查询次数') INT(11)"`
	Last  time.Time `xorm:"comment('上次查询时间') DATETIME"`
	Cid   int       `xorm:"comment('渠道商id') INT(11)"`
	Pid   int       `xorm:"comment('产品id') INT(11)"`
}
