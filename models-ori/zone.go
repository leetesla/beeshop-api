package models

type Zone struct {
	ZoneId    int    `xorm:"not null pk autoincr INT(11)"`
	CountryId int    `xorm:"not null INT(11)"`
	Name      string `xorm:"not null VARCHAR(128)"`
	Code      string `xorm:"not null VARCHAR(32)"`
	Status    int    `xorm:"not null default 1 TINYINT(1)"`
}
