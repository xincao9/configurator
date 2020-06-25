package authentication

import (
    "configurator/api/component/util"
    "configurator/api/constant"
    accountService "configurator/api/service/account"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

func Authentication(c *gin.Context) {
	t, err := c.Cookie("token") // 请求必须携带token
	if err != nil {
		c.Abort()
		util.RenderJSON(c, http.StatusBadRequest, "请求必须携带token")
		return
	}
	a, err := accountService.A.GetAccountByToken(t)
	if err != nil || a == nil || a.Expire.Before(time.Now()) { // 会话对象是否过期
		c.Abort()
		util.RenderJSON(c, http.StatusBadRequest, " 会话过期")
		return
	}
	c.Set(constant.SessionAccount, a) // 设置本地会话
}
