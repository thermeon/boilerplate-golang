package middleware

import (
	"github.com/google/uuid"
	"github.com/thermeon/boilerplate-golang/log"
	"net/http"
)

type TraceableHandler struct {
	handler http.Handler
}

func (t *TraceableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uid := uuid.New()
	ctx := log.WithTraceableCtx(r.Context(), uid)
	r = r.WithContext(ctx)
	log.Debug(ctx, "Request initiated with uuid: "+uid.String())
	t.handler.ServeHTTP(w, r)
}

func NewTraceableHandler(h http.Handler) *TraceableHandler {
	return &TraceableHandler{handler: h}
}
