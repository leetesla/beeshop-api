package models

type OrderVoucher struct {
	OrderVoucherId int    `xorm:"not null pk autoincr INT(11)"`
	OrderId        int    `xorm:"not null INT(11)"`
	VoucherId      int    `xorm:"not null INT(11)"`
	Description    string `xorm:"not null VARCHAR(255)"`
	Code           string `xorm:"not null VARCHAR(10)"`
	FromName       string `xorm:"not null VARCHAR(64)"`
	FromEmail      string `xorm:"not null VARCHAR(96)"`
	ToName         string `xorm:"not null VARCHAR(64)"`
	ToEmail        string `xorm:"not null VARCHAR(96)"`
	VoucherThemeId int    `xorm:"not null INT(11)"`
	Message        string `xorm:"not null TEXT"`
	Amount         string `xorm:"not null DECIMAL(15,4)"`
}
