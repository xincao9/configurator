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

func (as *appService) Save(a *app.App) error {
    if a.Id <= 0 {
        return as.o.Save(a).Error
    }
    oa, err := as.GetAppById(a.Id) // 修改
    if err != nil {
        return err
    }
    if oa == nil {
        return gorm.ErrRecordNotFound
    }
    return as.o.Save(a).Error
}

func (as *appService) GetAppById(id int64) (*app.App, error) {
    var a app.App
    err := as.o.Where("`id`=?", id).First(&a).Error
    if err == gorm.ErrRecordNotFound {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return &a, nil
}

func (as *appService) GetAllApps() ([]app.App, error) {
    var apps []app.App
    err := as.o.Find(&apps).Error
    return apps, err
}

func (as *appService) Delete(id int64) error {
    err := as.o.Unscoped().Where("`id`=?", id).Delete(app.App{}).Error
    return err
}
