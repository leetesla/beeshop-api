package models

type ProductToDistribution struct {
	ProductId int    `xorm:"not null index INT(11)"`
	Sales     int    `xorm:"not null default 0 comment('销量') INT(11)"`
	SalePrice string `xorm:"not null default 0.00 DECIMAL(15,2)"`
	Desc      string `xorm:"comment('描述') VARCHAR(255)"`
	Status    int    `xorm:"default 1 comment('是否启用，1是0否') TINYINT(3)"`
	CreatedAt int    `xorm:"comment('添加时间') INT(11)"`
	UpdatedAt int    `xorm:"comment('修改时间') INT(11)"`
}
