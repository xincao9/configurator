package app

import (
    "configurator/api/component/util"
    appService "configurator/api/service/app"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Route(engine *gin.RouterGroup) {
    engine.GET("apps", func(c *gin.Context) {
        apps, err := appService.A.GetAllApps()
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, "成功", apps)
    })
}
