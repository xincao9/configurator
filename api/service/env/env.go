package env

import (
    "configurator/api/component/db"
    "configurator/api/model/env"
    "github.com/jinzhu/gorm"
)

var (
    E *envService
)

func init() {
    E = new()
}

type envService struct {
    o *gorm.DB
}

func new() *envService {
    return &envService{o: db.O}
}

func (es *envService) GetAllEnvs() (envs []env.Env, err error) {
    err = es.o.Find(&envs).Error
    return
}

func (es *envService) GetEnvById(id int64) (*env.Env, error) {
    var e env.Env
    err := es.o.Where("`id`=?", id).First(&e).Error
    if err == gorm.ErrRecordNotFound {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return &e, nil
}
