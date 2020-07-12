package app

import (
	"configurator/api/component/util"
	"configurator/api/constant"
	"configurator/api/model/app"
	"configurator/api/model/user"
	appService "configurator/api/service/app"
	envService "configurator/api/service/env"
	userEnvService "configurator/api/service/user_env"
	"encoding/json"
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
		a := &app.App{}
		if err := c.ShouldBindJSON(a); err != nil {
			util.RenderJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		ues, err := userEnvService.U.GetUserEnvsByUserId(u.Id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		if ues == nil {
			util.RenderJSON(c, http.StatusOK, constant.Success)
			return
		}
		ok = false
		for _, ue := range ues {
			e, err := envService.E.GetEnvById(ue.EnvId)
			if err != nil {
				util.RenderJSON(c, http.StatusInternalServerError, err.Error())
				return
			}
			if e == nil {
				continue
			}
			if e.Name == a.Env {
				ok = true
				break
			}
		}
		if ok == false {
			util.RenderJSON(c, http.StatusOK, constant.Success)
			return
		}
		err = appService.A.Save(a)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		util.RenderJSONDetail(c, http.StatusOK, constant.Success, a)
	}
	engine.POST("/app", save)
	engine.PUT("/app", save)
	engine.GET("/apps", func(c *gin.Context) {
		su, ok := c.Get(constant.SessionUser)
		if ok == false {
			util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
			return
		}
		u := su.(*user.User)
		apps, err := appService.A.GetAppsByUserId(u.Id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		util.RenderJSONDetail(c, http.StatusOK, constant.Success, apps)
	})
	engine.GET("/app/:id", func(c *gin.Context) {
		su, ok := c.Get(constant.SessionUser)
		if ok == false {
			util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
			return
		}
		u := su.(*user.User)
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
		a, err := appService.A.GetAppById(id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		ues, err := userEnvService.U.GetUserEnvsByUserId(u.Id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		if ues == nil {
			util.RenderJSON(c, http.StatusOK, constant.Success)
			return
		}
		ok = false
		for _, ue := range ues {
			e, err := envService.E.GetEnvById(ue.EnvId)
			if err != nil {
				util.RenderJSON(c, http.StatusInternalServerError, err.Error())
				return
			}
			if e == nil {
				continue
			}
			if e.Name == a.Env {
				ok = true
				break
			}
		}
		if ok == false {
			util.RenderJSON(c, http.StatusOK, constant.Success)
			return
		}
		util.RenderJSONDetail(c, http.StatusOK, constant.Success, a)
	})
	engine.DELETE("/app/:id", func(c *gin.Context) {
		su, ok := c.Get(constant.SessionUser)
		if ok == false {
			util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
			return
		}
		u := su.(*user.User)
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
		a, err := appService.A.GetAppById(id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		ues, err := userEnvService.U.GetUserEnvsByUserId(u.Id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		if ues == nil {
			util.RenderJSON(c, http.StatusOK, constant.Success)
			return
		}
		ok = false
		for _, ue := range ues {
			e, err := envService.E.GetEnvById(ue.EnvId)
			if err != nil {
				util.RenderJSON(c, http.StatusInternalServerError, err.Error())
				return
			}
			if e == nil {
				continue
			}
			if e.Name == a.Env {
				ok = true
				break
			}
		}
		if ok == false {
			util.RenderJSON(c, http.StatusOK, constant.Success)
			return
		}
		err = appService.A.Delete(id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		util.RenderJSON(c, http.StatusOK, constant.Success)
	})
	engine.GET("/app/:id/properties", func(c *gin.Context) {
		su, ok := c.Get(constant.SessionUser)
		if ok == false {
			util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
			return
		}
		u := su.(*user.User)
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
		a, err := appService.A.GetAppById(id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		ues, err := userEnvService.U.GetUserEnvsByUserId(u.Id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		if ues == nil {
			util.RenderJSON(c, http.StatusOK, constant.Success)
			return
		}
		ok = false
		for _, ue := range ues {
			e, err := envService.E.GetEnvById(ue.EnvId)
			if err != nil {
				util.RenderJSON(c, http.StatusInternalServerError, err.Error())
				return
			}
			if e == nil {
				continue
			}
			if e.Name == a.Env {
				ok = true
				break
			}
		}
		if ok == false {
			util.RenderJSON(c, http.StatusOK, constant.Success)
			return
		}
		props, err := appService.A.GetPropertiesById(id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		if props == "" {
			util.RenderJSON(c, http.StatusOK, constant.Success)
			return
		}
		var m map[string]interface{}
		err = json.Unmarshal([]byte(props), &m)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		d, err := json.MarshalIndent(m, "", "\t")
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		util.RenderJSONDetail(c, http.StatusOK, constant.Success, string(d))
	})
	saveProperties := func(c *gin.Context) {
		su, ok := c.Get(constant.SessionUser)
		if ok == false {
			util.RenderJSON(c, http.StatusInternalServerError, constant.SystemError)
			return
		}
		u := su.(*user.User)
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
		a, err := appService.A.GetAppById(id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		ues, err := userEnvService.U.GetUserEnvsByUserId(u.Id)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		if ues == nil {
			util.RenderJSON(c, http.StatusOK, constant.Success)
			return
		}
		ok = false
		for _, ue := range ues {
			e, err := envService.E.GetEnvById(ue.EnvId)
			if err != nil {
				util.RenderJSON(c, http.StatusInternalServerError, err.Error())
				return
			}
			if e == nil {
				continue
			}
			if e.Name == a.Env {
				ok = true
				break
			}
		}
		if ok == false {
			util.RenderJSON(c, http.StatusOK, constant.Success)
			return
		}
		var m map[string]interface{}
		if err := c.ShouldBindJSON(&m); err != nil {
			util.RenderJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		d, err := json.Marshal(m)
		if err != nil {
			util.RenderJSON(c, http.StatusBadRequest, err.Error())
		}
		props, err := appService.A.SavePropertiesById(id, string(d))
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		util.RenderJSONDetail(c, http.StatusOK, constant.Success, props)
	}
	engine.POST("/app/:id/properties", saveProperties)
	engine.PUT("/app/:id/properties", saveProperties)
}
