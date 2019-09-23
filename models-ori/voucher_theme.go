package models

type VoucherTheme struct {
	VoucherThemeId int    `xorm:"not null pk autoincr INT(11)"`
	Image          string `xorm:"not null VARCHAR(255)"`
}
