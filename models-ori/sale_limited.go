package models

import (
	"time"
)

type SaleLimited struct {
	SaleId         int       `xorm:"not null pk autoincr INT(11)"`
	ManufacturerId int       `xorm:"not null INT(11)"`
	DateStart      time.Time `xorm:"not null DATETIME"`
	DateEnd        time.Time `xorm:"not null DATETIME"`
	Status         int       `xorm:"not null default 1 index TINYINT(1)"`
	Name           string    `xorm:"not null VARCHAR(255)"`
	Description    string    `xorm:"CHAR(255)"`
	Discount       string    `xorm:"CHAR(255)"`
	BannerSm       string    `xorm:"VARCHAR(255)"`
	BannerLg       string    `xorm:"VARCHAR(255)"`
	Logo           string    `xorm:"default '' VARCHAR(100)"`
	BgColor        string    `xorm:"default '#f4f4f4' VARCHAR(10)"`
	Visible        int       `xorm:"not null default 1 TINYINT(1)"`
	SortOrder      int       `xorm:"not null default 0 INT(11)"`
	SaleType       int       `xorm:"default 0 TINYINT(4)"`
	Hot            int       `xorm:"not null default 0 index TINYINT(1)"`
	PreSale        int       `xorm:"not null default 0 comment('预售状态:0 普通卖场 ; 1 预售卖场') TINYINT(1)"`
}
