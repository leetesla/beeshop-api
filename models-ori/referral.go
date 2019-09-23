package models

type Referral struct {
	CustomerId int `xorm:"not null pk INT(11)"`
	ParentId   int `xorm:"not null default 0 INT(11)"`
}
