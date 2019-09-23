package models

import (
	"time"
)

type OrderFraud struct {
	OrderId                        int       `xorm:"not null pk INT(11)"`
	CustomerId                     int       `xorm:"not null INT(11)"`
	CountryMatch                   string    `xorm:"not null VARCHAR(3)"`
	CountryCode                    string    `xorm:"not null VARCHAR(2)"`
	HighRiskCountry                string    `xorm:"not null VARCHAR(3)"`
	Distance                       int       `xorm:"not null INT(11)"`
	IpRegion                       string    `xorm:"not null VARCHAR(255)"`
	IpCity                         string    `xorm:"not null VARCHAR(255)"`
	IpLatitude                     string    `xorm:"not null DECIMAL(10,6)"`
	IpLongitude                    string    `xorm:"not null DECIMAL(10,6)"`
	IpIsp                          string    `xorm:"not null VARCHAR(255)"`
	IpOrg                          string    `xorm:"not null VARCHAR(255)"`
	IpAsnum                        int       `xorm:"not null INT(11)"`
	IpUserType                     string    `xorm:"not null VARCHAR(255)"`
	IpCountryConfidence            string    `xorm:"not null VARCHAR(3)"`
	IpRegionConfidence             string    `xorm:"not null VARCHAR(3)"`
	IpCityConfidence               string    `xorm:"not null VARCHAR(3)"`
	IpPostalConfidence             string    `xorm:"not null VARCHAR(3)"`
	IpPostalCode                   string    `xorm:"not null VARCHAR(10)"`
	IpAccuracyRadius               int       `xorm:"not null INT(11)"`
	IpNetSpeedCell                 string    `xorm:"not null VARCHAR(255)"`
	IpMetroCode                    int       `xorm:"not null INT(3)"`
	IpAreaCode                     int       `xorm:"not null INT(3)"`
	IpTimeZone                     string    `xorm:"not null VARCHAR(255)"`
	IpRegionName                   string    `xorm:"not null VARCHAR(255)"`
	IpDomain                       string    `xorm:"not null VARCHAR(255)"`
	IpCountryName                  string    `xorm:"not null VARCHAR(255)"`
	IpContinentCode                string    `xorm:"not null VARCHAR(2)"`
	IpCorporateProxy               string    `xorm:"not null VARCHAR(3)"`
	AnonymousProxy                 string    `xorm:"not null VARCHAR(3)"`
	ProxyScore                     int       `xorm:"not null INT(3)"`
	IsTransProxy                   string    `xorm:"not null VARCHAR(3)"`
	FreeMail                       string    `xorm:"not null VARCHAR(3)"`
	CarderEmail                    string    `xorm:"not null VARCHAR(3)"`
	HighRiskUsername               string    `xorm:"not null VARCHAR(3)"`
	HighRiskPassword               string    `xorm:"not null VARCHAR(3)"`
	BinMatch                       string    `xorm:"not null VARCHAR(10)"`
	BinCountry                     string    `xorm:"not null VARCHAR(2)"`
	BinNameMatch                   string    `xorm:"not null VARCHAR(3)"`
	BinName                        string    `xorm:"not null VARCHAR(255)"`
	BinPhoneMatch                  string    `xorm:"not null VARCHAR(3)"`
	BinPhone                       string    `xorm:"not null VARCHAR(32)"`
	CustomerPhoneInBillingLocation string    `xorm:"not null VARCHAR(8)"`
	ShipForward                    string    `xorm:"not null VARCHAR(3)"`
	CityPostalMatch                string    `xorm:"not null VARCHAR(3)"`
	ShipCityPostalMatch            string    `xorm:"not null VARCHAR(3)"`
	Score                          string    `xorm:"not null DECIMAL(10,5)"`
	Explanation                    string    `xorm:"not null TEXT"`
	RiskScore                      string    `xorm:"not null DECIMAL(10,5)"`
	QueriesRemaining               int       `xorm:"not null INT(11)"`
	MaxmindId                      string    `xorm:"not null VARCHAR(8)"`
	Error                          string    `xorm:"not null TEXT"`
	DateAdded                      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
