package models

import (
	"time"
)

type Customer struct {
	CustomerId      int       `xorm:"not null pk autoincr INT(11)"`
	StoreId         int       `xorm:"not null default 0 INT(11)"`
	Firstname       string    `xorm:"not null VARCHAR(32)"`
	Lastname        string    `xorm:"not null VARCHAR(32)"`
	Username        string    `xorm:"not null comment('用户名') VARCHAR(32)"`
	Email           string    `xorm:"not null index VARCHAR(96)"`
	Telephone       string    `xorm:"not null index VARCHAR(32)"`
	IsMobile        int       `xorm:"not null default 0 comment('是否是手机号注册登录用户') TINYINT(1)"`
	Fax             string    `xorm:"not null VARCHAR(32)"`
	Password        string    `xorm:"not null VARCHAR(40)"`
	Salt            string    `xorm:"not null VARCHAR(9)"`
	Cart            string    `xorm:"TEXT"`
	Wishlist        string    `xorm:"TEXT"`
	Newsletter      int       `xorm:"not null default 0 TINYINT(1)"`
	AddressId       int       `xorm:"not null default 0 INT(11)"`
	CustomerGroupId int       `xorm:"not null INT(11)"`
	Ip              string    `xorm:"not null default '0' VARCHAR(40)"`
	Status          int       `xorm:"not null TINYINT(1)"`
	Approved        int       `xorm:"not null TINYINT(1)"`
	Token           string    `xorm:"not null VARCHAR(255)"`
	DateAdded       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	HeadImage       string    `xorm:"default '' comment('头像地址') VARCHAR(255)"`
	HeadUrl         string    `xorm:"default '' comment('第三方登陆头像地址') VARCHAR(255)"`
	LoginType       int       `xorm:"default 0000 comment('登录类型
0：普通登录
1：微博登录
2：微信登录
3：qq登录') SMALLINT(4)"`
	UniqueId                 string    `xorm:"comment('第三方登陆唯一id') VARCHAR(255)"`
	VerCode                  string    `xorm:"default '' VARCHAR(32)"`
	VercodeTime              time.Time `xorm:"DATETIME"`
	BabySex                  int       `xorm:"not null default 0 comment('宝贝性别:女，男，孕育中') SMALLINT(4)"`
	BabyBirthday             time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('出生日期或者预产期') DATETIME"`
	StoreHouseId             int       `xorm:"not null default 1 INT(11)"`
	YunbanRewardEver         string    `xorm:"default 0.00 comment('收益：过往，暂时指本月之前') DECIMAL(15,2)"`
	YunbanRewardWithdrawable string    `xorm:"default 0.00 comment('收益：可提现') DECIMAL(15,2)"`
	Balance                  string    `xorm:"default 0.00 DECIMAL(15,2)"`
	PaymentPassword          string    `xorm:"default '' VARCHAR(40)"`
	PaymentPasswordSalt      string    `xorm:"default '' VARCHAR(9)"`
	CountryCode              string    `xorm:"VARCHAR(8)"`
	PaymentVerCode           string    `xorm:"VARCHAR(32)"`
	PaymentVercodeTime       time.Time `xorm:"DATETIME"`
	OpenId                   string    `xorm:"default '' VARCHAR(45)"`
	UnionId                  string    `xorm:"default '' VARCHAR(45)"`
	Weixin                   string    `xorm:"default '' comment('微信号') VARCHAR(45)"`
	DirectMailAddress        int       `xorm:"not null default 0 INT(11)"`
	IsAgent                  int       `xorm:"not null default 00 comment('是否是云伴代理商') TINYINT(2)"`
	AgentTime                time.Time `xorm:"comment('加入云伴代理商的时间') DATETIME"`
	YunbanTime               time.Time `xorm:"comment('加入云伴的时间') DATETIME"`
	MyInviteCode             string    `xorm:"default '' comment('我的邀请码') CHAR(6)"`
	InviterId                int       `xorm:"default 0 comment('邀请人customer id') INT(11)"`
	Sex                      int       `xorm:"not null default 0 comment('性别:0未知，1女，2男') SMALLINT(4)"`
	Birthday                 time.Time `xorm:"not null comment('生日') DATETIME"`
	AccountNo                string    `xorm:"not null default '' comment('账户名') VARCHAR(32)"`
}
