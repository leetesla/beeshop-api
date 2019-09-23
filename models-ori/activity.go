package models

import (
	"time"
)

type Activity struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	ActivityName string    `xorm:"default '' comment('活动名称') VARCHAR(45)"`
	StartTime    time.Time `xorm:"comment('活动开始时间') DATETIME"`
	EndTime      time.Time `xorm:"comment('活动结束时间') DATETIME"`
	Type         int       `xorm:"default 0 comment('活动类型') TINYINT(2)"`
	Extra        string    `xorm:"comment('额外信息,JSON') TEXT"`
}
