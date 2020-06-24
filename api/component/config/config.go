package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

var (
	C   *viper.Viper
	Env = os.Getenv("Env")
)

func init() {
	C = viper.New()
	if Env == "" {
		C.SetConfigName("config")
	} else {
		C.SetConfigName(fmt.Sprintf("config-%s", Env))
	}
	C.SetConfigType("yaml")
	C.AddConfigPath("./resources/conf")
    C.AddConfigPath("/usr/local/configurator-api/conf/")
	C.SetDefault("logger.dir", "/tmp/logs")
	C.SetDefault("logger.level", "debug")
	C.SetDefault("logger.filename", "configurator-api.log")
	C.SetDefault("server.mode", "debug")
	C.SetDefault("server.port", 8080)
	C.SetDefault("server.cors.accessControlAllowOrigin", "http://localhost:8081")
	C.SetDefault("manager.server.port", 8081)
	C.SetDefault("db.dataSourceName", "root:asdf@tcp(localhost:3306)/configurator?charset=utf8&parseTime=true")
	C.SetDefault("gpool.size", 10000)
	C.SetDefault("email.address", "smtp.exmail.qq.com:587")
	C.SetDefault("email.username", "")
	C.SetDefault("email.password", "")
	C.SetDefault("email.host", "smtp.exmail.qq.com")
	C.SetDefault("email.pool.maxSize", 4)
	err := C.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error conf : %v\n", err)
	}
}

func Route(engine *gin.Engine) {
	engine.GET("/conf", func(c *gin.Context) {
		c.JSON(http.StatusOK, C.AllSettings())
	})
}
