package models

type Migration struct {
	Version   string `xorm:"not null pk VARCHAR(180)"`
	ApplyTime int    `xorm:"INT(11)"`
}
