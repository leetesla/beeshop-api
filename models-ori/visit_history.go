package models

import (
	"time"
)

type VisitHistory struct {
	VisitId int       `xorm:"not null pk autoincr INT(11)"`
	OrderId int       `xorm:"not null INT(11)"`
	UserId  int       `xorm:"not null INT(11)"`
	Content string    `xorm:"not null TEXT"`
	Msg     string    `xorm:"not null TEXT"`
	IsMsg   int       `xorm:"not null default 0 comment('0 未发送；1 已发送') TINYINT(1)"`
	Addtime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('回访时间') DATETIME"`
}
