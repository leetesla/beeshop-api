package models

type OptionValue struct {
	OptionValueId int    `xorm:"not null pk autoincr INT(11)"`
	OptionId      int    `xorm:"not null INT(11)"`
	Image         string `xorm:"not null VARCHAR(255)"`
	SortOrder     int    `xorm:"not null INT(3)"`
}
