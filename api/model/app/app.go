package app

import (
    "configurator/api/component/db"
)

type App struct {
    db.Model
    Env     string `json:"env" gorm:"column:env" binding:"required"`
    Group   string `json:"group" gorm:"column:group" binding:"required"`
    Project string `json:"project" gorm:"column:project" binding:"required"`
    Version string `json:"version" gorm:"column:version" binding:"required"`
}
