package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func Init() {
	zapLogger, _ = zap.NewProduction()
}

func getLogger() *zap.Logger {
	if zapLogger == nil {
		Init()
	}
	return zapLogger
}

func Debug(msg string) {
	getLogger().Debug(msg)
}

func Debugf(format string, v ...any) {
	Debug(fmt.Sprintf(format, v...))
}

func Info(msg string) {
	getLogger().Info(msg)
}

func Infof(format string, v ...any) {
	Info(fmt.Sprintf(format, v...))
}

func Warn(msg string) {
	getLogger().Warn(msg)
}

func Warnf(format string, v ...any) {
	Warn(fmt.Sprintf(format, v...))
}

func Error(msg string) {
	getLogger().Error(msg)
}

func Errorf(format string, v ...any) {
	Error(fmt.Sprintf(format, v...))
}

func Fatal(msg string) {
	getLogger().Fatal(msg)
}

func Fatalf(format string, v ...any) {
	Fatal(fmt.Sprintf(format, v...))
}
