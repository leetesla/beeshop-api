package models

type OrderDistribution struct {
	OrderNo string `xorm:"not null pk default '' VARCHAR(32)"`
}
