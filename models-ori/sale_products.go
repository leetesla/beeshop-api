package models

type SaleProducts struct {
	SaleId    int `xorm:"not null pk index INT(11)"`
	ProductId int `xorm:"not null pk INT(11)"`
	Priority  int `xorm:"INT(4)"`
	Viewed    int `xorm:"not null default 0 INT(11)"`
	Purchased int `xorm:"not null default 0 INT(11)"`
	Stock     int `xorm:"not null TINYINT(2)"`
}
