package models

type YunbanPromoteLog struct {
	Id          int `xorm:"not null pk autoincr INT(11)"`
	CustomerId  int `xorm:"not null default 0 index INT(11)"`
	ParentId    int `xorm:"not null default 0 index INT(11)"`
	TitleIdFrom int `xorm:"INT(11)"`
	TitleIdTo   int `xorm:"comment('升级后title id') INT(11)"`
	CreatedAt   int `xorm:"comment('升级时间') index INT(11)"`
	IsExceed    int `xorm:"default 0 comment('是否越级，1是 0否') TINYINT(4)"`
	Status      int `xorm:"default 0 comment('0 新增；5 已返利') TINYINT(4)"`
}
