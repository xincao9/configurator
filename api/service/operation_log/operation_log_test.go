package operation_log

import (
    "configurator/api/model/operation_log"
    "testing"
)

func TestOperationLogService_Save(t *testing.T) {
    ol := &operation_log.OperationLog{
        Username: "caoxin",
        Message:  "新建表",
    }
    err := O.Save(ol)
    if err != nil {
        t.Error(err)
    }
    O.Delete(ol.Id)
}

func TestOperationLogService_GetAllOperationLogs(t *testing.T) {
    ol := &operation_log.OperationLog{
        Username: "caoxin",
        Message:  "新建表",
    }
    err := O.Save(ol)
    if err != nil {
        t.Error(err)
    }
    ols, err := O.GetAllOperationLogs()
    if err != nil {
        t.Error(err)
    }
    t.Log(ols)
    O.Delete(ol.Id)
}

func TestOperationLogService_GetOperationLogByUsername(t *testing.T) {
    ol := &operation_log.OperationLog{
        Username: "caoxin",
        Message:  "新建表",
    }
    err := O.Save(ol)
    if err != nil {
        t.Error(err)
    }
    ols, err := O.GetOperationLogByUsername("caoxin")
    if err != nil {
        t.Error(err)
    }
    t.Log(ols)
    O.Delete(ol.Id)
}
