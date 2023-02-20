package QUtil

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

type LoggerConfig struct {
	FileName   string //文件地址
	MaxSize    int    //分割文件大小 单位:MB
	MaxBackups int    //保留旧文件个数
	MaxAge     int    //保留旧文件天数
	Compress   bool   //是否压缩/归档旧文件
}

// GetLogger 获取日志对象
func GetLogger(cfg ...LoggerConfig) *zap.SugaredLogger {
	if logger == nil {
		newLogger(cfg...)
	}
	return logger
}

func newLogger(cfg ...LoggerConfig) {
	logWiter := getLogWriter(cfg...)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, logWiter, zapcore.DebugLevel)
	l := zap.New(core, zap.AddCaller())
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			print(111111)
		}
	}(l)
	logger = l.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConf := zap.NewProductionEncoderConfig()
	encoderConf.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConf)
	//return zapcore.NewConsoleEncoder(encoderConf)
}

func getLogWriter(cfg ...LoggerConfig) zapcore.WriteSyncer {

	option := &lumberjack.Logger{}
	if len(cfg) > 0 {
		for _, v := range cfg {
			option = &lumberjack.Logger{
				Filename:   v.FileName,
				MaxSize:    v.MaxSize,
				MaxAge:     v.MaxAge,
				MaxBackups: v.MaxBackups,
				LocalTime:  false,
				Compress:   false,
			}
		}
	} else {
		option = &lumberjack.Logger{
			Filename:   "./log/logger.log",
			MaxSize:    30,
			MaxAge:     30,
			MaxBackups: 60,
			LocalTime:  false,
			Compress:   false,
		}

	}
	lumberJackLogger := option
	return zapcore.AddSync(lumberJackLogger)
}
