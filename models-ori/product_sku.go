package models

type ProductSku struct {
	ProductId  int    `xorm:"not null INT(11)"`
	SkuCode    int    `xorm:"not null INT(11)"`
	LanguageId int    `xorm:"not null INT(11)"`
	SkuKvs     string `xorm:"comment('JSON') JSON"`
	Quantity   int    `xorm:"INT(11)"`
	Price      string `xorm:"not null default 0.00 DECIMAL(15,2)"`
}
