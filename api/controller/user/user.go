package user

import (
    "configurator/api/component/util"
    "configurator/api/constant"
    "configurator/api/model/user"
    userService "configurator/api/service/user"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "strconv"
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
    engine.DELETE("/session/:id", func(c *gin.Context) {
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
        err = userService.U.Save(u)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSON(c, http.StatusOK, constant.Success)
    })
    save := func(c *gin.Context) {
        su, ok := c.Get(constant.SessionUser)
        if ok == false {
            util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
            return
        }
        u := su.(*user.User)
        nu := &user.User{}
        if err := c.ShouldBindJSON(nu); err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
            return
        }
        ok = false
        if u.Role == constant.RoleManager {
            if nu.Role == constant.RoleCommon {
                ok = true
            }
        } else if u.Role == constant.RoleAdmin {
            if nu.Role == constant.RoleCommon || nu.Role == constant.RoleManager || nu.Role == constant.RoleAdmin {
                ok = true
            }
        }
        if ok == false {
            util.RenderJSON(c, http.StatusOK, constant.Success)
            return
        }
        nu.Token = uuid.New().String()
        nu.Expire = time.Now()
        err := userService.U.Save(nu)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, nu)
    }
    engine.POST("/user", save)
    engine.PUT("/user", save)
    engine.GET("/user", func(c *gin.Context) {
        su, ok := c.Get(constant.SessionUser)
        if ok == false {
            util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
            return
        }
        u := su.(*user.User)
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, u)
    })
    engine.GET("/user/:id", func(c *gin.Context) {
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
        u, err := userService.U.GetUserById(id)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, u)
    })
    engine.GET("/users", func(c *gin.Context) {
        su, ok := c.Get(constant.SessionUser)
        if ok == false {
            util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
            return
        }
        u := su.(*user.User)
        if u.Role == constant.RoleCommon {
            util.RenderJSON(c, http.StatusOK, constant.Success)
            return
        }
        var users []user.User
        var err error
        if u.Role == constant.RoleManager {
            users, err = userService.U.GetUsersByRole(constant.RoleCommon)
        } else {
            users, err = userService.U.GetAllUsers()
        }
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, users)
    })
}
