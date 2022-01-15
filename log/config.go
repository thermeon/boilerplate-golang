package log

import "github.com/thermeon/boilerplate-golang/common"

const (
	DebugLevel   common.LogLevel = "debug"
	WarningLevel common.LogLevel = "warning"
	InfoLevel    common.LogLevel = "info"
	ErrorLevel   common.LogLevel = "error"
	FatalLevel   common.LogLevel = "fatal"
	PanicLevel   common.LogLevel = "panic"
)

type Config struct {
	Level common.LogLevel
}

func GetDefaultConfigs() *Config {
	return &Config{
		Level: InfoLevel,
	}
}
