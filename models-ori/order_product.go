package models

type OrderProduct struct {
	OrderProductId int    `xorm:"not null pk autoincr INT(11)"`
	OrderId        int    `xorm:"not null index INT(11)"`
	ProductId      int    `xorm:"not null index INT(11)"`
	Name           string `xorm:"not null VARCHAR(255)"`
	Model          string `xorm:"not null VARCHAR(64)"`
	Quantity       int    `xorm:"not null INT(4)"`
	Price          string `xorm:"not null default 0.00 DECIMAL(15,2)"`
	Total          string `xorm:"not null default 0.00 DECIMAL(15,2)"`
	Weight         string `xorm:"not null default 0.00000000 DECIMAL(15,8)"`
	Tax            string `xorm:"not null default 0.0000 DECIMAL(15,4)"`
	Reward         int    `xorm:"not null INT(8)"`
	SaleId         int    `xorm:"default -1 INT(11)"`
	StoreHouseId   int    `xorm:"default 1 index TINYINT(4)"`
	Reviewed       int    `xorm:"not null default 0 TINYINT(1)"`
	SalePrice      string `xorm:"not null default 0.00 DECIMAL(15,2)"`
	VipPrice       string `xorm:"not null default 0.00 DECIMAL(15,2)"`
}
