package util

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"ports-server/configs"
)

type LoggerEnvironment int64

const (
	Development LoggerEnvironment = iota + 1
	Production
	DevelopmentJson
)

const (
	traceIdKey = "traceID"
)

type Logger struct {
	*zap.Logger
}

type LoggerConfig struct {
	Environment LoggerEnvironment
	Level       zapcore.Level
}

type Field struct {
	Key   string
	Value any
}

func NewField(key string, value any) *Field {
	return &Field{
		Key:   key,
		Value: value,
	}
}

func ErrorField(err error) *Field {
	return &Field{
		Key:   "error",
		Value: err,
	}
}

func New(c *configs.Config) (*Logger, error) {
	l, err := getLogger(c)
	if err != nil {
		return nil, err
	}

	return &Logger{
			l,
		},
		nil
}

func getLogger(c *configs.Config) (*zap.Logger, error) {
	var err error
	var l *zap.Logger
	loggerEnvironment := ToLoggerEnvironment(c.LoggerConfig.Environment)
	loggerLevel := ToLoggerLevel(c.LoggerConfig.Level)
	switch loggerEnvironment {
	case Production:
		lConfig := zap.Config{
			Level:             zap.NewAtomicLevelAt(loggerLevel),
			Development:       false,
			DisableCaller:     false,
			DisableStacktrace: true,
			Sampling:          nil,
			Encoding:          "console",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:          "M",
				LevelKey:            "L",
				TimeKey:             "T",
				NameKey:             "",
				CallerKey:           "C",
				FunctionKey:         "",
				StacktraceKey:       "",
				SkipLineEnding:      false,
				LineEnding:          zapcore.DefaultLineEnding,
				EncodeLevel:         zapcore.CapitalColorLevelEncoder,
				EncodeTime:          zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
				EncodeDuration:      nil,
				EncodeCaller:        zapcore.ShortCallerEncoder,
				EncodeName:          nil,
				NewReflectedEncoder: nil,
				ConsoleSeparator:    " | ",
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			InitialFields:    nil,
		}

		l, err = lConfig.Build()
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	case Development:
		lConfig := zap.Config{
			Level:             zap.NewAtomicLevelAt(loggerLevel),
			Development:       true,
			DisableCaller:     false,
			DisableStacktrace: true,
			Sampling:          nil,
			Encoding:          "console",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:          "M",
				LevelKey:            "L",
				TimeKey:             "T",
				NameKey:             "",
				CallerKey:           "C",
				FunctionKey:         "F",
				StacktraceKey:       "",
				SkipLineEnding:      false,
				LineEnding:          zapcore.DefaultLineEnding,
				EncodeLevel:         zapcore.CapitalColorLevelEncoder,
				EncodeTime:          zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
				EncodeDuration:      nil,
				EncodeCaller:        zapcore.ShortCallerEncoder,
				EncodeName:          nil,
				NewReflectedEncoder: nil,
				ConsoleSeparator:    " | ",
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			InitialFields:    nil,
		}

		l, err = lConfig.Build()
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	case DevelopmentJson:
		lConfig := zap.Config{
			Level:             zap.NewAtomicLevelAt(loggerLevel),
			Development:       true,
			DisableCaller:     false,
			DisableStacktrace: true,
			Sampling:          nil,
			Encoding:          "json",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:          "M",
				LevelKey:            "L",
				TimeKey:             "T",
				NameKey:             "",
				CallerKey:           "C",
				FunctionKey:         "F",
				StacktraceKey:       "",
				SkipLineEnding:      false,
				LineEnding:          zapcore.DefaultLineEnding,
				EncodeLevel:         zapcore.CapitalColorLevelEncoder,
				EncodeTime:          zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
				EncodeDuration:      nil,
				EncodeCaller:        zapcore.ShortCallerEncoder,
				EncodeName:          nil,
				NewReflectedEncoder: nil,
				ConsoleSeparator:    " | ",
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			InitialFields:    nil,
		}

		l, err = lConfig.Build()
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	default:
		l = zap.NewExample()
		l.Warn("app environment not set, set default")
	}

	zap.ReplaceGlobals(l)

	l.Info(fmt.Sprintf("logger instance created with config: %+v", *c))
	return l, nil
}

func getZapFieldsWithCtxData(ctx context.Context, fields ...*Field) []zap.Field {
	zapFields := getZapFields(fields)
	zapFields = append(zapFields, zap.String(traceIdKey, getTracingId(ctx)))
	return zapFields
}

func getZapFields(fields []*Field) []zap.Field {
	var zapFields []zap.Field
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}

func ToLoggerLevel(arg string) zapcore.Level {
	switch arg {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zap.DebugLevel
	}
}

func ToLoggerEnvironment(arg string) LoggerEnvironment {
	switch arg {
	case "development":
		return Development
	case "production":
		return Production
	case "development-json":
		return DevelopmentJson
	default:
		return Development
	}
}

func (log *Logger) InfoCtx(ctx context.Context, msg string, fields ...*Field) {
	log.Logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, getZapFieldsWithCtxData(ctx, fields...)...)
}

func (log *Logger) ErrorCtx(ctx context.Context, msg string, fields ...*Field) {
	log.Logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, getZapFieldsWithCtxData(ctx, fields...)...)
}

func getTracingId(ctx context.Context) string {
	return trace.SpanFromContext(ctx).SpanContext().TraceID().String()
}
