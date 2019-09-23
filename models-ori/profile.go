package models

type Profile struct {
	ProfileId      int    `xorm:"not null pk autoincr INT(11)"`
	SortOrder      int    `xorm:"not null INT(11)"`
	Status         int    `xorm:"not null TINYINT(4)"`
	Price          string `xorm:"not null DECIMAL(10,4)"`
	Frequency      string `xorm:"not null ENUM('day','month','semi_month','week','year')"`
	Duration       int    `xorm:"not null INT(10)"`
	Cycle          int    `xorm:"not null INT(10)"`
	TrialStatus    int    `xorm:"not null TINYINT(4)"`
	TrialPrice     string `xorm:"not null DECIMAL(10,4)"`
	TrialFrequency string `xorm:"not null ENUM('day','month','semi_month','week','year')"`
	TrialDuration  int    `xorm:"not null INT(10)"`
	TrialCycle     int    `xorm:"not null INT(10)"`
}
