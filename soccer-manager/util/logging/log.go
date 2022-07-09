package logging

import (
	"context"
	"io"
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	lr "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

const (
	stackKey          = "stack"
	reportLocationKey = "reportLocation"
	requestIdKey      = "request_id"
)

type Logger interface {
	Debug(string, ...interface{})
	Debugf(string, ...interface{})
	DebugD(string, Fields)

	DebugWithCtx(context.Context, string, ...interface{})
	DebugfWithCtx(context.Context, string, ...interface{})
	DebugDWithCtx(context.Context, string, Fields)

	Info(string, ...interface{})
	Infof(string, ...interface{})
	InfoD(string, Fields)

	InfoWithCtx(context.Context, string, ...interface{})
	InfofWithCtx(context.Context, string, ...interface{})
	InfoDWithCtx(context.Context, string, Fields)

	Warn(string, ...interface{})
	Warnf(string, ...interface{})
	WarnD(string, Fields)

	WarnWithCtx(context.Context, string, ...interface{})
	WarnfWithCtx(context.Context, string, ...interface{})
	WarnDWithCtx(context.Context, string, Fields)

	Error(string, ...interface{})
	Errorf(string, ...interface{})
	ErrorD(string, Fields)

	ErrorWithCtx(context.Context, string, ...interface{})
	ErrorfWithCtx(context.Context, string, ...interface{})
	ErrorDWithCtx(context.Context, string, Fields)

	Panic(string, ...interface{})
	Panicf(string, ...interface{})
	PanicD(string, Fields)

	PanicWithCtx(context.Context, string, ...interface{})
	PanicfWithCtx(context.Context, string, ...interface{})
	PanicDWithCtx(context.Context, string, Fields)

	Fatal(string, ...interface{})
	Fatalf(string, ...interface{})

	FatalWithCtx(context.Context, string, ...interface{})
	FatalfWithCtx(context.Context, string, ...interface{})

	Log(string, ...interface{})
	LogWithCtx(context.Context, string, ...interface{})

	SetLevel(level lr.Level)
	AddHook(hook Hook)
	StandardLogger() *lr.Logger
}

type logger struct {
	l *lr.Logger
}

type Fields lr.Fields
type Hook lr.Hook

func NewLogger(opts ...Option) Logger {
	la := &loggerOpts{
		output: os.Stdout,
		format: &lr.JSONFormatter{},
	}

	for _, opt := range opts {
		opt(la)
	}

	l := lr.New()

	l.SetOutput(la.output)
	l.SetFormatter(la.format)

	return logger{
		l: l,
	}
}

type loggerOpts struct {
	output io.Writer
	format lr.Formatter
}

type Option func(*loggerOpts)

func getRequestIdFromContext(ctx context.Context) string {
	data, ok := metadata.FromIncomingContext(ctx)
	requestId := ""
	if ok {
		if v, ok := data["sd-request-id"]; ok {
			requestId = v[0]
		}
	}
	return requestId
}

func SetOutput(i io.Writer) Option {
	return func(opts *loggerOpts) {
		opts.output = i
	}
}

func SetFormat(f lr.Formatter) Option {
	return func(opts *loggerOpts) {
		opts.format = f
	}
}

func (l logger) SetLevel(level lr.Level) {
	l.l.SetLevel(level)
}

func (l logger) AddHook(hook Hook) {
	l.l.AddHook(hook)
}

func (l logger) DebugWithCtx(ctx context.Context, s string, fs ...interface{}) {
	l.l.WithFields(lr.Fields{
		requestIdKey: getRequestIdFromContext(ctx),
	}).Debug(s, fs)
}

func (l logger) Debug(s string, fs ...interface{}) {
	l.l.Debug(s, fs)
}

func (l logger) DebugfWithCtx(ctx context.Context, s string, fs ...interface{}) {
	l.l.WithFields(lr.Fields{
		requestIdKey: getRequestIdFromContext(ctx),
	}).Debugf(s, fs)
}

func (l logger) Debugf(s string, fs ...interface{}) {
	l.l.Debugf(s, fs)
}

func (l logger) DebugDWithCtx(ctx context.Context, s string, f Fields) {
	f[requestIdKey] = getRequestIdFromContext(ctx)
	l.l.WithFields(f.format()).Debug(s)
}

func (l logger) DebugD(s string, f Fields) {
	l.l.WithFields(f.format()).Debug(s)
}

func (l logger) InfoWithCtx(ctx context.Context, s string, fs ...interface{}) {
	l.l.WithFields(lr.Fields{
		requestIdKey: getRequestIdFromContext(ctx),
	}).Info(s, fs)
}

func (l logger) Info(s string, fs ...interface{}) {
	l.l.Info(s, fs)
}

func (l logger) InfofWithCtx(ctx context.Context, s string, fs ...interface{}) {
	l.l.WithFields(lr.Fields{
		requestIdKey: getRequestIdFromContext(ctx),
	}).Infof(s, fs)
}

func (l logger) Infof(s string, fs ...interface{}) {
	l.l.Infof(s, fs)
}

func (l logger) InfoDWithCtx(ctx context.Context, s string, f Fields) {
	f[requestIdKey] = getRequestIdFromContext(ctx)
	l.l.WithFields(f.format()).Info(s)
}

func (l logger) InfoD(s string, f Fields) {
	l.l.WithFields(f.format()).Info(s)
}

func (l logger) WarnWithCtx(ctx context.Context, s string, fs ...interface{}) {
	f := getFields(fs)
	f[requestIdKey] = getRequestIdFromContext(ctx)
	l.l.WithFields(f.format()).Warn(s)
}

func (l logger) Warn(s string, fs ...interface{}) {
	f := getFields(fs)
	l.l.WithFields(f.format()).Warn(s)
}

func (l logger) WarnfWithCtx(ctx context.Context, s string, fs ...interface{}) {
	l.l.WithFields(lr.Fields{
		requestIdKey: getRequestIdFromContext(ctx),
	}).Warnf(s, fs)
}

func (l logger) Warnf(s string, fs ...interface{}) {
	l.l.Warnf(s, fs)
}

func (l logger) WarnDWithCtx(ctx context.Context, s string, f Fields) {
	f[requestIdKey] = getRequestIdFromContext(ctx)
	l.l.WithFields(f.format()).Warn(s)
}

func (l logger) WarnD(s string, f Fields) {
	l.l.WithFields(f.format()).Warn(s)
}

func (l logger) ErrorWithCtx(ctx context.Context, s string, fs ...interface{}) {
	l.l.WithFields(lr.Fields{
		requestIdKey: getRequestIdFromContext(ctx),
	}).Error(s, fs)
}

func (l logger) Error(s string, fs ...interface{}) {
	l.l.Error(s, fs)
}

func (l logger) ErrorfWithCtx(ctx context.Context, s string, fs ...interface{}) {
	l.l.WithFields(lr.Fields{
		requestIdKey: getRequestIdFromContext(ctx),
	}).Errorf(s, fs)
}

func (l logger) Errorf(s string, fs ...interface{}) {
	l.l.Errorf(s, fs)
}

func (l logger) ErrorDWithCtx(ctx context.Context, s string, f Fields) {
	f[requestIdKey] = getRequestIdFromContext(ctx)
	l.l.WithFields(f.format()).Error(s)
}

func (l logger) ErrorD(s string, f Fields) {
	l.l.WithFields(f.format()).Error(s)
}

func (l logger) PanicWithCtx(ctx context.Context, s string, fs ...interface{}) {
	l.l.WithFields(lr.Fields{
		requestIdKey: getRequestIdFromContext(ctx),
	}).Panic(s, fs)
}

func (l logger) Panic(s string, fs ...interface{}) {
	l.l.Panic(s, fs)
}

func (l logger) PanicfWithCtx(ctx context.Context, s string, fs ...interface{}) {
	l.l.WithFields(lr.Fields{
		requestIdKey: getRequestIdFromContext(ctx),
	}).Panicf(s, fs)
}

func (l logger) Panicf(s string, fs ...interface{}) {
	l.l.Panicf(s, fs)
}

func (l logger) PanicDWithCtx(ctx context.Context, s string, f Fields) {
	f = l.appendStack(f)
	f = l.appendReportLocation(f)
	f[requestIdKey] = getRequestIdFromContext(ctx)
	l.l.WithFields(f.format()).Panic(s)
}

func (l logger) PanicD(s string, f Fields) {
	f = l.appendStack(f)
	f = l.appendReportLocation(f)
	l.l.WithFields(f.format()).Panic(s)
}

func (l logger) FatalWithCtx(ctx context.Context, s string, fs ...interface{}) {
	l.l.WithFields(lr.Fields{
		requestIdKey: getRequestIdFromContext(ctx),
	}).Fatal(s, fs)
}

func (l logger) Fatal(s string, fs ...interface{}) {
	l.l.Fatal(s, fs)
}

func (l logger) FatalfWithCtx(ctx context.Context, s string, fs ...interface{}) {
	l.l.WithFields(lr.Fields{
		requestIdKey: getRequestIdFromContext(ctx),
	}).Fatalf(s, fs)
}

func (l logger) Fatalf(s string, fs ...interface{}) {
	l.l.Fatalf(s, fs)
}

func (l logger) LogWithCtx(ctx context.Context, s string, fs ...interface{}) {
	l.InfofWithCtx(ctx, s, fs)
}

func (l logger) Log(s string, fs ...interface{}) {
	l.Infof(s, fs)
}

func (l logger) StandardLogger() *lr.Logger {
	return l.l
}

func (l logger) appendStack(f Fields) Fields {
	f[stackKey] = string(debug.Stack())

	return f
}

func (l logger) appendReportLocation(f Fields) Fields {
	pc, _, line, _ := runtime.Caller(3)

	f[reportLocationKey] = map[string]interface{}{
		"line":         line,
		"functionName": runtime.FuncForPC(pc).Name(),
	}

	return f
}

func (f Fields) format() lr.Fields {
	return lr.Fields(f)
}

func getFields(vfs ...interface{}) Fields {
	fields := Fields{}
	if len(vfs) == 0 {
		return fields
	}
	fs := vfs[0].([]interface{})
	if len(fs) > 0 {
		if f, ok := fs[0].([]interface{}); ok {
			if len(f) > 0 {
				if val, ok := f[0].(Fields); ok {
					fields = val
				}
			}
		}
	}
	return fields
}

func PrettifyStack(stack string) string {
	lines := strings.Split(strings.TrimSpace(stack), "\n")

	if len(lines) > 0 {
		if first := lines[0]; strings.HasPrefix(first, "goroutine ") && strings.HasSuffix(first, ":") {
			lines = lines[1:]
		}
	}

	sb := strings.Builder{}

	for _, line := range lines {
		if strings.HasPrefix(line, "\t") {
			line = line[1:]
			if offset := strings.LastIndex(line, " +0x"); offset != -1 {
				line = line[:offset]
			}
			sb.WriteString(" (")
			sb.WriteString(line)
			sb.WriteString(")")
			continue
		}

		if sb.Len() > 0 {
			sb.WriteByte('\n')
		}
		sb.WriteString(line)
	}

	return sb.String()
}
