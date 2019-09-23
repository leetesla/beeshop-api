package models

type YunbanTeam struct {
	CustomerId int `xorm:"not null pk INT(11)"`
	ParentId   int `xorm:"not null INT(11)"`
	LeftI      int `xorm:"default 0 INT(11)"`
	RightI     int `xorm:"default 1 INT(11)"`
	Level      int `xorm:"default 1 INT(11)"`
	TreeId     int `xorm:"default 0 comment('root customer id') INT(11)"`
	TitleId    int `xorm:"default 0 comment('0 普通会员，10 经理， 20 总监') TINYINT(3)"`
	Status     int `xorm:"not null default 1 TINYINT(1)"`
}
