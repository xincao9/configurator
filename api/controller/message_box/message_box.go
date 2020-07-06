package message_box

import (
	"configurator/api/component/util"
	"configurator/api/constant"
	"configurator/api/model/message_box"
	"configurator/api/model/user"
	messageBoxService "configurator/api/service/message_box"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(engine *gin.RouterGroup) {
	save := func(c *gin.Context) {
		mb := &message_box.MessageBox{}
		if err := c.ShouldBindJSON(mb); err != nil {
			util.RenderJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		err := messageBoxService.M.Save(mb)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		util.RenderJSONDetail(c, http.StatusOK, constant.Success, mb)
	}
	engine.POST("/message_box", save)
	engine.PUT("/message_box", save)
	engine.GET("/message_boxes", func(c *gin.Context) {
		su, _ := c.Get(constant.SessionUser)
		u, _ := su.(*user.User)
		mbs, err := messageBoxService.M.GetMessageBoxesByUsername(u.Username)
		if err != nil {
			util.RenderJSON(c, http.StatusInternalServerError, err.Error())
			return
		}
		util.RenderJSONDetail(c, http.StatusOK, constant.Success, mbs)
	})
}
