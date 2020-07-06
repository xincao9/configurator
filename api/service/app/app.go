package app

import (
    "configurator/api/component/db"
    "configurator/api/component/logger"
    "configurator/api/model/app"
    envService "configurator/api/service/env"
    userEnvService "configurator/api/service/user_env"
    "fmt"
    "github.com/jinzhu/gorm"
    "github.com/xincao9/dkv/client"
    "time"
)

var (
    A *appService
)

func init() {
    var err error
    A, err = new()
    if err != nil {
        logger.L.Fatalf("")
    }
}

func new() (*appService, error) {
    c, err := client.New("localhost:9090", time.Second)
    if err != nil {
        return nil, err
    }
    return &appService{o: db.O, c: c}, nil
}

type appService struct {
    o *gorm.DB
    c *client.Client
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

func (as *appService) GetAppByEnv(env string) (apps []app.App, err error) {
    err = as.o.Where("`env`=?", env).Find(&apps).Error
    if err == gorm.ErrRecordNotFound {
        return
    }
    if err != nil {
        return
    }
    return
}

func (as *appService) GetAllApps() (apps []app.App, err error) {
    err = as.o.Find(&apps).Error
    return
}

func (as *appService) GetAppByUserId(userId int64) (apps []app.App, err error) {
    userEnvs, err := userEnvService.U.GetUserEnvByUserId(userId)
    if err != nil {
        return
    }
    for _, userEnv := range userEnvs {
        env, err := envService.E.GetEnvById(userEnv.EnvId)
        if err != nil {
            return apps, err
        }
        as, err := as.GetAppByEnv(env.Name)
        if err != nil {
            return apps, err
        }
        apps = append(apps, as...)
    }
    return
}

func (as *appService) Delete(id int64) error {
    return as.o.Unscoped().Where("`id`=?", id).Delete(app.App{}).Error
}

func (as *appService) GetPropertiesById(id int64) (string, error) {
    app, err := as.GetAppById(id)
    if err != nil {
        return "", err
    }
    r, err := as.c.GetOrRealtime(app.Key(), true)
    if err != nil {
        return "", err
    }
    if r.Code != 200 && r.Code != 404 {
        return "", fmt.Errorf("dkv response code: %d, message: %s\n", r.Code, r.Message)
    }
    return r.KV.V, nil
}

func (as *appService) SavePropertiesById(id int64, props string) (string, error) {
    app, err := as.GetAppById(id)
    if err != nil {
        return "", err
    }
    r, err := as.c.Put(app.Key(), props)
    if err != nil {
        return "", err
    }
    if r.Code != 200 {
        return "", fmt.Errorf("dkv response code: %d, message: %s\n", r.Code, r.Message)
    }
    return r.KV.V, nil
}
