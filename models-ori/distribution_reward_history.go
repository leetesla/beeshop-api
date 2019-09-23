package models

type DistributionRewardHistory struct {
	Id               int    `xorm:"not null pk autoincr INT(11)"`
	CustomerId       int    `xorm:"not null INT(11)"`
	ProductId        int    `xorm:"not null INT(11)"`
	MyRewardAmount   string `xorm:"not null default 0.000 comment('个人返利金额') DECIMAL(11,3)"`
	TeamRewardAmount string `xorm:"not null default 0.000 comment('团队返利金额') DECIMAL(11,3)"`
	MySales          int    `xorm:"not null default 0 comment('个人本月销量') INT(11)"`
	TeamSales        int    `xorm:"not null default 0 comment('团队商品销售数量') INT(11)"`
	TeamOrderNum     int    `xorm:"not null default 0 comment('团队订单总数') INT(11)"`
	SalePrice        string `xorm:"not null default 0.000 comment('进货价') DECIMAL(11,3)"`
	RewardPrice      string `xorm:"not null default 0.000 comment('返利后价格') DECIMAL(11,3)"`
	RewardLevel      string `xorm:"not null comment('分销级别') VARCHAR(255)"`
	RewardAt         int    `xorm:"comment('返利的具体时间 timestamp') INT(11)"`
}
