package models

type Banner struct {
	BannerId int    `xorm:"not null pk autoincr INT(11)"`
	Name     string `xorm:"not null VARCHAR(64)"`
	Status   int    `xorm:"not null TINYINT(1)"`
}
