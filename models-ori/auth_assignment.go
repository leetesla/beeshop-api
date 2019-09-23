package models

type AuthAssignment struct {
	ItemName  string `xorm:"not null pk VARCHAR(64)"`
	UserId    string `xorm:"not null pk index VARCHAR(64)"`
	CreatedAt int    `xorm:"INT(11)"`
}
