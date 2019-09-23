package models

type CustomerGroup struct {
	CustomerGroupId   int `xorm:"not null pk autoincr INT(11)"`
	Approval          int `xorm:"not null INT(1)"`
	CompanyIdDisplay  int `xorm:"not null INT(1)"`
	CompanyIdRequired int `xorm:"not null INT(1)"`
	TaxIdDisplay      int `xorm:"not null INT(1)"`
	TaxIdRequired     int `xorm:"not null INT(1)"`
	SortOrder         int `xorm:"not null INT(3)"`
}
