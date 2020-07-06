package operation_log

import (
	"configurator/api/component/util"
	"configurator/api/constant"
	"configurator/api/model/operation_log"
	"configurator/api/model/user"
	operationLogService "configurator/api/service/operation_log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(engine *gin.RouterGroup) {
	engine.GET("/operation_logs", func(c *gin.Context) {
		su, _ := c.Get(constant.SessionUser)
		u, _ := su.(*user.User)
		var ols []operation_log.OperationLog
		var err error
		if u.Role == constant.RoleCommon {
			ols, err = operationLogService.O.GetOperationLogsByUsername(u.Username)
		} else {
			ols, err = operationLogService.O.GetAllOperationLogs()
		}
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		util.RenderJSONDetail(c, http.StatusOK, constant.Success, ols)
	})
}
