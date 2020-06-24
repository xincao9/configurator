package session

import (
    "configurator/api/component/db"
	"time"
)

type Session struct {
	db.Model
	Username  string    `json:"username" gorm:"column:username" binding:"required"`
	Password string    `json:"password" gorm:"column:password" binding:"required"`
	Expire   time.Time `json:"expire" gorm:"column:expire"`
	Token    string    `json:"token" gorm:"column:token"`
}