package message_box

import (
    "configurator/api/constant"
    "configurator/api/model/message_box"
    "testing"
)

func TestMessageBoxService_Save(t *testing.T) {
    mb := &message_box.MessageBox{
        Username: "caoxin",
        Message:  "申请已经完成",
        Status:   constant.StatusNormal,
    }
    err := M.Save(mb)
    if err != nil {
        t.Error(err)
    }
    M.Delete(mb.Id)
}

func TestMessageBoxService_GetAllMessageBoxs(t *testing.T) {
    mb := &message_box.MessageBox{
        Username: "caoxin",
        Message:  "申请已经完成",
        Status:   constant.StatusNormal,
    }
    err := M.Save(mb)
    if err != nil {
        t.Error(err)
    }
    mbs, err := M.GetAllMessageBoxes()
    if err != nil {
        t.Error(err)
    }
    t.Log(mbs)
    M.Delete(mb.Id)
}

func TestMessageBoxService_GetMessageBoxByUsername(t *testing.T) {
    mb := &message_box.MessageBox{
        Username: "caoxin",
        Message:  "申请已经完成",
        Status:   constant.StatusNormal,
    }
    err := M.Save(mb)
    if err != nil {
        t.Error(err)
    }
    mbs, err := M.GetMessageBoxByUsername("caoxin")
    if err != nil {
        t.Error(err)
    }
    t.Log(mbs)
    M.Delete(mb.Id)
}
