package main

import (
	"context"
	"git.slygods.com/evoplay/wss-go/pkg/helper"
	"github.com/google/uuid"
	"github.com/rs/xid"
	"net/http"
)

func (app *application) requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.Logger.IncomingRequest(r)

		next.ServeHTTP(w, r)
	})
}

func (app *application) traceIds(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqIdHeaderName := helper.Env("REQUEST_ID_HEADER_NAME", "X-Request-Id")
		requestId := r.Header.Get(reqIdHeaderName)
		if requestId == "" {
			requestId = uuid.New().String()
		}

		spanIdHeaderName := helper.Env("PARENT_SPAN_ID_HEADER_NAME", "X-Span-Id")
		parentSpanId := r.Header.Get(spanIdHeaderName)
		spanId := xid.New().String()

		ctx := r.Context()
		ctx = context.WithValue(ctx, keyRequestId, requestId)
		ctx = context.WithValue(ctx, keySpanId, spanId)
		ctx = context.WithValue(ctx, keyParentSpanId, parentSpanId)
		r = r.WithContext(ctx)

		headers := w.Header()
		headers[reqIdHeaderName] = []string{requestId}
		headers[spanIdHeaderName] = []string{spanId}

		next.ServeHTTP(w, r)
	})
}
