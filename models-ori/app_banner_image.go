package models

type AppBannerImage struct {
	Id          int    `xorm:"not null pk autoincr INT(11)"`
	AppBannerId int    `xorm:"not null comment('app_banner表主键id') INT(11)"`
	Image       string `xorm:"default '' comment('banner图地址') VARCHAR(255)"`
	Type        int    `xorm:"default 0 comment('banner跳转类型: 0 HTML链接;1 单品页;2 卖场页;3 搜索页;') TINYINT(2)"`
	Param       string `xorm:"default '' comment('跳转参数') VARCHAR(255)"`
}
