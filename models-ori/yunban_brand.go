package models

type YunbanBrand struct {
	Id        int    `xorm:"not null pk autoincr INT(11)"`
	Title     string `xorm:"default '' comment('品牌名字') VARCHAR(128)"`
	Url       string `xorm:"default '' comment('url') VARCHAR(128)"`
	Banner    string `xorm:"default '' comment('banner url') VARCHAR(128)"`
	Status    int    `xorm:"not null default 0 comment('停用 还是 启用') TINYINT(1)"`
	SortOrder int    `xorm:"not null default 0 comment('排序') INT(11)"`
}
