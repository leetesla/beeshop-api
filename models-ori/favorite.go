package models

import (
	"time"
)

type Favorite struct {
	FavoriteId int       `xorm:"not null pk autoincr INT(11)"`
	CustomerId int       `xorm:"not null index INT(11)"`
	ProductId  int       `xorm:"not null index INT(11)"`
	CreatedAt  time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
