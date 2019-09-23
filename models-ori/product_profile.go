package models

type ProductProfile struct {
	ProductId       int `xorm:"not null pk INT(11)"`
	ProfileId       int `xorm:"not null pk INT(11)"`
	CustomerGroupId int `xorm:"not null pk INT(11)"`
}
