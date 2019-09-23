package models

import (
	"time"
)

type OrderRecurring struct {
	OrderRecurringId   int       `xorm:"not null pk autoincr INT(11)"`
	OrderId            int       `xorm:"not null INT(11)"`
	Created            time.Time `xorm:"not null DATETIME"`
	Status             int       `xorm:"not null TINYINT(4)"`
	ProductId          int       `xorm:"not null INT(11)"`
	ProductName        string    `xorm:"not null VARCHAR(255)"`
	ProductQuantity    int       `xorm:"not null INT(11)"`
	ProfileId          int       `xorm:"not null INT(11)"`
	ProfileName        string    `xorm:"not null VARCHAR(255)"`
	ProfileDescription string    `xorm:"not null VARCHAR(255)"`
	RecurringFrequency string    `xorm:"not null VARCHAR(25)"`
	RecurringCycle     int       `xorm:"not null SMALLINT(6)"`
	RecurringDuration  int       `xorm:"not null SMALLINT(6)"`
	RecurringPrice     string    `xorm:"not null DECIMAL(10,4)"`
	Trial              int       `xorm:"not null TINYINT(1)"`
	TrialFrequency     string    `xorm:"not null VARCHAR(25)"`
	TrialCycle         int       `xorm:"not null SMALLINT(6)"`
	TrialDuration      int       `xorm:"not null SMALLINT(6)"`
	TrialPrice         string    `xorm:"not null DECIMAL(10,4)"`
	ProfileReference   string    `xorm:"not null VARCHAR(255)"`
}
