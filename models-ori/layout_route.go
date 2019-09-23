package models

type LayoutRoute struct {
	LayoutRouteId int    `xorm:"not null pk autoincr INT(11)"`
	LayoutId      int    `xorm:"not null INT(11)"`
	StoreId       int    `xorm:"not null INT(11)"`
	Route         string `xorm:"not null VARCHAR(255)"`
}
