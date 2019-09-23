package models

type DepositType struct {
	DepositTypeId int    `xorm:"not null pk autoincr INT(11)"`
	PlanId        int    `xorm:"not null INT(11)"`
	Description   string `xorm:"not null default '' VARCHAR(255)"`
	Amount        string `xorm:"not null default 0.00 DECIMAL(15,2)"`
	Bonus         string `xorm:"not null default 0.00 DECIMAL(15,2)"`
}
