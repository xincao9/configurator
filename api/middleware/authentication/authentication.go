package authentication

import (
    "configurator/api/component/util"
    "configurator/api/constant"
    sessionService "configurator/api/service/session"
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
	s, err := sessionService.S.GetSessionByToken(t)
	if err != nil || s == nil || s.Expire.Before(time.Now()) { // 会话对象是否过期
		c.Abort()
		util.RenderJSON(c, http.StatusBadRequest, " 会话过期")
		return
	}
	c.Set(constant.Session, s) // 设置本地会话
}
