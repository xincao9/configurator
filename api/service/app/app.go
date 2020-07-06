package app

import (
	"configurator/api/component/config"
	"configurator/api/component/constant"
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
		logger.L.Fatalf("Fatal error appService: %v\n", err)
	}
}

func new() (*appService, error) {
	c, err := client.New(config.C.GetString(constant.DKVAddress), time.Second)
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
	oa, err := as.GetAppById(a.Id)
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

func (as *appService) GetAppsByEnv(env string) ([]app.App, error) {
	var apps []app.App
	err := as.o.Where("`env`=?", env).Find(&apps).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func (as *appService) GetAllApps() ([]app.App, error) {
	var apps []app.App
	err := as.o.Find(&apps).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func (as *appService) GetAppsByUserId(userId int64) ([]app.App, error) {
	var apps []app.App
	userEnvs, err := userEnvService.U.GetUserEnvsByUserId(userId)
	if err != nil {
		return nil, err
	}
	if userEnvs == nil {
		return nil, nil
	}
	for _, userEnv := range userEnvs {
		env, err := envService.E.GetEnvById(userEnv.EnvId)
		if err != nil {
			return nil, err
		}
		if env == nil {
			continue
		}
		as, err := as.GetAppsByEnv(env.Name)
		if err != nil {
			return nil, err
		}
		if as == nil {
			continue
		}
		apps = append(apps, as...)
	}
	return apps, nil
}

func (as *appService) Delete(id int64) error {
	return as.o.Unscoped().Where("`id`=?", id).Delete(app.App{}).Error
}

func (as *appService) GetPropertiesById(id int64) (string, error) {
	app, err := as.GetAppById(id)
	if err != nil {
		return "", err
	}
	if app == nil {
		return "", nil
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
	if app == nil {
		return "", nil
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
