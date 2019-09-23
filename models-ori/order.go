package models

import (
	"time"
)

type Order struct {
	OrderId               int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId            int       `xorm:"not null default 0 index INT(11)"`
	OrderNo               string    `xorm:"index VARCHAR(32)"`
	OrderTradeNo          string    `xorm:"comment('支付宝交易订单号') VARCHAR(32)"`
	UserSource            string    `xorm:"VARCHAR(32)"`
	UserSourceRaw         string    `xorm:"VARCHAR(255)"`
	OrderSource           string    `xorm:"VARCHAR(32)"`
	OrderSourceRaw        string    `xorm:"VARCHAR(255)"`
	InvoiceNo             int       `xorm:"not null default 0 INT(11)"`
	InvoicePrefix         string    `xorm:"not null VARCHAR(26)"`
	StoreId               int       `xorm:"default 0 INT(11)"`
	StoreName             string    `xorm:"VARCHAR(64)"`
	StoreUrl              string    `xorm:"VARCHAR(255)"`
	CustomerGroupId       int       `xorm:"not null default 0 INT(11)"`
	Firstname             string    `xorm:"VARCHAR(32)"`
	Lastname              string    `xorm:"VARCHAR(32)"`
	Email                 string    `xorm:"not null VARCHAR(96)"`
	Telephone             string    `xorm:"VARCHAR(32)"`
	Fax                   string    `xorm:"VARCHAR(32)"`
	PaymentMethod         string    `xorm:"VARCHAR(128)"`
	PaymentCode           string    `xorm:"VARCHAR(128)"`
	PaymentType           int       `xorm:"not null default 00000000000 INT(11)"`
	PaymentAccount        string    `xorm:"not null default '' VARCHAR(45)"`
	ShippingAddress1      string    `xorm:"not null VARCHAR(128)"`
	ShippingAddress2      string    `xorm:"VARCHAR(128)"`
	ShippingCity          string    `xorm:"not null VARCHAR(128)"`
	ShippingPostcode      string    `xorm:"VARCHAR(10)"`
	ShippingCountry       string    `xorm:"not null VARCHAR(128)"`
	ShippingCountryId     int       `xorm:"not null INT(11)"`
	ShippingZone          string    `xorm:"not null VARCHAR(128)"`
	ShippingZoneId        int       `xorm:"not null INT(11)"`
	ShippingAddressFormat string    `xorm:"not null TEXT"`
	ShippingMethod        string    `xorm:"not null VARCHAR(128)"`
	ShippingCode          string    `xorm:"not null VARCHAR(128)"`
	Comment               string    `xorm:"TEXT"`
	Total                 string    `xorm:"not null default 0.00 DECIMAL(15,2)"`
	OrderStatusId         int       `xorm:"not null default 0 index INT(11)"`
	AffiliateId           int       `xorm:"INT(11)"`
	Commission            string    `xorm:"DECIMAL(15,4)"`
	LanguageId            int       `xorm:"INT(11)"`
	CurrencyId            int       `xorm:"INT(11)"`
	CurrencyCode          string    `xorm:"VARCHAR(3)"`
	CurrencyValue         string    `xorm:"default 1.00000000 DECIMAL(15,8)"`
	Ip                    string    `xorm:"not null VARCHAR(40)"`
	ForwardedIp           string    `xorm:"not null VARCHAR(40)"`
	UserAgent             string    `xorm:"not null VARCHAR(255)"`
	AcceptLanguage        string    `xorm:"VARCHAR(255)"`
	DateAdded             time.Time `xorm:"not null index DATETIME"`
	DateModified          time.Time `xorm:"not null DATETIME"`
	ShippingFullname      string    `xorm:"not null VARCHAR(32)"`
	ShippingProvince      string    `xorm:"not null VARCHAR(32)"`
	ShippingDistrict      string    `xorm:"not null VARCHAR(32)"`
	ShippingMobilePhone   string    `xorm:"not null VARCHAR(32)"`
	ShippingLandlinePhone string    `xorm:"not null VARCHAR(32)"`
	ShippingRealname      string    `xorm:"VARCHAR(32)"`
	ShippingIdentityId    string    `xorm:"VARCHAR(18)"`
	OrderTrackNumber      string    `xorm:"index VARCHAR(32)"`
	OrderTrackCompany     string    `xorm:"VARCHAR(32)"`
	OrderTrackStatus      string    `xorm:"VARCHAR(3000)"`
	StoreHouseId          int       `xorm:"not null default 1 TINYINT(1)"`
	StoreHouseName        string    `xorm:"not null default '' VARCHAR(45)"`
	OfflineOrderId        string    `xorm:"not null default '' VARCHAR(15)"`
	PreSale               int       `xorm:"not null default 0 comment('预售状态: 0 非预售 ; 1 预售') TINYINT(1)"`
	Version               string    `xorm:"not null default '' comment('下单app版本号') VARCHAR(25)"`
	ShippedTime           time.Time `xorm:"comment('变更为可提货状态的时间') DATETIME"`
	OrderType             int       `xorm:"default 0 comment('0 自提订单；
1 直邮订单') INT(11)"`
	CompletedTime time.Time `xorm:"comment('订单状态变成已完成状态的时间') DATETIME"`
	HasReward     int       `xorm:"not null comment('订单是否已经提取过云伴收益') TINYINT(1)"`
	RewardId      int       `xorm:"comment('订单reward之后的对应id') INT(11)"`
}
