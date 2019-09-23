package models

type Language struct {
	LanguageId int    `xorm:"not null pk autoincr INT(11)"`
	Name       string `xorm:"not null index VARCHAR(32)"`
	Code       string `xorm:"not null VARCHAR(5)"`
	Locale     string `xorm:"not null VARCHAR(255)"`
	Image      string `xorm:"not null VARCHAR(64)"`
	Directory  string `xorm:"not null VARCHAR(32)"`
	Filename   string `xorm:"not null VARCHAR(64)"`
	SortOrder  int    `xorm:"not null default 0 INT(3)"`
	Status     int    `xorm:"not null TINYINT(1)"`
}
