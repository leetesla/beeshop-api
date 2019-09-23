package models

type ProductImage struct {
	ProductImageId int    `xorm:"not null pk autoincr INT(11)"`
	ProductId      int    `xorm:"not null INT(11)"`
	SkuCode        int    `xorm:"not null INT(11)"`
	Image          string `xorm:"VARCHAR(255)"`
	SortOrder      int    `xorm:"not null default 0 INT(3)"`
}
