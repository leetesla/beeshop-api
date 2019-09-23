package models

type AuthItem struct {
	Name        string `xorm:"not null pk VARCHAR(64)"`
	Type        int    `xorm:"not null index SMALLINT(6)"`
	Description string `xorm:"TEXT"`
	RuleName    string `xorm:"index VARCHAR(64)"`
	Data        []byte `xorm:"BLOB"`
	CreatedAt   int    `xorm:"INT(11)"`
	UpdatedAt   int    `xorm:"INT(11)"`
}
