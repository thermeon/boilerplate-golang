package log

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/thermeon/boilerplate-golang/common"
)

const uuidKey common.ContextKey = "uuid"

type TraceableLog interface {
	GetUUID() (uid uuid.UUID)
	SetUUID(uid uuid.UUID)
}

type traceableLog struct {
	UID uuid.UUID
}

func (t *traceableLog) GetUUID() (uid uuid.UUID) {
	return t.UID
}

func (t *traceableLog) SetUUID(uid uuid.UUID) {
	t.UID = uid
	return
}

func WithTraceableCtx(ctx context.Context, uid uuid.UUID) (ctxR context.Context) {
	tl := traceableLog{}
	tl.SetUUID(uid)
	return context.WithValue(ctx, uuidKey, &tl)
}

func getTraceableLog(ctx context.Context) (tl TraceableLog, err error) {
	ok := true
	tl, ok = ctx.Value(uuidKey).(TraceableLog)
	if !ok {
		err = errors.New("unable to find traceable context")
	}
	return
}
