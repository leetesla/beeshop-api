package models

type ProductDescription struct {
	ProductId        int    `xorm:"not null pk INT(11)"`
	LanguageId       int    `xorm:"not null pk INT(11)"`
	Name             string `xorm:"not null index VARCHAR(255)"`
	Description      string `xorm:"not null TEXT"`
	MetaDescription  string `xorm:"not null VARCHAR(255)"`
	MetaKeyword      string `xorm:"not null VARCHAR(255)"`
	Tag              string `xorm:"not null TEXT"`
	ShortDescription string `xorm:"VARCHAR(1024)"`
	ShortName        string `xorm:"VARCHAR(8)"`
	SearchKeyword    string `xorm:"not null comment('搜索关键词') VARCHAR(255)"`
}
