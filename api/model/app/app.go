package app

import (
	"configurator/api/component/db"
	"fmt"
)

const (
	keyFormat = "configurator|%d|%s|%s|%s"
)

type App struct {
	db.Model
	EnvId   int64  `json:"env_id" gorm:"column:env_id" binding:"required"`
	Group   string `json:"group" gorm:"column:group" binding:"required"`
	Project string `json:"project" gorm:"column:project" binding:"required"`
	Version string `json:"version" gorm:"column:version" binding:"required"`
}

func (a *App) Key() string {
	return fmt.Sprintf(keyFormat, a.EnvId, a.Group, a.Project, a.Version)
}
