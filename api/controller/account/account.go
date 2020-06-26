package account

import (
    "configurator/api/component/util"
    "configurator/api/constant"
    "configurator/api/model/account"
    accountService "configurator/api/service/account"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "time"
)

func Route(engine *gin.Engine) {
    save := func(c *gin.Context) {
        a := &account.Account{}
        if err := c.ShouldBindJSON(a); err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
            return
        }
        err := accountService.A.Login(a) // 登录校验
        if err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
            return
        }
        a.Token = uuid.New().String()
        a.Expire = time.Now().Add(time.Hour * time.Duration(constant.SessionExpireHour))
        err = accountService.A.Save(a) // 更新登录信息
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        c.SetCookie(constant.Token, a.Token, (int)((time.Hour*time.Duration(constant.SessionExpireHour))/time.Second), "/", "*", false, false)
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, a)
    }
    engine.POST("/account", save)
    engine.PUT("/account", save)
}

func AuthenticationRoute(engine *gin.RouterGroup) {
    engine.DELETE("/account", func(c *gin.Context) {
        sa, ok := c.Get(constant.SessionAccount)
        if ok == false {
            util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
            return
        }
        a := sa.(*account.Account)
        a, err := accountService.A.GetAccountByUsername(a.Username)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        a.Expire = time.Now()
        util.RenderJSON(c, http.StatusOK, constant.Success)
    })
}
