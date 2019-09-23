package models

import (
	"time"
)

type Affiliate struct {
	AffiliateId       int       `xorm:"not null pk autoincr INT(11)"`
	Firstname         string    `xorm:"not null VARCHAR(32)"`
	Lastname          string    `xorm:"not null VARCHAR(32)"`
	Email             string    `xorm:"not null VARCHAR(96)"`
	Telephone         string    `xorm:"not null VARCHAR(32)"`
	Fax               string    `xorm:"not null VARCHAR(32)"`
	Password          string    `xorm:"not null VARCHAR(40)"`
	Salt              string    `xorm:"not null VARCHAR(9)"`
	Company           string    `xorm:"not null VARCHAR(32)"`
	Website           string    `xorm:"not null VARCHAR(255)"`
	Address1          string    `xorm:"not null VARCHAR(128)"`
	Address2          string    `xorm:"not null VARCHAR(128)"`
	City              string    `xorm:"not null VARCHAR(128)"`
	Postcode          string    `xorm:"not null VARCHAR(10)"`
	CountryId         int       `xorm:"not null INT(11)"`
	ZoneId            int       `xorm:"not null INT(11)"`
	Code              string    `xorm:"not null VARCHAR(64)"`
	Commission        string    `xorm:"not null default 0.00 DECIMAL(4,2)"`
	Tax               string    `xorm:"not null VARCHAR(64)"`
	Payment           string    `xorm:"not null VARCHAR(6)"`
	Cheque            string    `xorm:"not null VARCHAR(100)"`
	Paypal            string    `xorm:"not null VARCHAR(64)"`
	BankName          string    `xorm:"not null VARCHAR(64)"`
	BankBranchNumber  string    `xorm:"not null VARCHAR(64)"`
	BankSwiftCode     string    `xorm:"not null VARCHAR(64)"`
	BankAccountName   string    `xorm:"not null VARCHAR(64)"`
	BankAccountNumber string    `xorm:"not null VARCHAR(64)"`
	Ip                string    `xorm:"not null VARCHAR(40)"`
	Status            int       `xorm:"not null TINYINT(1)"`
	Approved          int       `xorm:"not null TINYINT(1)"`
	DateAdded         time.Time `xorm:"not null DATETIME"`
}
