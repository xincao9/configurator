package account

import (
    "configurator/api/component/db"
    "configurator/api/model/account"
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "github.com/jinzhu/gorm"
)

var (
    A *accountService
)

func init() {
    A = new()
}

type accountService struct {
    o *gorm.DB
}

func new() *accountService {
    return &accountService{o: db.O}
}

func (as *accountService) Login(a *account.Account) error {
    u := a.Username
    oa, err := as.GetAccountByUsername(u)
    if err != nil {
        return err
    }
    if oa == nil {
        return fmt.Errorf("user %s does not exist", u)
    }
    h := md5.New()
    h.Write([]byte(a.Password))
    a.Password = hex.EncodeToString(h.Sum(nil))
    if oa.Password != a.Password {
        return fmt.Errorf("user %s password is incorrect", u)
    }
    return nil
}

func (as *accountService) Save(a *account.Account) error {
    u := a.Username
    oa, err := as.GetAccountByUsername(u)
    if err != nil {
        return err
    }
    if oa == nil { // 第一次登录
        h := md5.New()
        h.Write([]byte(a.Password))
        p := hex.EncodeToString(h.Sum(nil))
        a.Password = p
        return as.o.Save(a).Error
    }
    a.Id = oa.Id
    err = as.o.Save(a).Error
    return err
}

func (as *accountService) GetAccountByUsername(username string) (*account.Account, error) {
    a := &account.Account{}
    err := as.o.Where("`username`=?", username).First(a).Error
    if err == gorm.ErrRecordNotFound {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return a, nil
}

func (as *accountService) GetAccountByToken(token string) (*account.Account, error) {
    a := &account.Account{}
    err := as.o.Where("`token`=?", token).First(a).Error
    if err == gorm.ErrRecordNotFound {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return a, nil
}

func (as *accountService) GetAllAccounts() ([]account.Account, error) {
    var accounts []account.Account
    err := as.o.Find(&accounts).Error
    return accounts, err
}

func (as *accountService) Delete(id int64) error {
    return as.o.Unscoped().Where("`id`=?", id).Delete(account.Account{}).Error
}
