package app

import (
    "configurator/api/component/util"
    "configurator/api/constant"
    "configurator/api/model/app"
    appService "configurator/api/service/app"
    "encoding/json"
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
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, a)
    }
    engine.POST("/app", save)
    engine.PUT("/app", save)
    engine.GET("apps", func(c *gin.Context) {
        apps, err := appService.A.GetAllApps()
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, apps)
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
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, a)
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
        util.RenderJSON(c, http.StatusOK, constant.Success)
    })
    engine.GET("/app/:id/properties", func(c *gin.Context) {
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
        props, err := appService.A.GetPropertiesById(id)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        var m map[string]interface{}
        err = json.Unmarshal([]byte(props), &m)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        d, err := json.MarshalIndent(m, "", "\t")
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, string(d))
    })
    saveProperties := func(c *gin.Context) {
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
        var m map[string]interface{}
        if err := c.ShouldBindJSON(&m); err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
            return
        }
        d, err := json.Marshal(m)
        if err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
        }
        props, err := appService.A.SavePropertiesById(id, string(d))
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, props)
    }
    engine.POST("/app/:id/properties", saveProperties)
    engine.PUT("/app/:id/properties", saveProperties)
}
