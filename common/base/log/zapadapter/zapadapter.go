package zapadapter

import (
	"bytes"
	"os"
	"project/common/tool"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type adapter struct {
	zapLogger *zap.Logger
}

var (
	adapterInstance *adapter
	loadAdapterOnce sync.Once
)

// New new log adapter
func New(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) *adapter {
	loadAdapterOnce.Do(func() {
		adapterInstance = &adapter{}
		core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
		adapterInstance.zapLogger = zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", serviceName)))
	})
	return adapterInstance
}

// NewCustom new custom log adapter
func NewCustom(fileName, filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) *adapter {
	var buffer bytes.Buffer
	buffer.WriteString(filePath)
	buffer.WriteString(tool.GetFileNameWithTimestamp(fileName))
	fileFullPath := buffer.String()
	core := newCore(fileFullPath, level, maxSize, maxBackups, maxAge, compress)
	adapterInstance = &adapter{}
	adapterInstance.zapLogger = zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", serviceName)))
	return adapterInstance
}

func newCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
	//日志文件路径配置
	hook := lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
}

func (log *adapter) Debug(msg string) {
	log.zapLogger.Debug(msg)
}

func (log *adapter) Info(msg string) {
	log.zapLogger.Info(msg)
}

func (log *adapter) Warn(msg string) {
	log.zapLogger.Warn(msg)
}

func (log *adapter) Error(msg string) {
	log.zapLogger.Error(msg)
}

func (log *adapter) Panic(msg string) {
	log.zapLogger.Panic(msg)
}

func (log *adapter) Fatal(msg string) {
	log.zapLogger.Fatal(msg)
}
