package models

type ProductSpecification struct {
	Id   int    `xorm:"not null pk autoincr INT(11)"`
	Name string `xorm:"not null comment('规格名字') VARCHAR(32)"`
}
