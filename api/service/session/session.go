package session

import (
    "configurator/api/component/db"
    "configurator/api/model/session"
    "fmt"
    "github.com/jinzhu/gorm"
)

var (
	S *sessionService
)

func init() {
	S = new()
}

type sessionService struct {
	o *gorm.DB
}

func new() *sessionService {
	return &sessionService{o: db.O}
}

func (ss *sessionService) Login(s *session.Session) error {
    u := s.Username
    os, err := ss.GetSessionByUsername(u)
    if err != nil {
        return err
    }
    if os == nil {
        return fmt.Errorf("用户 %s 不存在", u)
    }
    if os.Password != s.Password {
        return fmt.Errorf("用户 %s 输入密码不正确", u)
    }
	return nil
}

func (ss *sessionService) Save(s *session.Session) error {
	u := s.Username
	os, err := ss.GetSessionByUsername(u)
	if err != nil {
		return err
	}
	if os == nil { // 第一次登录
		return ss.o.Save(s).Error
	}
	s.Id = os.Id
	err = ss.o.Save(s).Error
	return err
}

func (ss *sessionService) GetSessionByUsername(username string) (*session.Session, error) {
	s := &session.Session{}
	err := ss.o.Where("`username`=?", username).First(s).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (ss *sessionService) GetSessionByToken(token string) (*session.Session, error) {
	s := &session.Session{}
	err := ss.o.Where("`token`=?", token).First(s).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (ss *sessionService) Delete(id int64) error {
	return ss.o.Unscoped().Where("`id`=?", id).Delete(session.Session{}).Error
}
