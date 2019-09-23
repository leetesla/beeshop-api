package models

type OrderSuborder struct {
	OrderId           int    `xorm:"not null pk index INT(11)"`
	StoreHouseId      int    `xorm:"not null pk index TINYINT(4)"`
	OrderTrackNumber  string `xorm:"index VARCHAR(32)"`
	OrderTrackCompany string `xorm:"VARCHAR(32)"`
	OrderTrackStatus  string `xorm:"VARCHAR(3000)"`
	OrderStatusId     int    `xorm:"index TINYINT(4)"`
}
