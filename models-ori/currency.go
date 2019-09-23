package models

import (
	"time"
)

type Currency struct {
	CurrencyId   int       `xorm:"not null pk autoincr INT(11)"`
	Title        string    `xorm:"not null VARCHAR(32)"`
	Code         string    `xorm:"not null VARCHAR(3)"`
	SymbolLeft   string    `xorm:"not null VARCHAR(12)"`
	SymbolRight  string    `xorm:"not null VARCHAR(12)"`
	DecimalPlace string    `xorm:"not null CHAR(1)"`
	Value        float32   `xorm:"not null FLOAT(15,8)"`
	Status       int       `xorm:"not null TINYINT(1)"`
	DateModified time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
