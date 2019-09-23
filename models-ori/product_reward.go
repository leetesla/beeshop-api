package models

type ProductReward struct {
	ProductRewardId int `xorm:"not null pk autoincr INT(11)"`
	ProductId       int `xorm:"not null default 0 INT(11)"`
	CustomerGroupId int `xorm:"not null default 0 INT(11)"`
	Points          int `xorm:"not null default 0 INT(8)"`
}
