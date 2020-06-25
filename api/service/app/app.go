package app

import (
    "configurator/api/component/db"
    "configurator/api/model/app"
    "github.com/jinzhu/gorm"
)

var (
    A *appService
)

func init () {
    A = new()
}

func new () *appService {
    return &appService{o: db.O}
}

type appService struct {
    o *gorm.DB
}

func (as *appService) GetAllApps() ([]app.App, error) {
    var apps []app.App
    err := as.o.Find(&apps).Error
    return apps, err
}
