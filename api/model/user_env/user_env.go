package user_env

import "configurator/api/component/db"

type UserEnv struct {
	db.Model
	UserId int64 `json:"user_id" gorm:"column:user_id"`
	EnvId  int64 `json:"env_id" gorm:"env_id"`
}
