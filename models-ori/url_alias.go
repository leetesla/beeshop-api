package models

type UrlAlias struct {
	UrlAliasId int    `xorm:"not null pk autoincr INT(11)"`
	Query      string `xorm:"not null VARCHAR(255)"`
	Keyword    string `xorm:"not null VARCHAR(255)"`
}
