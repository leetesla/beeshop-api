package models

type Option struct {
	OptionId  int    `xorm:"not null pk autoincr INT(11)"`
	Type      string `xorm:"not null VARCHAR(32)"`
	SortOrder int    `xorm:"not null INT(3)"`
}
