package models

type BannerImageDescription struct {
	BannerImageId int    `xorm:"not null pk INT(11)"`
	LanguageId    int    `xorm:"not null pk INT(11)"`
	BannerId      int    `xorm:"not null INT(11)"`
	Title         string `xorm:"not null VARCHAR(64)"`
}
