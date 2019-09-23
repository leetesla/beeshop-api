package models

import (
	"time"
)

type ProductSpecial struct {
	ProductSpecialId int       `xorm:"not null pk autoincr INT(11)"`
	ProductId        int       `xorm:"not null index INT(11)"`
	CustomerGroupId  int       `xorm:"not null INT(11)"`
	Priority         int       `xorm:"not null default 1 INT(5)"`
	Price            string    `xorm:"not null default 0.0000 DECIMAL(15,4)"`
	DateStart        time.Time `xorm:"DATETIME"`
	DateEnd          time.Time `xorm:"DATETIME"`
}
