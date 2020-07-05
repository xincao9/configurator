package user_env

import (
    "configurator/api/component/util"
    "configurator/api/constant"
    "configurator/api/model/user"
    "configurator/api/model/user_env"
    userEnvService "configurator/api/service/user_env"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Route(engine *gin.RouterGroup) {
    save := func(c *gin.Context) {
        ue := &user_env.UserEnv{}
        if err := c.ShouldBindJSON(ue); err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
            return
        }
        err := userEnvService.U.Save(ue)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, ue)
    }
    engine.POST("/user_env", save)
    engine.PUT("/user_env", save)
    engine.GET("/user_env", func(c *gin.Context) {
        su, ok := c.Get(constant.SessionUser)
        if ok == false {
            util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
            return
        }
        u := su.(*user.User)
        userEnvs, err := userEnvService.U.GetUserEnvByUserId(u.Id)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, userEnvs)
    })
}
