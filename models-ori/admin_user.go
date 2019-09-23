package models

import (
	"time"
)

type AdminUser struct {
	UserId      int       `xorm:"not null pk autoincr INT(11)"`
	UserGroupId int       `xorm:"not null INT(11)"`
	Username    string    `xorm:"not null VARCHAR(20)"`
	Password    string    `xorm:"not null VARCHAR(40)"`
	Salt        string    `xorm:"not null VARCHAR(9)"`
	Firstname   string    `xorm:"not null VARCHAR(32)"`
	Lastname    string    `xorm:"not null VARCHAR(32)"`
	Email       string    `xorm:"not null VARCHAR(96)"`
	Code        string    `xorm:"not null VARCHAR(40)"`
	Ip          string    `xorm:"not null VARCHAR(40)"`
	Status      int       `xorm:"not null TINYINT(1)"`
	ProviderId  int       `xorm:"not null default 0 INT(11)"`
	DateAdded   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	WarehouseId int       `xorm:"not null default 0 comment('手持设备权限') INT(11)"`
}
