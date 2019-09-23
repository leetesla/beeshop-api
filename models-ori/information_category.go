package models

type InformationCategory struct {
	CategoryId   int    `xorm:"not null pk autoincr INT(11)"`
	CategoryName string `xorm:"VARCHAR(45)"`
}
