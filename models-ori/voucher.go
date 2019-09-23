package models

import (
	"time"
)

type Voucher struct {
	VoucherId      int       `xorm:"not null pk autoincr INT(11)"`
	OrderId        int       `xorm:"not null INT(11)"`
	Code           string    `xorm:"not null VARCHAR(10)"`
	FromName       string    `xorm:"not null VARCHAR(64)"`
	FromEmail      string    `xorm:"not null VARCHAR(96)"`
	ToName         string    `xorm:"not null VARCHAR(64)"`
	ToEmail        string    `xorm:"not null VARCHAR(96)"`
	VoucherThemeId int       `xorm:"not null INT(11)"`
	Message        string    `xorm:"not null TEXT"`
	Amount         string    `xorm:"not null DECIMAL(15,4)"`
	Status         int       `xorm:"not null TINYINT(1)"`
	DateAdded      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
