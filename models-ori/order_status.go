package models

type OrderStatus struct {
	OrderStatusId int    `xorm:"not null pk autoincr INT(11)"`
	LanguageId    int    `xorm:"not null pk INT(11)"`
	Name          string `xorm:"not null VARCHAR(32)"`
}
