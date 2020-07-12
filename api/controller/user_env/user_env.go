package user_env

import (
	"configurator/api/component/util"
	"configurator/api/constant"
	"configurator/api/model/user"
	"configurator/api/model/user_env"
	userEnvService "configurator/api/service/user_env"
	"github.com/gin-gonic/gin"
	"net/http"
    "strconv"
)

func Route(engine *gin.RouterGroup) {
	save := func(c *gin.Context) {
		su, ok := c.Get(constant.SessionUser)
		if ok == false {
			util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
			return
		}
		u := su.(*user.User)
		ue := &user_env.UserEnv{}
		if err := c.ShouldBindJSON(ue); err != nil {
			util.RenderJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		ok = false
		if u.Role == constant.RoleManager || u.Role == constant.RoleAdmin {
			ues, err := userEnvService.U.GetUserEnvsByUserId(u.Id)
			if err != nil {
				util.RenderJSON(c, http.StatusInternalServerError, err.Error())
				return
			}
			if ues == nil {
				util.RenderJSON(c, http.StatusOK, constant.Success)
				return
			}
			for _, userEnv := range ues {
				if userEnv.EnvId == ue.EnvId {
					ok = true
					break
				}
			}
		}
		if ok == false {
			util.RenderJSON(c, http.StatusOK, constant.Success)
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
		userEnvs, err := userEnvService.U.GetUserEnvsByUserId(u.Id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		util.RenderJSONDetail(c, http.StatusOK, constant.Success, userEnvs)
	})
    engine.GET("/user_env/user_id/:user_id", func(c *gin.Context) {
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
        n := c.Param("user_id")
        if n == "" {
            util.RenderJSON(c, http.StatusBadRequest, "user_id is required")
            return
        }
        userId, err := strconv.ParseInt(n, 10, 64)
        if err != nil {
            util.RenderJSON(c, http.StatusBadRequest, err.Error())
            return
        }
        userEnvs, err := userEnvService.U.GetUserEnvsByUserId(userId)
        if err != nil {
            util.RenderJSON(c, http.StatusInternalServerError, err.Error())
            return
        }
        util.RenderJSONDetail(c, http.StatusOK, constant.Success, userEnvs)
    })
}
