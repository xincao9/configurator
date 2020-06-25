package app

import (
    "configurator/api/component/util"
    "configurator/api/model/app"
    appService "configurator/api/service/app"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func Route(engine *gin.RouterGroup) {
    save := func(c *gin.Context) {
        a := &app.App{}
        if err := c.ShouldBindJSON(a); err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
            return
        }
        err := appService.A.Save(a)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSON(c, http.StatusOK, "成功")
    }
    engine.POST("/app", save)
    engine.PUT("/app", save)
    engine.GET("apps", func(c *gin.Context) {
        apps, err := appService.A.GetAllApps()
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, "成功", apps)
    })
    engine.GET("/app/:id", func(c *gin.Context) {
        n := c.Param("id")
        if n == "" {
            util.RenderJSON(c, http.StatusBadRequest, "id is required")
            return
        }
        id, err := strconv.ParseInt(n, 10, 64)
        if err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
            return
        }
        a, err := appService.A.GetAppById(id)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, "成功", a)
    })
    engine.DELETE("/app/:id", func(c *gin.Context) {
        n := c.Param("id")
        if n == "" {
            util.RenderJSON(c, http.StatusBadRequest, "id is required")
            return
        }
        id, err := strconv.ParseInt(n, 10, 64)
        if err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
            return
        }
        err = appService.A.Delete(id)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSON(c, http.StatusOK, "成功")
    })
}
