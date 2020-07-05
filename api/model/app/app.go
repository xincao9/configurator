package app

import (
    "configurator/api/component/db"
    "fmt"
)

const (
    keyFormat = "configurator|%s|%s|%s|%s"
)

type App struct {
    db.Model
    Env     string `json:"env" gorm:"column:env" binding:"required"`
    Group   string `json:"group" gorm:"column:group" binding:"required"`
    Project string `json:"project" gorm:"column:project" binding:"required"`
    Version string `json:"version" gorm:"column:version" binding:"required"`
}

func (a *App) Key() string {
    return fmt.Sprintf(keyFormat, a.Env, a.Group, a.Project, a.Version)
}
