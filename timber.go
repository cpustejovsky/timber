package timber

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger represent common interface for logging function
type Logger interface {
	Sync() error
	CatchPanic()
	Errorf(format string, args ...interface{})
	Errorw(format string, keysAndValues ...interface{})
	Error(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Infow(format string, keysAndValues ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

type zapSugarLogger struct {
	lw *zap.SugaredLogger
}

func NewNopZapLogger() *zapSugarLogger {
	return &zapSugarLogger{lw: zap.NewNop().Sugar()}
}

// NewZapLogger constructs a zapSugarLogger struct that satisfies the Logger interface
// the returned zapSugarLogger writes to stdout and provides human-readable timestamps.
func NewZapLogger(service string) (*zapSugarLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": service,
	}

	log, err := config.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	l := zapSugarLogger{
		lw: log.Sugar(),
	}
	return &l, nil
}

func (z *zapSugarLogger) Sync() error {
	return z.lw.Sync()
}

func (z *zapSugarLogger) CatchPanic() {
	if p := recover(); p != nil {
		z.Errorf("%+v\n", p)
	}
}

func (z *zapSugarLogger) Errorf(format string, args ...interface{}) {
	z.lw.Errorf(format, args...)
}

func (z *zapSugarLogger) Errorw(format string, keysAndValues ...interface{}) {
	z.lw.Errorw(format, keysAndValues...)
}

func (z *zapSugarLogger) Error(args ...interface{}) {
	z.lw.Error(args...)
}

func (z *zapSugarLogger) Fatalf(format string, args ...interface{}) {
	z.lw.Fatalf(format, args...)
}

func (z *zapSugarLogger) Fatal(args ...interface{}) {
	z.lw.Fatal(args...)
}

func (z *zapSugarLogger) Infof(format string, args ...interface{}) {
	z.lw.Infof(format, args...)
}

func (z *zapSugarLogger) Infow(format string, keysAndValues ...interface{}) {
	z.lw.Infow(format, keysAndValues...)
}

func (z *zapSugarLogger) Info(args ...interface{}) {
	z.lw.Info(args...)
}

func (z *zapSugarLogger) Warnf(format string, args ...interface{}) {
	z.lw.Warnf(format, args...)
}

func (z *zapSugarLogger) Debugf(format string, args ...interface{}) {
	z.lw.Debugf(format, args...)
}

func (z *zapSugarLogger) Debug(args ...interface{}) {
	z.lw.Debug(args...)
}

func (z *zapSugarLogger) Printf(format string, args ...interface{}) {
	z.lw.Infof(format, args...)
}

func (z *zapSugarLogger) Println(args ...interface{}) {
	z.lw.Info(args...)
}
