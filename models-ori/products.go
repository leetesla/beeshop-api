package models

import (
	"time"
)

type Products struct {
	Id        int       `xorm:"not null pk autoincr INT(10)"`
	CreatedAt time.Time `xorm:"TIMESTAMP"`
	UpdatedAt time.Time `xorm:"TIMESTAMP"`
	DeletedAt time.Time `xorm:"index TIMESTAMP"`
	Code      string    `xorm:"VARCHAR(255)"`
	Price     int       `xorm:"INT(10)"`
}
