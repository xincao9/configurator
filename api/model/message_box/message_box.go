package message_box

import (
    "configurator/api/component/db"
)

type MessageBox struct {
    db.Model
    Username string `json:"username" gorm:"column:username"`
    Message  string `json:"message" gorm:"column:message"`
    Status   int32  `json:"status" gorm:"column:status"`
}
