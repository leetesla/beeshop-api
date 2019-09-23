package models

type AuthRule struct {
	Name      string `xorm:"not null pk VARCHAR(64)"`
	Data      []byte `xorm:"BLOB"`
	CreatedAt int    `xorm:"INT(11)"`
	UpdatedAt int    `xorm:"INT(11)"`
}
