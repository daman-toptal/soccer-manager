package logging

import (
	"context"

	"github.com/sirupsen/logrus"
)

var (
	log = NewLogger()
)

func AddHook(hook Hook) {
	log.AddHook(hook)
}

func StandardLogger() *logrus.Logger {
	return log.StandardLogger()
}

func SetLevel(level logrus.Level) {
	log.SetLevel(level)
}

func Debug(s string, fs ...interface{}) {
	log.Debug(s, fs)
}

func Debugf(s string, fs ...interface{}) {
	log.Debugf(s, fs)
}

func DebugD(s string, f Fields) {
	log.DebugD(s, f)
}

func Info(s string, fs ...interface{}) {
	log.Info(s, fs)
}

func Infof(s string, fs ...interface{}) {
	log.Infof(s, fs)
}

func InfoD(s string, f Fields) {
	log.InfoD(s, f)
}

func Warn(s string, fs ...interface{}) {
	log.Warn(s, fs)
}

func Warnf(s string, fs ...interface{}) {
	log.Warnf(s, fs)
}

func WarnD(s string, f Fields) {
	log.WarnD(s, f)
}

func Error(s string, fs ...interface{}) {
	log.Error(s, fs)
}

func Errorf(s string, fs ...interface{}) {
	log.Errorf(s, fs)
}

func ErrorD(s string, f Fields) {
	log.ErrorD(s, f)
}

func Panic(s string, fs ...interface{}) {
	log.Panic(s, fs)
}

func Panicf(s string, fs ...interface{}) {
	log.Panicf(s, fs)
}

func PanicD(s string, f Fields) {
	log.PanicD(s, f)
}

func Fatal(s string, fs ...interface{}) {
	log.Panic(s, fs)
}

func Fatalf(s string, fs ...interface{}) {
	log.Panicf(s, fs)
}

func DebugWithCtx(ctx context.Context, s string, fs ...interface{}) {
	log.DebugWithCtx(ctx, s, fs)
}

func DebugfWithCtx(ctx context.Context, s string, fs ...interface{}) {
	log.DebugfWithCtx(ctx, s, fs)
}

func DebugDWithCtx(ctx context.Context, s string, f Fields) {
	log.DebugDWithCtx(ctx, s, f)
}

func InfoWithCtx(ctx context.Context, s string, fs ...interface{}) {
	log.InfoWithCtx(ctx, s, fs)
}

func InfofWithCtx(ctx context.Context, s string, fs ...interface{}) {
	log.InfofWithCtx(ctx, s, fs)
}

func InfoDWithCtx(ctx context.Context, s string, f Fields) {
	log.InfoDWithCtx(ctx, s, f)
}

func WarnWithCtx(ctx context.Context, s string, fs ...interface{}) {
	log.WarnWithCtx(ctx, s, fs)
}

func WarnfWithCtx(ctx context.Context, s string, fs ...interface{}) {
	log.WarnfWithCtx(ctx, s, fs)
}

func WarnDWithCtx(ctx context.Context, s string, f Fields) {
	log.WarnDWithCtx(ctx, s, f)
}

func ErrorWithCtx(ctx context.Context, s string, fs ...interface{}) {
	log.ErrorWithCtx(ctx, s, fs)
}

func ErrorfWithCtx(ctx context.Context, s string, fs ...interface{}) {
	log.ErrorfWithCtx(ctx, s, fs)
}

func ErrorDWithCtx(ctx context.Context, s string, f Fields) {
	log.ErrorDWithCtx(ctx, s, f)
}

func PanicWithCtx(ctx context.Context, s string, fs ...interface{}) {
	log.PanicWithCtx(ctx, s, fs)
}

func PanicfWithCtx(ctx context.Context, s string, fs ...interface{}) {
	log.PanicfWithCtx(ctx, s, fs)
}

func PanicDWithCtx(ctx context.Context, s string, f Fields) {
	log.PanicDWithCtx(ctx, s, f)
}

func FatalWithCtx(ctx context.Context, s string, fs ...interface{}) {
	log.PanicWithCtx(ctx, s, fs)
}

func FatalfWithCtx(ctx context.Context, s string, fs ...interface{}) {
	log.PanicfWithCtx(ctx, s, fs)
}

func SetupLogging(level string) {
	switch level {
	case "DEBUG":
		log.SetLevel(logrus.DebugLevel)
		break
	case "INFO":
		log.SetLevel(logrus.InfoLevel)
		break
	default:
		log.SetLevel(logrus.ErrorLevel)
		break
	}
}

func GetLogger() Logger {
	return log
}
