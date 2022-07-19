package logger

import (
	"github.com/Rock2k3/notes-core/internal/appV2/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type AppLogger interface {
	Info(args ...any)
	Infof(s string, args ...any)
	Debug(args ...any)
	Debugf(s string, args ...any)
	Warn(args ...any)
	Warnf(s string, args ...any)
	Error(args ...any)
	Errorf(s string, args ...any)
	Fatal(args ...any)
	Fatalf(s string, args ...any)
}

type appLogger struct {
	logLevel       string
	loggerLevelMap map[string]zapcore.Level
	logger         *zap.Logger
	sugaredLogger  *zap.SugaredLogger
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
	l.sugaredLogger = logger.Sugar()
}

func (l appLogger) getLoggerLevel() zapcore.Level {
	loggerLevel, exist := l.loggerLevelMap[l.logLevel]
	if exist {
		return loggerLevel
	}
	return zapcore.InfoLevel
}

func (l *appLogger) Sync() {
	l.logger.Sync()
	l.sugaredLogger.Sync()
}

func (l *appLogger) Info(args ...any) {
	l.sugaredLogger.Info(args...)
}

func (l *appLogger) Infof(s string, args ...any) {
	l.sugaredLogger.Infof(s, args...)
}

func (l *appLogger) Debug(args ...any) {
	l.sugaredLogger.Debug(args...)
}

func (l *appLogger) Debugf(s string, args ...any) {
	l.sugaredLogger.Debugf(s, args...)
}

func (l *appLogger) Warn(args ...any) {
	l.sugaredLogger.Warn(args...)
}

func (l *appLogger) Warnf(s string, args ...any) {
	l.sugaredLogger.Warnf(s, args...)
}

func (l *appLogger) Error(args ...any) {
	l.sugaredLogger.Error(args...)
}

func (l *appLogger) Errorf(s string, args ...any) {
	l.sugaredLogger.Errorf(s, args...)
}

func (l *appLogger) Fatal(args ...any) {
	l.sugaredLogger.Fatal(args...)
}

func (l *appLogger) Fatalf(s string, args ...any) {
	l.sugaredLogger.Fatalf(s, args...)
}
