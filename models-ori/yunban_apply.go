package models

type YunbanApply struct {
	Id            int    `xorm:"not null pk autoincr INT(11)"`
	CustomerId    int    `xorm:"not null index INT(11)"`
	Realname      string `xorm:"not null default '' VARCHAR(32)"`
	MobilePhone   string `xorm:"default '' VARCHAR(32)"`
	IdentityNo    string `xorm:"not null default '' comment('认证编号，如身份证号') VARCHAR(18)"`
	Intro         string `xorm:"default '' comment('自我介绍：限定字数256') VARCHAR(512)"`
	CountryId     int    `xorm:"default 44 INT(11)"`
	Province      string `xorm:"not null default '' VARCHAR(32)"`
	City          string `xorm:"not null default '' VARCHAR(128)"`
	District      string `xorm:"not null default '' VARCHAR(32)"`
	AddressDetail string `xorm:"not null default '' VARCHAR(128)"`
	AuditDetail   string `xorm:"default '' comment('处理详情') VARCHAR(128)"`
	Target        int    `xorm:"default 20 comment('申请目标对应title_id,默认20，为总监') TINYINT(3)"`
	AuditorId     int    `xorm:"default 0 comment('审核人id') INT(11)"`
	Status        int    `xorm:"default 0 comment('0 待处理，10 已通过， -10 已驳回') TINYINT(3)"`
	CreatedAt     int    `xorm:"comment('创建时间') INT(11)"`
	AuditAt       int    `xorm:"comment('审核时间') INT(11)"`
	IsExceed      int    `xorm:"default 0 comment('审核时是否越级') TINYINT(3)"`
}
