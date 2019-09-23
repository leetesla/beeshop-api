package models

type YunbanRewardHistory struct {
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	CustomerId      int    `xorm:"INT(11)"`
	RewardPercent   string `xorm:"default 0.00000 comment('团队返利比例') DECIMAL(8,5)"`
	RewardAmount    string `xorm:"default 0.000 comment('个人返利金额') DECIMAL(11,3)"`
	TeamSalesAmount string `xorm:"default 0.000 comment('团队销售额') DECIMAL(11,3)"`
	TeamOrderNum    int    `xorm:"default 0 comment('团队订单总数') INT(11)"`
	RewardMonthAt   int    `xorm:"comment('返利月份：201811 冗余字段 ') INT(11)"`
	RewardAt        int    `xorm:"comment('返利的具体时间 timestamp') INT(11)"`
}
