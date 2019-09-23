package models

type Setting struct {
	SettingId  int    `xorm:"not null pk autoincr INT(11)"`
	StoreId    int    `xorm:"not null default 0 INT(11)"`
	Group      string `xorm:"not null VARCHAR(32)"`
	Key        string `xorm:"not null VARCHAR(64)"`
	Value      string `xorm:"not null TEXT"`
	Serialized int    `xorm:"not null TINYINT(1)"`
}
