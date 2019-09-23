package models

type Warehouse struct {
	Id            int    `xorm:"not null pk autoincr INT(11)"`
	WarehouseId   int    `xorm:"not null INT(11)"`
	WarehouseName string `xorm:"default '' VARCHAR(45)"`
}
