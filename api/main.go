package main

import (
	"configurator/api/component/config"
    "configurator/api/component/constant"
    "configurator/api/component/logger"
	_ "configurator/api/component/manager"
	"configurator/api/controller/account"
	"configurator/api/controller/app"
	"configurator/api/middleware/authentication"
	"fmt"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	assetsDir = "/usr/local/configurator/assets"
)

var (
	jsDir  = fmt.Sprintf("%s/js", assetsDir)
	cssDir = fmt.Sprintf("%s/css", assetsDir)
	imgDir = fmt.Sprintf("%s/img", assetsDir)
)

func main() {
	gin.SetMode(config.C.GetString(constant.ServerMode))
	engine := gin.New()
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: logger.L.WriterLevel(logrus.DebugLevel)}), gin.RecoveryWithWriter(logger.L.WriterLevel(logrus.ErrorLevel)))
	engine.Use(cors.New(cors.Options{
		AllowedOrigins: []string{config.C.GetString(constant.ServerCorsAccessControlAllowOrigin)},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))
	account.Route(engine)
	authorized := engine.Group("/", authentication.Authentication)
	account.AuthenticationRoute(authorized)
	app.Route(authorized)
	engine.Static("/assets", assetsDir)
	engine.Static("/js", jsDir)
	engine.Static("/css", cssDir)
	engine.Static("/img", imgDir)
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/assets")
	})
	addr := fmt.Sprintf(":%d", config.C.GetInt(constant.ServerPort))
	logger.L.Infof("Listening and serving HTTP on : %s", addr)
	if err := engine.Run(addr); err != nil {
		logger.L.Fatalf("Fatal error configurator-api: %v\n", err)
	}
}
