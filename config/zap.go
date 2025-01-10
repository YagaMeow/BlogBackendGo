package config

import (
	"go.uber.org/zap/zapcore"
)

/*
*

	level: info
	prefix: "[yagami_blog/server]"
	format: console
	director: log
	encode_level: LowercaseColorLeveleEncoder
	stacktrace_key: stacktrace
	max_age: 0
	show_line: true
	log_in_console: true
*/
type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Prefix        string `mapstructure:"prefix_string" json:"prefix_string" yaml:"prefix_string"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Director      string `mapstructure:"director" json:"director" yaml:"director"`
	EncodeLevel   string `mapstructure:"encode_level" json:"encode_level" yaml:"encode_level"`
	StacktraceKey string `mapstructure:"stacktrace_key" json:"stacktrace_key" yaml:"stacktrace_key"`
	MaxAge        int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	ShowLine      bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	LogInConsole  bool   `mapstructure:"log_in_console" json:"log_in_console" yaml:"log_in_console"`
}

func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	switch z.EncodeLevel {
	case "LowercaseLevelEncoder":
		return zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder":
		return zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder":
		return zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder":
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

func (z *Zap) TransportLevel() zapcore.Level {
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
