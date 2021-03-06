package elog

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config ...
type Config struct {
	Debug               bool          // 配置文件控制日志输出到终端还是文件，true到终端，false到文件
	Dir                 string        // 日志输出目录，默认logs
	Name                string        // 日志文件名称，默认框架日志mocro.sys，业务日志default.log
	Level               string        // 日志初始等级，默认info级别
	AddCaller           bool          // 是否添加调用者信息，默认不加调用者信息
	MaxSize             int           // 日志输出文件最大长度，超过改值则截断，默认500M
	MaxAge              int           // 日志存储最大时间，默认最大保存天数为7天
	MaxBackup           int           // 日志存储最大数量，默认最大保存文件个数为10个
	Interval            time.Duration // 日志轮转时间，默认1天
	Async               bool          // 是否异步，默认异步
	FlushBufferSize     int           // 缓冲大小，默认256 * 1024B
	FlushBufferInterval time.Duration // 缓冲时间，默认5秒
	Fields              []zap.Field   // 日志初始化字段
	CallerSkip          int
	Core                zapcore.Core
	EncoderConfig       *zapcore.EncoderConfig
	configKey           string
}

// Filename ...
func (config *Config) Filename() string {
	return fmt.Sprintf("%s/%s", config.Dir, config.Name)
}

// DefaultConfig ...
func DefaultConfig() *Config {
	return &Config{
		Name:                DefaultLoggerName,
		Dir:                 "./logs",
		Level:               "info",
		FlushBufferSize:     defaultBufferSize,
		FlushBufferInterval: 5 * time.Second,
		MaxSize:             500, // 500M
		MaxAge:              1,   // 1 day
		MaxBackup:           10,  // 10 backup
		Interval:            24 * time.Hour,
		CallerSkip:          1,
		AddCaller:           false,
		Async:               true,
		EncoderConfig:       DefaultZapConfig(),
	}
}
