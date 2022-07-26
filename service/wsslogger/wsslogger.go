package wsslogger

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"net/http"
)

type ContextKey string

type Logger struct {
	driver      zerolog.Logger
	contextKeys []ContextKey
}

func NewLogger(service string, writers []io.Writer, keysFromContext []ContextKey) *Logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	multi := zerolog.MultiLevelWriter(writers...)
	logger := zerolog.New(multi).
		With().
		Timestamp().
		Str("service", service).
		Logger()

	return &Logger{
		driver:      logger,
		contextKeys: keysFromContext,
	}
}

func (l *Logger) Info(msg string, ctx context.Context) {
	e := l.setStandardFields(l.driver.Info(), ctx)
	e.Msg(msg)
}

func (l *Logger) Err(err error, ctx context.Context) {
	e := l.setStandardFields(l.driver.Error(), ctx)
	e.Err(err).Stack().Msg("")
}

func (l *Logger) IncomingRequest(r *http.Request) {
	e := l.setRequestLogFields(l.setStandardFields(l.driver.Log(), r.Context()), r)
	e.Msg("Incoming request")
}

func (l *Logger) OutgoingRequest(r *http.Request) {
	e := l.setRequestLogFields(l.setStandardFields(l.driver.Log(), r.Context()), r)
	e.Msg("Outgoing request")
}

func (l *Logger) OutgoingResponse(r *http.Request) {
	e := l.setRequestLogFields(l.setStandardFields(l.driver.Log(), r.Context()), r)
	e.Msg("Outgoing response")
}

func (l *Logger) IncomingResponse(r *http.Request) {
	e := l.setRequestLogFields(l.setStandardFields(l.driver.Log(), r.Context()), r)
	e.Msg("Incoming request")
}

func (l *Logger) setStandardFields(e *zerolog.Event, ctx context.Context) *zerolog.Event {
	if ctx != nil {
		for _, key := range l.contextKeys {
			cVal, ok := ctx.Value(key).(string)
			if ok {
				e.Str(string(key), cVal)
			}
		}
	}

	return e
}

func (l *Logger) setRequestLogFields(e *zerolog.Event, r *http.Request) *zerolog.Event {
	return e.Str("remote_ip", r.RemoteAddr).
		Str("http_method", r.Method).
		Str("uri", r.URL.RequestURI())
}
