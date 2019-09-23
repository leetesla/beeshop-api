package models

type Consignee struct {
	ConsigneeId int    `xorm:"not null pk autoincr INT(11)"`
	CustomerId  int    `xorm:"not null index INT(11)"`
	Realname    string `xorm:"not null default '' comment('提货人全名') VARCHAR(32)"`
	CountryCode string `xorm:"not null default '0086' VARCHAR(6)"`
	MobilePhone string `xorm:"not null default '' index VARCHAR(16)"`
	Tag         string `xorm:"not null default '' comment('标签') VARCHAR(64)"`
	Type        int    `xorm:"not null default 0 comment('0 普通；1 默认') TINYINT(3)"`
}
