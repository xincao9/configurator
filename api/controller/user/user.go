package user

import (
    "configurator/api/component/util"
    "configurator/api/constant"
    "configurator/api/model/user"
    userService "configurator/api/service/user"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "time"
)

func Route(engine *gin.Engine) {
    save := func(c *gin.Context) {
        u := &user.User{}
        if err := c.ShouldBindJSON(u); err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
            return
        }
        err := userService.U.Login(u) // 登录校验
        if err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
            return
        }
        u.Token = uuid.New().String()
        u.Expire = time.Now().Add(time.Hour * time.Duration(constant.SessionExpireHour))
        err = userService.U.Save(u) // 更新登录信息
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        c.SetCookie(constant.Token, u.Token, (int)((time.Hour*time.Duration(constant.SessionExpireHour))/time.Second), "/", "*", false, false)
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, u)
    }
    engine.PUT("/session", save)
    engine.POST("/session", save)
}

func AuthenticationRoute(engine *gin.RouterGroup) {
    save := func(c *gin.Context) {
        u := &user.User{}
        if err := c.ShouldBindJSON(u); err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
            return
        }
        u.Token = uuid.New().String()
        u.Expire = time.Now()
        err := userService.U.Save(u)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, u)
    }
    engine.POST("/user", save)
    engine.PUT("/user", save)
    engine.DELETE("/session", func(c *gin.Context) {
        su, ok := c.Get(constant.SessionUser)
        if ok == false {
            util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
            return
        }
        u := su.(*user.User)
        u, err := userService.U.GetUserByUsername(u.Username)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        u.Expire = time.Now()
        util.RenderJSON(c, http.StatusOK, constant.Success)
    })
    engine.GET("/users", func(c *gin.Context) {
        users, err := userService.U.GetAllUsers()
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, users)
    })
}
