package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

var (
	logger       *logrus.Logger
	loggerConfig = &LoggerConfig{
		Filename:     "logger.log",
		Format:       "text",
		NeedFileInfo: true,
		Level:        int(logrus.InfoLevel),
		Rotate: RotateConfig{
			MaxSize:    1,
			MaxBackups: 3,
			MaxAge:     1,
			Compress:   true,
		},
	}
)

type LoggerConfig struct {
	Filename     string       `yaml:"fileName"`
	Format       string       `yaml:"format"`
	NeedFileInfo bool         `yaml:"needFileInfo"`
	Level        int          `yaml:"level"`
	Rotate       RotateConfig `yaml:"rotate"`
}

type RotateConfig struct {
	MaxSize    int  `yaml:"maxSize"`
	MaxBackups int  `yaml:"maxBackups"`
	MaxAge     int  `yaml:"maxAge"`
	Compress   bool `yaml:"compress"`
}

func Init(c *LoggerConfig) {
	if logger != nil {
		return
	}
	logger = logrus.New()
	if c != nil {
		*loggerConfig = *c
	}
	if loggerConfig.Format == "json" {
		logger.Formatter = &logrus.JSONFormatter{}
	}
	logger.Out = &lumberjack.Logger{
		Filename:   loggerConfig.Filename,
		MaxSize:    loggerConfig.Rotate.MaxSize,
		MaxBackups: loggerConfig.Rotate.MaxBackups,
		MaxAge:     loggerConfig.Rotate.MaxAge,
		Compress:   loggerConfig.Rotate.Compress,
	}
	logger.Level = logrus.Level(loggerConfig.Level)
}

func SetLogLevel(level logrus.Level) {
	logger.Level = level
}

func SetLogFormatter(formatter logrus.Formatter) {
	logger.Formatter = formatter
}

func Debug(args ...interface{}) {
	if logger != nil {
		if logger.Level >= logrus.DebugLevel {
			entry := logger.WithFields(logrus.Fields{})
			if loggerConfig.NeedFileInfo {
				entry.Data["file"] = fileInfo(2)
			}
			entry.Debug(args)
		}
	}
}

func Debugf(format string, args ...interface{}) {
	if logger != nil {
		if logger.Level >= logrus.DebugLevel {
			entry := logger.WithFields(logrus.Fields{})
			if loggerConfig.NeedFileInfo {
				entry.Data["file"] = fileInfo(2)
			}
			entry.Debugf(format, args)
		}
	}
}

func Info(args ...interface{}) {
	if logger != nil {
		if logger.Level >= logrus.InfoLevel {
			entry := logger.WithFields(logrus.Fields{})
			if loggerConfig.NeedFileInfo {
				entry.Data["file"] = fileInfo(2)
			}
			entry.Info(args...)
		}
	}
}

func Infof(format string, args ...interface{}) {
	if logger != nil {
		if logger.Level >= logrus.InfoLevel {
			entry := logger.WithFields(logrus.Fields{})
			if loggerConfig.NeedFileInfo {
				entry.Data["file"] = fileInfo(2)
			}
			entry.Infof(format, args...)
		}
	}
}

func Warn(args ...interface{}) {
	if logger != nil {
		if logger.Level >= logrus.WarnLevel {
			entry := logger.WithFields(logrus.Fields{})
			if loggerConfig.NeedFileInfo {
				entry.Data["file"] = fileInfo(2)
			}
			entry.Warn(args...)
		}
	}
}

func Warnf(format string, args ...interface{}) {
	if logger != nil {
		if logger.Level >= logrus.WarnLevel {
			entry := logger.WithFields(logrus.Fields{})
			if loggerConfig.NeedFileInfo {
				entry.Data["file"] = fileInfo(2)
			}
			entry.Warnf(format, args...)
		}
	}
}

func Error(args ...interface{}) {
	if logger != nil {
		if logger.Level >= logrus.ErrorLevel {
			entry := logger.WithFields(logrus.Fields{})
			if loggerConfig.NeedFileInfo {
				entry.Data["file"] = fileInfo(2)
			}
			entry.Error(args...)
		}
	}
}

func Errorf(format string, args ...interface{}) {
	if logger != nil {
		if logger.Level >= logrus.ErrorLevel {
			entry := logger.WithFields(logrus.Fields{})
			if loggerConfig.NeedFileInfo {
				entry.Data["file"] = fileInfo(2)
			}
			entry.Errorf(format, args...)
		}
	}
}

func Fatal(args ...interface{}) {
	if logger != nil {
		if logger.Level >= logrus.FatalLevel {
			entry := logger.WithFields(logrus.Fields{})
			if loggerConfig.NeedFileInfo {
				entry.Data["file"] = fileInfo(2)
			}
			entry.Fatal(args...)
		}
	}
}

func Fatalf(format string, args ...interface{}) {
	if logger != nil {
		if logger.Level >= logrus.FatalLevel {
			entry := logger.WithFields(logrus.Fields{})
			if loggerConfig.NeedFileInfo {
				entry.Data["file"] = fileInfo(2)
			}
			entry.Fatalf(format, args...)
		}
	}
}

func Panic(args ...interface{}) {
	if logger != nil {
		if logger.Level >= logrus.PanicLevel {
			entry := logger.WithFields(logrus.Fields{})
			if loggerConfig.NeedFileInfo {
				entry.Data["file"] = fileInfo(2)
			}
			entry.Panic(args...)
		}
	}
}

func Panicf(format string, args ...interface{}) {
	if logger != nil {
		if logger.Level >= logrus.PanicLevel {
			entry := logger.WithFields(logrus.Fields{})
			if loggerConfig.NeedFileInfo {
				entry.Data["file"] = fileInfo(2)
			}
			entry.Panicf(format, args...)
		}
	}
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
