package operation_log

import (
	"configurator/api/component/db"
)

type OperationLog struct {
	db.Model
	Username string `json:"username" gorm:"column:username"`
	Message  string `json:"message" gorm:"column:message"`
}
