package models

type Store struct {
	StoreId int    `xorm:"not null pk autoincr INT(11)"`
	Name    string `xorm:"not null VARCHAR(64)"`
	Url     string `xorm:"not null VARCHAR(255)"`
	Ssl     string `xorm:"not null VARCHAR(255)"`
}
