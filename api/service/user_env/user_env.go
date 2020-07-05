package user_env

import (
    "configurator/api/component/db"
    "github.com/jinzhu/gorm"
)

var (
    U *userEnvService
)

func init() {
    U = new()
}

func new() *userEnvService {
    return &userEnvService{
        o: db.O,
    }
}

type userEnvService struct {
    o *gorm.DB
}
