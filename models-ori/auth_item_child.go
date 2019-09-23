package models

type AuthItemChild struct {
	Parent string `xorm:"not null pk VARCHAR(64)"`
	Child  string `xorm:"not null pk index VARCHAR(64)"`
}
