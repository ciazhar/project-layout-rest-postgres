package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const (
	//DebugD has verbose message
	DebugD = "debug"
	//InfoD is default log level
	InfoD = "info"
	//WarnD is for logging messages about possible issues
	WarnD = "warn"
	//ErrorD is for logging errors
	ErrorD = "error"
	//FatalD is for logging fatal messages. The sytem shutsdown after logging the message.
	FatalD = "fatal"
)

type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case InfoD:
		return zapcore.InfoLevel
	case WarnD:
		return zapcore.WarnLevel
	case DebugD:
		return zapcore.DebugLevel
	case ErrorD:
		return zapcore.ErrorLevel
	case FatalD:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func newZapLogger(config Configuration) (*zap.SugaredLogger, error) {
	var cores []zapcore.Core

	if config.EnableConsole {
		level := getZapLevel(config.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(config.ConsoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	if config.EnableFile {
		level := getZapLevel(config.FileLevel)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: config.FileLocation,
			MaxSize:  100,
			Compress: true,
			MaxAge:   28,
		})
		core := zapcore.NewCore(getEncoder(config.FileJSONFormat), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
		zap.Fields(zap.String("service", "project-layout-rest-postgres")),
	).Sugar()

	return logger, nil
}
