package models

type CustomField struct {
	CustomFieldId int    `xorm:"not null pk autoincr INT(11)"`
	Type          string `xorm:"not null VARCHAR(32)"`
	Value         string `xorm:"not null TEXT"`
	Required      int    `xorm:"not null TINYINT(1)"`
	Location      string `xorm:"not null VARCHAR(32)"`
	Position      int    `xorm:"not null INT(3)"`
	SortOrder     int    `xorm:"not null INT(3)"`
}
