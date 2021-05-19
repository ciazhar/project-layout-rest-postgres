package logger

import (
	"go.uber.org/zap"
	"sync"
)

var once sync.Once

var sugaredLogger *zap.SugaredLogger

// Init returns an instance of logger
func Init() {

	once.Do(func() {
		config := Configuration{
			EnableConsole:     true,
			ConsoleLevel:      DebugD,
			ConsoleJSONFormat: true,
			EnableFile:        true,
			FileLevel:         InfoD,
			FileJSONFormat:    true,
			FileLocation:      "logs/app.log",
		}

		var err error
		sugaredLogger, err = newZapLogger(config)
		if err != nil {
			panic(err.Error())
		}
	})
}

func Info(format string, args ...interface{}) {
	sugaredLogger.Infof(format, args...)
}

func Error(format string, args ...interface{}) {
	sugaredLogger.Errorf(format, args...)
}
