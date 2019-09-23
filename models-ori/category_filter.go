package models

type CategoryFilter struct {
	CategoryId int `xorm:"not null pk INT(11)"`
	FilterId   int `xorm:"not null pk INT(11)"`
}
