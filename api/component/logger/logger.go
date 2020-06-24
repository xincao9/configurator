package logger

import (
    "configurator/api/component/config"
    "github.com/sirupsen/logrus"
    "gopkg.in/natefinch/lumberjack.v2"
    "log"
    "path/filepath"
)

var (
	L *logrus.Logger
)

func init() {
	// 日志设置
	L = logrus.New()
	level, err := logrus.ParseLevel(config.C.GetString("logger.level"))
	if err != nil {
		log.Fatalf("Fatal error logger : %v\n", err)
	}
	fn := filepath.Join(config.C.GetString("logger.dir"), config.C.GetString("logger.filename"))
	L.Out = &lumberjack.Logger{
		Filename:   fn,
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   false,
	}
	L.SetLevel(level)
	L.Formatter = &logrus.TextFormatter{}
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetOutput(L.WriterLevel(logrus.InfoLevel))
}
