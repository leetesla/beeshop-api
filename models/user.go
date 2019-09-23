package models

import (
	"time"
)

type User struct {
	UserId       int       `json:"user_id" xorm:"not null pk autoincr INT(11)"`
	UserGroupId  int       `xorm:"not null INT(11)"`
	Username     string    `xorm:"not null VARCHAR(20)"`
	PasswordHash string    `xorm:"not null default '' VARCHAR(80)"`
	Salt         string    `xorm:"not null VARCHAR(9)"`
	Firstname    string    `xorm:"not null VARCHAR(32)"`
	Lastname     string    `xorm:"not null VARCHAR(32)"`
	Email        string    `xorm:"not null VARCHAR(96)"`
	Code         string    `xorm:"not null VARCHAR(40)"`
	Ip           string    `xorm:"not null VARCHAR(40)"`
	Status       int       `xorm:"not null TINYINT(1)"`
	ProviderId   int       `xorm:"not null default 0 INT(11)"`
	DateAdded    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	WarehouseId  int       `xorm:"not null default 0 comment('手持设备权限') INT(11)"`
	AuthKey      string    `xorm:"VARCHAR(125)"`
}


func (u User) TableName() string {
	return "user"
	//if u.Role == "admin" {
	//	return "admin_users"
	//} else {
	//	return "users"
	//}
	// return "profiles"
}

// Disable table name's pluralization, if set to true, `User`'s table name will be `user`
// db.SingularTable(true)
