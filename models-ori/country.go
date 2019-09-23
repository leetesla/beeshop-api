package models

type Country struct {
	CountryId        int    `xorm:"not null pk autoincr INT(11)"`
	Name             string `xorm:"not null VARCHAR(128)"`
	IsoCode2         string `xorm:"not null VARCHAR(2)"`
	IsoCode3         string `xorm:"not null VARCHAR(3)"`
	AddressFormat    string `xorm:"not null TEXT"`
	PostcodeRequired int    `xorm:"not null TINYINT(1)"`
	Status           int    `xorm:"not null default 1 TINYINT(1)"`
}
