package user_env

import (
	"configurator/api/component/db"
	"configurator/api/model/user_env"
	"github.com/jinzhu/gorm"
)

var (
	U *userEnvService
)

func init() {
	U = new()
}

func new() *userEnvService {
	return &userEnvService{
		o: db.O,
	}
}

type userEnvService struct {
	o *gorm.DB
}

func (ues *userEnvService) Save(ue *user_env.UserEnv) error {
	if ue.Id <= 0 {
	    userEnvs, err := ues.GetUserEnvsByUserId(ue.UserId)
	    if err != nil {
	        return err
        }
        for _, userEnv := range userEnvs {
            if userEnv.UserId == ue.UserId && userEnv.EnvId == ue.EnvId {
                return nil
            }
        }
		return ues.o.Save(ue).Error
	}
	oue, err := ues.GetUserEnvById(ue.Id) // 修改
	if err != nil {
		return err
	}
	if oue == nil {
		return gorm.ErrRecordNotFound
	}
	return ues.o.Save(ue).Error
}

func (ues *userEnvService) GetUserEnvById(id int64) (*user_env.UserEnv, error) {
	var ue user_env.UserEnv
	err := ues.o.Where("`id`=?", id).First(&ue).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &ue, nil
}

func (ues *userEnvService) GetUserEnvsByUserId(userId int64) ([]user_env.UserEnv, error) {
	var userEnvs []user_env.UserEnv
	err := ues.o.Where("`user_id`=?", userId).Find(&userEnvs).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return userEnvs, nil
}
