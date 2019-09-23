package models

type OpenbayFaq struct {
	Id    int    `xorm:"not null pk autoincr INT(11)"`
	Route string `xorm:"not null TEXT"`
}
