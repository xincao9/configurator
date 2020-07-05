package env

import (
    "configurator/api/component/util"
    envService "configurator/api/service/env"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Route(engine *gin.RouterGroup) {
    engine.GET("/envs", func(c *gin.Context) {
        envs, err := envService.E.GetAllEnvs()
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, "成功", envs)
    })
}
