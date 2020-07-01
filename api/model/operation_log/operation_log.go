package operation_log

import (
    "configurator/api/component/db"
)

type OperationLog struct {
    db.Model
    Username string `json:"share_id" gorm:"column:share_id"`
    Message  string `json:"message" gorm:"column:message"`
}
