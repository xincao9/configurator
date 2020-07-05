package authentication

import (
    "configurator/api/component/util"
    "configurator/api/constant"
    "configurator/api/model/operation_log"
    operationLogService "configurator/api/service/operation_log"
    userService "configurator/api/service/user"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

func Authentication(c *gin.Context) {
    t, err := c.Cookie(constant.Token) // 请求必须携带token
    if err != nil {
        c.Abort()
        util.RenderJSON(c, http.StatusBadRequest, "the request must carry u token")
        return
    }
    u, err := userService.U.GetUserByToken(t)
    if err != nil || u == nil || u.Expire.Before(time.Now()) { // 会话对象是否过期
        c.Abort()
        util.RenderJSON(c, http.StatusBadRequest, "session expired or nonexistent session")
        return
    }
    c.Set(constant.SessionUser, u) // 设置本地会话
    if c.Request.Method == "GET" {
        return
    } else if c.Request.Method == "DELETE" || c.Request.Method == "POST" || c.Request.Method == "PUT" {
        operationLogService.O.Save(&operation_log.OperationLog{
            Username: u.Username,
            Message:  fmt.Sprintf("method = %s, uri = %s", c.Request.Method, c.Request.URL.RequestURI()),
        })
    }
}
