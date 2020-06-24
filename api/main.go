package main

import (
    "configurator/api/component/config"
    "configurator/api/component/logger"
    "configurator/api/controller/session"
    "configurator/api/middleware/authentication"
    "fmt"
    "github.com/gin-gonic/gin"
    cors "github.com/rs/cors/wrapper/gin"
    "github.com/sirupsen/logrus"
    "net/http"
)

func main() {
    gin.SetMode(config.C.GetString("server.mode"))
    engine := gin.New()
    engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: logger.L.WriterLevel(logrus.DebugLevel)}), gin.RecoveryWithWriter(logger.L.WriterLevel(logrus.ErrorLevel)))
    engine.Use(cors.New(cors.Options{
        AllowedOrigins: []string{config.C.GetString("server.cors.accessControlAllowOrigin")},
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
    session.Route(engine)
    authorized := engine.Group("/", authentication.Authentication)
    session.AuthenticationRoute(authorized)
    engine.Static("/assets", "./assets")
    engine.Static("/js", "./assets/js")
    engine.Static("/css", "./assets/css")
    engine.Static("/img", "./assets/img")
    engine.GET("/", func(c *gin.Context) {
        c.Redirect(http.StatusFound, "/assets")
    })
    addr := fmt.Sprintf(":%d", config.C.GetInt("server.port"))
    logger.L.Infof("Listening and serving HTTP on : %s", addr)
    if err := engine.Run(addr); err != nil {
        logger.L.Fatalf("Fatal error configurator-api: %v\n", err)
    }
}
