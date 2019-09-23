package models

type Mfgr struct {
	MfgrId    int    `xorm:"not null pk autoincr INT(11)"`
	Name      string `xorm:"not null VARCHAR(64)"`
	Image     string `xorm:"VARCHAR(255)"`
	SortOrder int    `xorm:"not null INT(3)"`
	Hot       int    `xorm:"not null default 0 TINYINT(4)"`
}
