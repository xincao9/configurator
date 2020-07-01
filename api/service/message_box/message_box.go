package message_box

import (
    "configurator/api/component/db"
    "configurator/api/constant"
    "configurator/api/model/message_box"
    "github.com/jinzhu/gorm"
)

var (
    M *messageBoxService
)

func init() {
    M = new()
}

type messageBoxService struct {
    o *gorm.DB
}

func new() *messageBoxService {
    return &messageBoxService{o: db.O}
}

func (mbs *messageBoxService) Save(mb *message_box.MessageBox) error {
    if mb.Id <= 0 {
        return mbs.o.Save(mb).Error
    }
    omb, err := mbs.GetMessageBoxById(mb.Id)
    if err != nil {
        return err
    }
    if omb == nil {
        return gorm.ErrRecordNotFound
    }
    return mbs.o.Save(mb).Error
}

func (mbs *messageBoxService) GetAllMessageBoxes() ([]message_box.MessageBox, error) {
    var messageBoxes []message_box.MessageBox
    err := mbs.o.Find(&messageBoxes).Error
    return messageBoxes, err
}

func (mbs *messageBoxService) GetMessageBoxByUsername(username string) ([]message_box.MessageBox, error) {
    var messageBoxes []message_box.MessageBox
    err := mbs.o.Where("`username`=? and `status`=?", username, constant.StatusNormal).Find(&messageBoxes).Error
    return messageBoxes, err
}

func (mbs *messageBoxService) GetMessageBoxById(id int64) (*message_box.MessageBox, error) {
    var mb message_box.MessageBox
    err := mbs.o.Where("`id`=?", id).First(&mb).Error
    if err == gorm.ErrRecordNotFound {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return &mb, nil
}

func (mbs *messageBoxService) Delete(id int64) error {
    return mbs.o.Unscoped().Where("`id`=?", id).Delete(message_box.MessageBox{}).Error
}
