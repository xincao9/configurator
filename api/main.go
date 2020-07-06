package main

import (
	"configurator/api/component/config"
	"configurator/api/component/constant"
	"configurator/api/component/logger"
	_ "configurator/api/component/manager"
	"configurator/api/controller/app"
	"configurator/api/controller/env"
	"configurator/api/controller/message_box"
	"configurator/api/controller/operation_log"
	"configurator/api/controller/user"
	"configurator/api/controller/user_env"
	"configurator/api/middleware/authentication"
	"fmt"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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
	user.Route(engine)
	authorized := engine.Group("/", authentication.Authentication)
	user.AuthenticationRoute(authorized)
	app.Route(authorized)
	env.Route(authorized)
	user_env.Route(authorized)
	message_box.Route(authorized)
	operation_log.Route(authorized)
	engine.Static("/assets", config.C.GetString(constant.AssetsRootDir))
	engine.Static("/js", config.C.GetString(constant.AssetsJsDir))
	engine.Static("/css", config.C.GetString(constant.AssetsCssDir))
	engine.Static("/img", config.C.GetString(constant.AssetsImgDir))
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/assets")
	})
	addr := fmt.Sprintf(":%d", config.C.GetInt(constant.ServerPort))
	logger.L.Infof("Listening and serving HTTP on : %s", addr)
	if err := engine.Run(addr); err != nil {
		logger.L.Fatalf("Fatal error configurator-api: %v\n", err)
	}
}
