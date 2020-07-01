package operation_log

import (
    "configurator/api/component/util"
    operationLogService "configurator/api/service/operation_log"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Route(engine *gin.RouterGroup) {
    engine.GET("/operation_logs", func(c *gin.Context) {
        ols, err := operationLogService.O.GetAllOperationLogs()
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, "成功", ols)
    })
}
