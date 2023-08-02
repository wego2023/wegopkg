package logrus

import (
	"context"
	"github.com/wegoteam/wepkg/log/wlog"
	"io"

	"github.com/sirupsen/logrus"
)

var _ wlog.FullLogger = (*Logger)(nil)

// Logger logrus impl
type Logger struct {
	l *logrus.Logger
}

// NewLogger create a logger
func NewLogger(opts ...Option) *Logger {
	cfg := defaultConfig()

	// apply options
	for _, opt := range opts {
		opt.apply(cfg)
	}

	// attach hook
	for _, hook := range cfg.hooks {
		cfg.logger.AddHook(hook)
	}

	return &Logger{
		l: cfg.logger,
	}
}

func (l *Logger) Logger() *logrus.Logger {
	return l.l
}

func (l *Logger) Trace(v ...interface{}) {
	l.l.Trace(v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.l.Debug(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.l.Info(v...)
}

func (l *Logger) Notice(v ...interface{}) {
	l.l.Warn(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.l.Warn(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.l.Error(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.l.Fatal(v...)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.l.Tracef(format, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.l.Debugf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.l.Infof(format, v...)
}

func (l *Logger) Noticef(format string, v ...interface{}) {
	l.l.Warnf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.l.Warnf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.l.Errorf(format, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.l.Fatalf(format, v...)
}

func (l *Logger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Tracef(format, v...)
}

func (l *Logger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Debugf(format, v...)
}

func (l *Logger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Infof(format, v...)
}

func (l *Logger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Warnf(format, v...)
}

func (l *Logger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Warnf(format, v...)
}

func (l *Logger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Errorf(format, v...)
}

func (l *Logger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Fatalf(format, v...)
}

func (l *Logger) SetLevel(level wlog.Level) {
	var lv logrus.Level
	switch level {
	case wlog.LevelTrace:
		lv = logrus.TraceLevel
	case wlog.LevelDebug:
		lv = logrus.DebugLevel
	case wlog.LevelInfo:
		lv = logrus.InfoLevel
	case wlog.LevelWarn, wlog.LevelNotice:
		lv = logrus.WarnLevel
	case wlog.LevelError:
		lv = logrus.ErrorLevel
	case wlog.LevelFatal:
		lv = logrus.FatalLevel
	default:
		lv = logrus.WarnLevel
	}
	l.l.SetLevel(lv)
}

func (l *Logger) SetOutput(writer io.Writer) {
	l.l.SetOutput(writer)
}
