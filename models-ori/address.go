package models

type Address struct {
	AddressId      int    `xorm:"not null pk autoincr INT(11)"`
	CustomerId     int    `xorm:"not null index INT(11)"`
	Firstname      string `xorm:"VARCHAR(32)"`
	Lastname       string `xorm:"VARCHAR(32)"`
	Company        string `xorm:"VARCHAR(32)"`
	CompanyId      string `xorm:"VARCHAR(32)"`
	TaxId          string `xorm:"VARCHAR(32)"`
	Address1       string `xorm:"not null VARCHAR(128)"`
	Address2       string `xorm:"not null VARCHAR(128)"`
	City           string `xorm:"not null VARCHAR(128)"`
	Postcode       string `xorm:"VARCHAR(10)"`
	CountryId      int    `xorm:"default 0 INT(11)"`
	ZoneId         int    `xorm:"default 0 INT(11)"`
	Province       string `xorm:"not null VARCHAR(32)"`
	District       string `xorm:"not null VARCHAR(32)"`
	Fullname       string `xorm:"not null VARCHAR(32)"`
	MobilePhone    string `xorm:"VARCHAR(32)"`
	LandlinePhone  string `xorm:"VARCHAR(32)"`
	Realname       string `xorm:"not null default '' VARCHAR(32)"`
	IdentityId     string `xorm:"not null default '' VARCHAR(18)"`
	ShippingOption int    `xorm:"not null default 0 TINYINT(1)"`
	AddressType    int    `xorm:"not null default 0 comment('0 自提；1 直邮') INT(11)"`
	Default        int    `xorm:"not null default 0 comment('状态:0 普通，1 默认') TINYINT(3)"`
	CountryCode    string `xorm:"not null default '0086' comment('地区码') VARCHAR(8)"`
	Tag            string `xorm:"not null default '' comment('标签名') VARCHAR(32)"`
}
