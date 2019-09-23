package models

type Menu struct {
	Id     int    `xorm:"not null pk autoincr INT(11)"`
	Name   string `xorm:"not null VARCHAR(128)"`
	Parent int    `xorm:"index INT(11)"`
	Route  string `xorm:"VARCHAR(255)"`
	Order  int    `xorm:"INT(11)"`
	Data   []byte `xorm:"BLOB"`
}
