package models

type Channel struct {
	Id      int    `xorm:"not null pk autoincr comment('序列号') INT(11)"`
	Name    string `xorm:"default '' comment('渠道商名字') VARCHAR(32)"`
	Contact string `xorm:"default '' comment('联系人') VARCHAR(32)"`
	Phone   string `xorm:"default '' comment('联系电话') VARCHAR(32)"`
}
