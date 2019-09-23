package models

type StoreCategorySort struct {
	StoreId    int `xorm:"not null pk INT(11)"`
	CategoryId int `xorm:"not null pk INT(11)"`
	SortOrder  int `xorm:"not null default 0 INT(11)"`
}
