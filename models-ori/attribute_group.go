package models

type AttributeGroup struct {
	AttributeGroupId int `xorm:"not null pk autoincr INT(11)"`
	SortOrder        int `xorm:"not null INT(3)"`
}
