package db

import (
    "configurator/api/component/config"
    "configurator/api/component/logger"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "time"
)

var (
	O *gorm.DB
)

func init() {
	var err error
	O, err = gorm.Open("mysql", config.C.GetString("db.dataSourceName"))
	if err != nil {
		logger.L.Fatalf("Fatal error db: %v\n", err)
	}
	O.SingularTable(true)
}

type Model struct {
	Id        int64      `gorm:"primary_key"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
