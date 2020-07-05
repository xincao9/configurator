package env

import "configurator/api/component/db"

type Env struct {
    db.Model
    Name string `json:"name" gorm:"column:name"`
}
