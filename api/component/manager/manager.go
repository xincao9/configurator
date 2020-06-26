package manager

import (
    "configurator/api/component/config"
    "configurator/api/component/constant"
    "configurator/api/component/logger"
    "configurator/api/component/metrics"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    swaggerFiles "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
)

func init() {
	gin.SetMode(config.C.GetString("server.mode"))
	engine := gin.New()
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: logger.L.WriterLevel(logrus.DebugLevel)}), gin.RecoveryWithWriter(logger.L.WriterLevel(logrus.ErrorLevel)))
	config.Route(engine) // 配置服务接口
	metrics.Use(engine)
	engine.Static("/api", "./api")
	url := ginSwagger.URL("/api/swagger.yaml") // The url pointing to API definition
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	addr := fmt.Sprintf(":%d", config.C.GetInt(constant.ManagerServerPort))
	logger.L.Infof("Management listening and serving HTTP on : %s", addr)
	go func() {
		if err := engine.Run(addr); err != nil {
			logger.L.Fatalf("Fatal error manager: %v\n", err)
		}
	}()
}
