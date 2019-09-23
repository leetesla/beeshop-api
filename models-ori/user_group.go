package models

type UserGroup struct {
	UserGroupId int    `xorm:"not null pk autoincr INT(11)"`
	Name        string `xorm:"not null VARCHAR(64)"`
	Permission  string `xorm:"not null TEXT"`
}
