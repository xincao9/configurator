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
