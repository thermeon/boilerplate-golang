package log

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"os"
)

const zeroLogName = "zeroLog"
const traceableId = "uid"

func init() {
	Init(GetDefaultConfigs())
}

func Init(c *Config) {
	lvl, err := zerolog.ParseLevel(string(c.Level))
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(lvl)
	zl := zerolog.New(os.Stderr).With().Timestamp().CallerWithSkipFrameCount(4).Logger()
	inst = &zeroLog{
		L: zl,
	}
}

type zeroLog struct {
	L zerolog.Logger
}

func (z zeroLog) Debug(ctx context.Context, format string, args ...interface{}) {
	z.L.Debug().Str(traceableId, z.getTraceableId(ctx)).Msg(fmt.Sprintf(format, args...))
}

func (z zeroLog) Warning(ctx context.Context, format string, args ...interface{}) {
	z.L.Warn().Str(traceableId, z.getTraceableId(ctx)).Msg(fmt.Sprintf(format, args...))
}

func (z zeroLog) Info(ctx context.Context, format string, args ...interface{}) {
	z.L.Info().Str(traceableId, z.getTraceableId(ctx)).Msg(fmt.Sprintf(format, args...))
}

func (z zeroLog) Error(ctx context.Context, format string, args ...interface{}) {
	z.L.Error().Str(traceableId, z.getTraceableId(ctx)).Msg(fmt.Sprintf(format, args...))
}

func (z zeroLog) Fatal(ctx context.Context, format string, args ...interface{}) {
	z.L.Fatal().Str(traceableId, z.getTraceableId(ctx)).Msg(fmt.Sprintf(format, args...))
}

func (z zeroLog) Panic(ctx context.Context, format string, args ...interface{}) {
	z.L.Panic().Str(traceableId, z.getTraceableId(ctx)).Msg(fmt.Sprintf(format, args...))
}

func (z zeroLog) Name() string {
	return zeroLogName
}

func (z zeroLog) getTraceableId(ctx context.Context) (uid string) {
	uid = ""
	tl, err := getTraceableLog(ctx)
	if err != nil {
		uid = uuid.New().String()
	} else {
		tlUid := tl.GetUUID()
		uid = tlUid.String()
	}
	return
}
