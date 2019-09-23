package models

import (
	"time"
)

type AffiliateTransaction struct {
	AffiliateTransactionId int       `xorm:"not null pk autoincr INT(11)"`
	AffiliateId            int       `xorm:"not null INT(11)"`
	OrderId                int       `xorm:"not null INT(11)"`
	Description            string    `xorm:"not null TEXT"`
	Amount                 string    `xorm:"not null DECIMAL(15,4)"`
	DateAdded              time.Time `xorm:"not null DATETIME"`
}
