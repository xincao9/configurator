package operation_log

import (
    "configurator/api/component/db"
    "configurator/api/model/operation_log"
    "github.com/jinzhu/gorm"
)

var (
    O *operationLogService
)

func init() {
    O = new()
}

type operationLogService struct {
    o *gorm.DB
}

func new() *operationLogService {
    return &operationLogService{o: db.O}
}

func (ols *operationLogService) Save(ol *operation_log.OperationLog) error {
    if ol.Id <= 0 {
        return ols.o.Save(ol).Error
    }
    ool, err := ols.GetOperationLogById(ol.Id)
    if err != nil {
        return err
    }
    if ool == nil {
        return gorm.ErrRecordNotFound
    }
    return ols.o.Save(ol).Error
}

func (ols *operationLogService) GetAllOperationLogs() ([]operation_log.OperationLog, error) {
    var operationLogs []operation_log.OperationLog
    err := ols.o.Find(&operationLogs).Error
    return operationLogs, err
}

func (ols *operationLogService) GetOperationLogByUsername(username string) ([]operation_log.OperationLog, error) {
    var operationLogs []operation_log.OperationLog
    err := ols.o.Where("`username`=?", username).Find(&operationLogs).Error
    return operationLogs, err
}

func (ols *operationLogService) GetOperationLogById(id int64) (*operation_log.OperationLog, error) {
    var ol operation_log.OperationLog
    err := ols.o.Where("`id`=?", id).First(&ol).Error
    if err == gorm.ErrRecordNotFound {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return &ol, nil
}

func (ols *operationLogService) Delete(id int64) error {
    return ols.o.Unscoped().Where("`id`=?", id).Delete(operation_log.OperationLog{}).Error
}
