package models

type Attribute struct {
	AttributeId      int `xorm:"not null pk autoincr INT(11)"`
	AttributeGroupId int `xorm:"not null INT(11)"`
	SortOrder        int `xorm:"not null INT(3)"`
}
