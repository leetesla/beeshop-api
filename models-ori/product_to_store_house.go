package models

type ProductToStoreHouse struct {
	ProductId        int    `xorm:"not null pk INT(11)"`
	StoreHouseId     int    `xorm:"not null pk INT(11)"`
	Quantity         int    `xorm:"not null default 0 INT(11)"`
	SalePrice        string `xorm:"not null default 0.00 DECIMAL(15,2)"`
	VipPrice         string `xorm:"not null default 0.00 DECIMAL(15,2)"`
	VipCompanyPrice  string `xorm:"not null default 0.00 comment('大客户价格') DECIMAL(15,2)"`
	OfflineSalePrice string `xorm:"not null default 0.00 DECIMAL(15,2)"`
	OfflineVipPrice  string `xorm:"not null default 0.00 DECIMAL(15,2)"`
}
