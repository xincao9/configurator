package session

import (
    "configurator/api/component/util"
    "configurator/api/constant"
    "configurator/api/model/session"
    sessionService "configurator/api/service/session"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "time"
)

func Route(engine *gin.Engine) {
	save := func(c *gin.Context) {
		s := &session.Session{}
		if err := c.ShouldBindJSON(s); err != nil {
			util.RenderJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		err := sessionService.S.Login(s) // 校验shareId
		if err != nil {
			util.RenderJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		s.Token = uuid.New().String()
		s.Expire = time.Now().Add(time.Hour * time.Duration(24))
		err = sessionService.S.Save(s) // 保存会话
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.SetCookie("token", s.Token, (int)((time.Hour*time.Duration(24))/time.Second), "/", "*", false, false)
		util.RenderJSONDetail(c, http.StatusOK, "成功", s)
	}
	engine.POST("/session", save)
	engine.PUT("/session", save)
}

func AuthenticationRoute(engine *gin.RouterGroup) {
	engine.DELETE("/session", func(c *gin.Context) {
		s, ok := c.Get(constant.Session)
		if ok == false {
			util.RenderJSON(c, http.StatusInternalServerError, "system error")
			return
		}
		ss := s.(*session.Session)
		ss, err := sessionService.S.GetSessionByUsername(ss.Username)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		ss.Expire = time.Now()
		util.RenderJSON(c, http.StatusOK, "成功")
	})
}
