package models

type ProductToDownload struct {
	ProductId  int `xorm:"not null pk INT(11)"`
	DownloadId int `xorm:"not null pk INT(11)"`
}
