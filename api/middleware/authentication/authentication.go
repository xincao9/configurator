package authentication

import (
    "configurator/api/component/util"
    "configurator/api/constant"
    "configurator/api/model/operation_log"
    accountService "configurator/api/service/account"
    operationLogService "configurator/api/service/operation_log"
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

func Authentication(c *gin.Context) {
    t, err := c.Cookie(constant.Token) // 请求必须携带token
    if err != nil {
        c.Abort()
        util.RenderJSON(c, http.StatusBadRequest, "the request must carry a token")
        return
    }
    a, err := accountService.A.GetAccountByToken(t)
    if err != nil || a == nil || a.Expire.Before(time.Now()) { // 会话对象是否过期
        c.Abort()
        util.RenderJSON(c, http.StatusBadRequest, "session expired or nonexistent session")
        return
    }
    c.Set(constant.SessionAccount, a) // 设置本地会话
    if c.Request.Method == "GET" {
        return
    } else if c.Request.Method == "DELETE" || c.Request.Method == "POST" || c.Request.Method == "PUT" {
        operationLogService.O.Save(&operation_log.OperationLog{
            Username: a.Username,
            Message:  fmt.Sprintf("method = %s, uri = %s", c.Request.Method, c.Request.URL.RequestURI()),
        })
    }
}
