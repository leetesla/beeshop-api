package models

import (
	"time"
)

type AppBanner struct {
	Id      int    `xorm:"not null pk autoincr INT(11)"`
	Name    string `xorm:"default '' comment('banner名字') VARCHAR(45)"`
	Pattern int    `xorm:"default 1 comment('1 一张图
3 三张图
4 四张图') TINYINT(1)"`
	Status int `xorm:"default 0 comment('0 停用
1 启用') TINYINT(1)"`
	ShowTimeStart time.Time `xorm:"comment('开始展示时间') DATETIME"`
	ShowTimeEnd   time.Time `xorm:"comment('停止展示时间') DATETIME"`
	BannerOrder   int       `xorm:"default 0 comment('排序') INT(11)"`
	Type          int       `xorm:"default 0 comment('banner类型') INT(11)"`
}
