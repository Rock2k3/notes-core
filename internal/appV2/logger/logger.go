package logger

import (
	"github.com/Rock2k3/notes-core/internal/appV2/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type appLogger struct {
	logLevel       string
	loggerLevelMap map[string]zapcore.Level
	logger         *zap.Logger
}

func NewAppLogger(c *config.AppConfig) *appLogger {
	loggerLevelMap := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}

	return &appLogger{
		logLevel:       c.LogLevel(),
		loggerLevelMap: loggerLevelMap,
	}
}

func (l *appLogger) Init() {
	loggerLevel := l.getLoggerLevel()

	//var encoderConfig zapcore.EncoderConfig
	//if l.devMode {
	//	encoderConfig = zap.NewDevelopmentEncoderConfig()
	//} else {
	//	encoderConfig = zap.NewProductionEncoderConfig()
	//}
	encoderConfig := zap.NewDevelopmentEncoderConfig()

	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	consoleDebugging := zapcore.AddSync(os.Stdout)

	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  "./logs/app.log",
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, syncWriter, zap.NewAtomicLevelAt(loggerLevel)),
		zapcore.NewCore(consoleEncoder, consoleDebugging, zap.NewAtomicLevelAt(loggerLevel)),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.logger = logger
}

func (l appLogger) getLoggerLevel() zapcore.Level {
	loggerLevel, exist := l.loggerLevelMap[l.logLevel]
	if exist {
		return loggerLevel
	}
	return zapcore.InfoLevel
}

func (l appLogger) Logger() *zap.Logger {
	return l.logger
}
