package models

type DemandProduct struct {
	DemandProductId int    `xorm:"not null pk autoincr INT(11)"`
	DemandId        int    `xorm:"not null INT(11)"`
	ProductId       int    `xorm:"not null default 0 INT(11)"`
	Name            string `xorm:"not null default '' VARCHAR(200)"`
	Model           string `xorm:"not null default '' VARCHAR(200)"`
	GoodsCount      int    `xorm:"not null default 0 comment('申请数量') INT(11)"`
	OutCount        int    `xorm:"not null default 0 comment('出库数量') INT(11)"`
	InCount         int    `xorm:"not null default 0 comment('清点数量') INT(11)"`
}
