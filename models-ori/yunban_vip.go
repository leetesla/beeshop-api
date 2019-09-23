package models

type YunbanVip struct {
	CustomerId int `xorm:"not null pk INT(11)"`
	CreatedAt  int `xorm:"comment('邀请云伴成为会员时间') INT(11)"`
	ExpiredAt  int `xorm:"comment('该会员过期时间') INT(11)"`
}
