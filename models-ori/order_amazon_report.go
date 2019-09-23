package models

type OrderAmazonReport struct {
	OrderId      int    `xorm:"not null index INT(11)"`
	SubmissionId string `xorm:"not null pk VARCHAR(255)"`
	Status       string `xorm:"not null ENUM('error','processing','success')"`
	Text         string `xorm:"not null TEXT"`
}
