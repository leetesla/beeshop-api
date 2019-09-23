package models

type StoreHouse struct {
	StoreHouseId int    `xorm:"not null pk autoincr INT(4)"`
	Name         string `xorm:"VARCHAR(50)"`
	Nickname     string `xorm:"VARCHAR(50)"`
	PriceName    string `xorm:"default '' VARCHAR(20)"`
	Status       int    `xorm:"not null default 0 comment('0 停用;1 启用') TINYINT(1)"`
	Sort         int    `xorm:"default 0 INT(11)"`
}
