package models

import (
	"time"
)

type AlipayTransaction struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	OutTradeNo  string    `xorm:"not null default '' comment('对应order表的order_no') VARCHAR(32)"`
	TradeNo     string    `xorm:"not null default '' comment('支付宝流水号') VARCHAR(32)"`
	TotalFee    string    `xorm:"default 0.00 comment('交易金额') DECIMAL(15,2)"`
	GmtCreate   string    `xorm:"default '' comment('交易创建时间') VARCHAR(32)"`
	GmtClose    string    `xorm:"default '' comment('交易关闭时间') VARCHAR(32)"`
	SignType    string    `xorm:"default '' comment('md5,rsa') VARCHAR(5)"`
	NotifyId    string    `xorm:"default '' VARCHAR(45)"`
	TradeStatus string    `xorm:"default '' comment('交易结果') VARCHAR(45)"`
	Addtime     time.Time `xorm:"DATETIME"`
	TradeType   string    `xorm:"default '0' comment('交易类型：0 充值； 1 支付；2 退款') VARCHAR(45)"`
	SellerId    int64     `xorm:"default 0 comment('卖家支付宝账号，以2088开头由16位纯数字组成的字符串，一般情况下收款账号就是签约账号') BIGINT(16)"`
}
